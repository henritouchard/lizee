package products

import "lizee/pkg/errortypes"

type datesOrder struct {
	from string
	to   string
}

// ProcessRentalOrder Check if a product is available in base and
// execute rental if possible. in case it's not, no insert are executed,
// and a list of unavailable products is returned
func ProcessRentalOrder(orders []ProductOrder) ([]UnAvailable, error) {

	// check dates
	fromTo := make(map[datesOrder]bool)
	// check if products are still available
	for i := range orders {
		from, to := orders[i].FromDate, orders[i].ToDate
		from, to, err := parseDates(from, to)
		if err != nil {
			return nil, err
		}

		if fromTo[datesOrder{from, to}] == false {
			fromTo[datesOrder{from, to}] = true
		}
	}

	var availables []Available
	unavailables := make([]UnAvailable, 0)
	var err error
	if len(fromTo) == 1 {
		// all orders are done at the same period
		for key := range fromTo {
			availables, err = AllAvailable(&ProductQuery{0, key.from, key.to})
			if err != nil {
				return nil, err
			}
		}

		for _, o := range orders {
			for _, av := range availables {
				if o.Product.ID == av.ProductID && o.Quantity > av.Available {
					unavailables = append(unavailables, UnAvailable{o.Product.ID, o.Quantity, av.Available})
				}
			}
		}
	}

	// TODO : Same task if there are several periods asked in orders
	if len(unavailables) > 0 {
		return unavailables, errortypes.New(errortypes.UnavailableProduct)
	}

	// No error found, prepare db statement
	insertStmt, err := db.Prepare(queryInsert)
	if err != nil {
		return nil, err
	}
	// Execute queries.
	// TODO: find a way to aggregate queries and execute everything in once
	for _, order := range orders {
		order.FromDate, order.ToDate, err = parseDates(order.FromDate, order.ToDate)
		_, err := insertStmt.Exec(order.Product.ID, order.Quantity, order.FromDate, order.ToDate)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
