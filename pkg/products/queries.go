package products

const (
	queryProductAvailability = `
	SELECT p.stock - 
	(
		SELECT COALESCE(SUM(r.quantity), 0) 
			FROM rental_order as r 
			WHERE r.product = p.id
				AND (r.start_date, r.end_date + INTERVAL '5 DAYS') OVERLAPS ($2::date - INTERVAL '2 DAYS', $3::date)
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
				AND (r.start_date, r.end_date + INTERVAL '5 DAYS') OVERLAPS ($1::date - INTERVAL '2 DAYS', $2::date) 
				IS TRUE
	) AS availability, p.picture
		FROM product as p WHERE cstr_category=$3
	`

	queryCheckAvailabilityByProduct = `
	SELECT $2 <= p.stock - (
		SELECT COALESCE(SUM(r.quantity), 0) as total
		FROM rental_order as r 
		WHERE r.product = p.id
		AND (r.start_date, r.end_date + INTERVAL '5 DAYS') OVERLAPS ($3::date - INTERVAL '2 DAYS', $4::date)) 
		AS available FROM product as p WHERE p.id = $1;
	`

	queryCheckAllAvailable = `
	SELECT cast(p.id as text) as product_id, p.stock - (
		SELECT COALESCE(SUM(r.quantity), 0) as total
		FROM rental_order as r 
		WHERE r.product = p.id
		AND (r.start_date, r.end_date + 
        INTERVAL '5 DAYS') OVERLAPS ($1::date - INTERVAL '2 DAYS', $2::date)) 
		AS available FROM product as p;
	`

	queryInsert = `
	INSERT INTO rental_order(start_date, end_date, returned, product, quantity) VALUES($3, $4, FALSE, $1, $2);
	`
)
