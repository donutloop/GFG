package product

type product struct {
	ProductID int         `json:"-"`
	UUID string           `json:"uuid"`
	Name string           `json:"name"`
	Brand string          `json:"brand"`
	Stock int             `json:"stock"`
	SellerUUID string `json:"seller_uuid"`
}