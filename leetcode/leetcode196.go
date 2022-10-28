package main

var _ = `
	DELETE a 
FROM
	Person a,
	Person b 
WHERE
	b.id < a.id 
	AND b.email = a.email
`

var _ = `
	DELETE 
FROM
	Person 
WHERE
	id NOT IN (
	SELECT
		id 
	FROM
	( SELECT MIN( id ) AS id FROM Person GROUP BY email ) t)
`
