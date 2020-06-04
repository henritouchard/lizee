package products

const (
	// Note about the OVERLAPS psql function :
	// there was only 3 days of delay from shipping to rental and 4 days from rental to
	// cleaning. OVERLAPS will return true if dates overlaps however this
	//   (DATE '2011-01-28', DATE '2011-02-01') OVERLAPS
	//   (DATE '2011-02-01', DATE '2011-02-01');
	// returns false. for our case it means that we could ship the product the same day
	// it's cleaned, what we can't do since it's back in stock the day after.
	// That's why we compare OVERLAPS with one day more than delay, both at begining and end of rental
	queryProductAvailability = `
	SELECT p.stock - 
	(
		SELECT COALESCE(SUM(r.quantity), 0) 
			FROM rental_order as r 
			WHERE r.product = p.id
				AND (r.start_date - INTERVAL '3 DAYS', r.end_date + INTERVAL '4 DAYS') 
					OVERLAPS ($2::date - INTERVAL '4 DAYS', $3::date + INTERVAL '5 DAYS')
					IS TRUE
	) AS availability
	FROM product as p WHERE p.id=$1
	`

	queryProductAvailabilityByCategory = `
	SELECT p.id , p.name, p.stock - 
	(
		SELECT COALESCE(SUM(r.quantity), 0)  
			FROM rental_order as r 
			WHERE r.product = p.id
				AND (r.start_date - INTERVAL '3 DAYS', r.end_date + INTERVAL '4 DAYS') 
					OVERLAPS ($1::date - INTERVAL '4 DAYS', $2::date + INTERVAL '5 DAYS') 
					IS TRUE
	) AS availability, p.picture
		FROM product as p WHERE cstr_category=$3
	`

	queryCheckAvailabilityByProduct = `
	SELECT $2 <= p.stock - (
		SELECT COALESCE(SUM(r.quantity), 0) as total
		FROM rental_order as r 
		WHERE r.product = p.id
		AND (r.start_date - INTERVAL '3 DAYS', r.end_date + INTERVAL '4 DAYS') 
			OVERLAPS ($3::date - INTERVAL '4 DAYS', $4::date + INTERVAL '5 DAYS')) 
		AS available FROM product as p WHERE p.id = $1;
	`

	queryCheckAllAvailable = `
	SELECT cast(p.id as text) as product_id, p.stock - (
		SELECT COALESCE(SUM(r.quantity), 0) as total
		FROM rental_order as r 
		WHERE r.product = p.id
		AND (r.start_date - INTERVAL '3 DAYS', r.end_date + INTERVAL '4 DAYS') 
			OVERLAPS ($1::date - INTERVAL '4 DAYS', $2::date + INTERVAL '5 DAYS')) 
		AS available FROM product as p;
	`

	queryInsert = `
	INSERT INTO rental_order(start_date, end_date, returned, product, quantity) VALUES($3, $4, FALSE, $1, $2);
	`
	queryModifyQuantity = `UPDATE product SET stock=$2 WHERE id=$1`
)
