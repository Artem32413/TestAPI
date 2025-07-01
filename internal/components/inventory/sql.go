package inventory

import (
	"context"
	"fmt"
)

var (
	priceIndication   = `UPDATE Inventory SET price = $1 WHERE warehouse_id = $2 AND product_id = $3`
	updateQuantity    = `UPDATE Inventory SET quantity = $1 WHERE warehouse_id = $2 AND product_id = $3`
	creatingADiscount = `UPDATE Inventory SET discount = $1 WHERE warehouse_id = $2 AND product_id = $3`
	listOfGoods       = `SELECT 
    p.name, 
    p.description, 
    i.price, 
    i.discount,
    (i.price * (1 - COALESCE(i.discount, 0) / 100)) AS discounted_price
FROM Inventory i
JOIN Products p ON i.product_id = p.id
WHERE i.warehouse_id = $1
ORDER BY p.name
LIMIT $2 OFFSET (($3 - 1) * $2)`
	listProduct = `SELECT * FROM Inventory WHERE warehouse_id = $1 AND product_id = $2`
	listCount   = `WITH product_data AS (
        SELECT 
            i.product_id,
            i.price, 
            i.discount,
            (i.price * (1 - COALESCE(i.discount, 0) / 100)) AS discounted_price
        FROM inventory i
        WHERE i.warehouse_id = $1 AND i.product_id = ANY($2)
    )
    SELECT 
        pd.product_id,
        pd.price,
        pd.discounted_price,
        p.quantity,
        CASE 
            WHEN pd.discounted_price > 0 THEN pd.discounted_price * p.quantity
            ELSE pd.price * p.quantity
        END AS subtotal
    FROM unnest($3::int[]) AS p(product_id, quantity)
    JOIN product_data pd ON pd.product_id = p.product_id`
	quantityCheck   = `SELECT quantity FROM Inventory WHERE quantity = $1`
	purchaseProduct = `UPDATE Inventory SET quantity = $1 WHERE warehouse_id = $2 AND product_id = $3`
)

type ListByWarehouse struct {
	Identifier string  `json:"identifier"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Discount   float64 `json:"discount"`
}

type AllInformationAboutTheProduct struct {
	Identifier      string  `json:"identifier"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Characteristics string  `json:"characteristics"`
	Barcode         string  `json:"barcode"`
	Price           float64 `json:"price"`
	Discount        float64 `json:"discount"`
	Quantity        int     `json:"quantity"`
}

type SummingUp struct {
	Sum float64 `json:"sum"`
}

type NewProd struct {
	Quantity int
}

func (s *InventoryService) ConnectingTable(price Inventory) error {

	if _, err := s.Db.Exec(context.Background(), priceIndication, price.Price, price.Warehouse_id, price.Product_id); err != nil {
		return err
	}

	return nil
}

func (s *InventoryService) UpdateQuantity(inventory Inventory) error {

	if _, err := s.Db.Exec(context.Background(), updateQuantity, inventory.Quantity, inventory.Warehouse_id, inventory.Product_id); err != nil {
		return err
	}

	return nil
}

func (s *InventoryService) CreatingADiscount(discount Inventory) error {

	if _, err := s.Db.Exec(context.Background(), creatingADiscount, discount.Discount, discount.Warehouse_id, discount.Product_id); err != nil {
		return err
	}

	return nil
}

func (s *InventoryService) List(product Inventory, perPage int, offset int) ([]ListByWarehouse, error) {

	r, err := s.Db.Query(context.Background(), listOfGoods, product.Warehouse_id, perPage, offset)
	if err != nil {
		return nil, err
	}

	var newL []ListByWarehouse

	for r.Next() {

		var n ListByWarehouse

		err = r.Scan(&n.Identifier, &n.Name, &n.Price, &n.Discount)
		if err != nil {
			return nil, err
		}

		newL = append(newL, n)
	}

	return newL, nil
}

func (s *InventoryService) ListProduct(product Inventory) (AllInformationAboutTheProduct, error) {

	r, err := s.Db.Query(context.Background(), listProduct, product.Warehouse_id, product.Product_id)
	if err != nil {
		return AllInformationAboutTheProduct{}, err
	}

	var n AllInformationAboutTheProduct

	for r.Next() {
		err = r.Scan(&n.Identifier, &n.Name, &n.Description, &n.Characteristics, &n.Barcode, &n.Price, &n.Discount, &n.Quantity)
		if err != nil {
			return AllInformationAboutTheProduct{}, err
		}
	}

	return n, nil
}

func (s *InventoryService) ListCount(count Inventory) (SummingUp, error) {

	r, err := s.Db.Query(context.Background(), listCount, count.Warehouse_id, count.Product_id, count.Quantity)
	if err != nil {
		return SummingUp{}, err
	}

	var n SummingUp

	for r.Next() {
		err = r.Scan(&n.Sum)
		if err != nil {
			return SummingUp{}, err
		}
	}

	return n, nil
}

func (s *InventoryService) Purchase(purchase NewInventory) error {

	r, err := s.Db.Query(context.Background(), quantityCheck, purchase.Quantity)
	if err != nil {
		return err
	}

	var sale NewProd

	if r.Next() {
		if err = r.Scan(sale.Quantity); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("На складе отсутствует товар")
	}

	if purchase.Quantity > sale.Quantity {
		return fmt.Errorf("Товар отсутствует или это количество товара на складе отсутствует")
	}

	if _, err := s.Db.Exec(context.Background(), purchaseProduct, purchase.Quantity, purchase.Warehouse_id, purchase.Product_id); err != nil {
		return err
	}

	return nil
}