package query

const AddProduct = `
INSERT INTO product_details 
	(product_name,product_description,product_price) 
	VALUES(?,?,?)
	RETURNING id
`
const AddPaymentRecord = `
INSERT INTO payment_record
	(payment_due,user_id,product_id)
	VALUES(?,?,?)
	RETURNING id
`

const AddPaymentLogs = `
INSERT INTO payment_logs
	(credit,debit,record_id)
	VALUES(?,?,?)
`
