package seller

type Seller struct {
	SellerID int `json:"-"`
	UUID string  `json:"uuid"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
