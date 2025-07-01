package product

import "context"

var (
	displayProducts = `SELECT * FROM Products`
	addingAProducts = `INSERT INTO Products (identifier, name, description, weight, barcode) VALUES ($1, $2, $3, $4, $5)`
	updateAProducts = `UPDATE Products SET description = $1, keyvalue = $2 WHERE identifier = $3`
	productCheck    = `SELECT EXISTS(SELECT 1 FROM Products WHERE identifier = $1)`
)

func (s *InventoryService) DisplayProducts() ([]Products, error) {

	r, err := s.Db.Query(context.Background(), displayProducts)
	if err != nil {
		return nil, err
	}

	var newSl []Products

	for r.Next() {
		var np Products

		if err = r.Scan(&np.Identifier, &np.Name, &np.Description, &np.KeyValue, &np.Weight, &np.Barcode); err != nil {
			return nil, err
		}

		newSl = append(newSl, Products{np.Identifier, np.Name, np.Description, np.KeyValue, np.Weight, np.Barcode})
	}

	return newSl, nil
}

func (s *InventoryService) AdditionProducts(products Products) error {

	if _, err := s.Db.Exec(context.Background(), addingAProducts, products.Description, products.KeyValue); err != nil {
		return err
	}

	return nil
}

func (s *InventoryService) UpdateProd(products Products, identifier int) error {

	var exists bool
	err := s.Db.QueryRow(context.Background(), productCheck, products.Identifier).Scan(&exists)

	if err != nil {
		return err
	}

	if !exists {
		return err
	}

	result, err := s.Db.Exec(context.Background(), updateAProducts, products.Description, products.KeyValue, identifier)

	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return err
	}

	return nil
}