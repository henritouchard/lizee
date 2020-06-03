package products

import (
	"strconv"
	"strings"
)

const (
	queryProductAvailability = `
	SELECT COALESCE(SUM(r.quantity), 0) as quantity 
		FROM rental_order 
		WHERE product=$1 
			AND (start_date, end_date + INTERVAL '4 DAYS') 
				OVERLAPS ($2::date - INTERVAL '2 DAYS', $3::date) 
			IS FALSE;`
	queryProductAvailabilityByCategory = `
	SELECT p.id , p.name, p.stock - 
	(
		SELECT COALESCE(SUM(r.quantity), 0)  
			FROM rental_order as r 
			WHERE r.product = p.id
				AND (r.start_date, r.end_date + INTERVAL '5 DAYS') OVERLAPS ($1::date - INTERVAL '2 DAYS', $2::date) 
				IS TRUE
	) AS availability, p.picture
		FROM product as p WHERE cstr_category=$3
	`
	// following qRO (queryRentalOrder) are part of a block of code that will be executed by postgres
	// The matter is that this kind of block doesn't accept any argument.
	qROEqualFromDate = `
	DO
	$do$
	BEGIN
		IF (
		    SELECT p.stock - (
		    SELECT COALESCE(SUM(r.quantity), 0) as total
		    FROM rental_order as r 
		    WHERE r.product = p.id AND (r.start_date, r.end_date + INTERVAL '5 DAYS') OVERLAPS ('`
	qROEqualToDate   = `'::date - INTERVAL '2 DAYS', '`
	qROEqualPID      = `'::date)) FROM product as p WHERE p.id =`
	qROEqualQuantity = `) >=`
	qROInsert        = ` THEN INSERT INTO rental_order(start_date, end_date, returned, product, quantity) VALUES(`
	qROEnd           = `);
		ELSE 
		   UPDATE rental_order SET quantity=2  WHERE product=3;
		END IF;
	END
	$do$
	`
)

// following queries are part of a block of code that will be executed by postgres
// The matter is that this kind of block doesn't accept any argument.
func buildRentalQuery(order ProductOrder) (string, error) {
	fromDate := order.FromDate
	toDate := order.ToDate

	pid := strconv.Itoa(order.Product.ID)
	quantity := strconv.Itoa(order.Quantity)

	insertValues := []string{"'" + fromDate + "'", "'" + toDate + "'", "FALSE", pid, quantity}

	query := qROEqualFromDate + fromDate +
		qROEqualToDate + toDate +
		qROEqualPID + pid +
		qROEqualQuantity + quantity +
		qROInsert + strings.Join(insertValues, ",") +
		qROEnd
	return query, nil
}
