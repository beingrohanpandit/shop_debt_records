package domain

type GetUser struct {
	UserId   int     `json:"user_id"`
	FullName string  `json:"full_name"`
	PhoneNo  string  `json:"phone_no"`
	Address  string  `json:"address"`
	Balance  float64 `json:"balance"`
}

type User struct {
	UserId    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	PhoneNo   string `json:"phone_no"`
	Address   string `json:"address"`
}
