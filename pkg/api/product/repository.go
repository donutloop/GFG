package product

import (
	"database/sql"
	"errors"
	"reflect"
)

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

type repository struct {
	db *sql.DB
}

func (r *repository) delete(product *Product) error {
	rows, err := r.db.Query("DELETE FROM product WHERE uuid = ?", product.UUID)

	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

func (r *repository) insert(product *Product) error {
	rows, err := r.db.Query(
		"INSERT INTO product (id_product, name, brand, stock, fk_seller, uuid) VALUES(?,?,?,?,(SELECT id_seller FROM seller WHERE uuid = ?),?)",
		product.ProductID, product.Name, product.Brand, product.Stock, product.SellerUUID, product.UUID,
	)

	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

func (r *repository) update(product *Product) error {
	rows, err := r.db.Query(
		"UPDATE product SET name = ?, brand = ?, stock = ? WHERE uuid = ?",
		product.Name, product.Brand, product.Stock, product.UUID,
	)

	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

func (r *repository) list(offset int, limit int) ([]*Product, error) {
	rows, err := r.db.Query(
		"SELECT p.id_product, p.name, p.brand, p.stock, s.uuid, p.uuid FROM product p " +
			"INNER JOIN seller s ON(s.id_seller = p.fk_seller) LIMIT ? OFFSET ?",
			limit, offset,
		)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*Product

	for rows.Next() {
		product := &Product{}

		err = rows.Scan(&product.ProductID, &product.Name, &product.Brand, &product.Stock, &product.SellerUUID, &product.UUID)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

var ErrNotFound = errors.New("Object not found")

func (r *repository) findByUUID(uuid string, product interface{}) (error) {

	reflectValue := reflect.ValueOf(product)
	if reflectValue.Kind() != reflect.Ptr {
		panic("values isn't a pointer")
	}

	row := r.db.QueryRow(
		"SELECT p.id_product, p.name, p.brand, p.stock, s.uuid, p.uuid FROM product p " +
			"INNER JOIN seller s ON(s.id_seller = p.fk_seller) WHERE p.uuid = ?",
			uuid,
		)

	if row.Err() != nil {
		if row.Err() == sql.ErrNoRows {
			return ErrNotFound
		}
		return row.Err()
	}

	var err error
	if productV1, ok := product.(*Product); ok {
		err = row.Scan(&productV1.ProductID, &productV1.Name, &productV1.Brand, &productV1.Stock, &productV1.SellerUUID, &productV1.UUID)
	} else if productV2, ok := product.(*ProductV2); ok {
		err = row.Scan(&productV2.ProductID, &productV2.Name, &productV2.Brand, &productV2.Stock, &productV2.Seller.UUID, &productV2.UUID)
	} else {
		panic("injected the wrong object")
	}

	if err != nil {
		return err
	}

	return nil
}