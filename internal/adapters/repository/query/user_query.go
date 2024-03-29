package query

const GetUser = `
SELECT 
	u.user_id,
	CONCAT(u.first_name, ' ',u.last_name) as full_name,
	u.phone_no,
	u.address,
	pr.payment_due as balance
FROM users u
	LEFT JOIN payment_record pr 
	ON pr.user_id = u.user_id
WHERE u.user_id = ?
`

const GetUsers = `
SELECT 
	u.user_id,
	CONCAT(u.first_name, ' ',u.last_name) as full_name,
	u.phone_no,
	u.address,
	pr.payment_due as balance
FROM users u
	LEFT JOIN payment_record pr 
	ON pr.user_id = u.user_id
`
const AddUser = `
INSERT INTO users 
	(first_name,last_name,phone_no,address) 
	VALUES('?','?','?','?')
`

const UpdateUser = `
UPDATE users SET 
	first_name = ?, 
	last_name = ?,
	phone_no = ?,
	address = ?
WHERE user_id = ?
`

const DeleteUser = `
DELETE
FROM users
WHERE user_id = ?
`
