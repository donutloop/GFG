package product

type BaseProduct struct {
	ProductID int    `json:"-"`
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	Brand     string `json:"brand"`
	Stock     int    `json:"stock"`
}

type Product struct {
	BaseProduct
	SellerUUID string `json:"seller_uuid"`
}

type ProductV2 struct {
	BaseProduct
	Seller Seller `json:"seller"`
}

type Seller struct {
	UUID  string `json:"uuid"`
	Links Links  `json:"_links"`
}

// todo(marcel) move to generic pkg
type Links struct {
	Self SelfLink `json:"self"`
}

type SelfLink struct {
	Href string `json:"href"`
}
