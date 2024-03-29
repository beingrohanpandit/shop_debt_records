package domain

type Record struct {
	UserId         int            `json:"user_id"`
	CreditAmount   float64        `json:"credit_amount"`
	ProductDetails ProductDetails `json:"product_details"`
}

type ProductDetails struct {
	ProductId          int     `json:"product_id"`
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	ProductPrice       float64 `json:"product_price"`
}
