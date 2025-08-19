package postgreSQL

import (
	"inventory/internal/config/databaseConfig"

	"github.com/jackc/pgx/v5"
)

var (
	priceInsert    = `INSERT INTO Inventory (warehouse_id, product_id, quantity, price, discount) VALUES ($1, $2, $3, $4, $5)`
	priceUpdate    = `UPDATE Inventory SET price = $1 WHERE warehouse_id = $2 AND product_id = $3`
	updateQuantity = `UPDATE Inventory 
						SET quantity = quantity + $1
						WHERE warehouse_id = $2 AND product_id = $3`
	creatingADiscount = `UPDATE Inventory SET discount = $1 WHERE warehouse_id = $2 AND product_id = $3`
	listOfGoods       = `SELECT 
						p.product_id,
						p.name,
						i.price,
						ROUND(i.price * (1 - i.discount/100), 2) as discounted_price
						FROM inventory i
						JOIN products p ON i.product_id = p.product_id
						WHERE i.warehouse_id = $1
						LIMIT $2 OFFSET $3`
	listInventory = `	SELECT 
						COALESCE(quantity, 0) AS quantity,
						COALESCE(price, 0) AS price,
						COALESCE(discount, 0) AS discount 
						FROM Inventory 
						WHERE warehouse_id = $1 AND product_id = $2`
	oneProduct = `SELECT 
					p.product_id,
					p.name,
					p.description,
					p.weight,
					p.barcode
					FROM 
					Products p
					WHERE product_id = $1`
	listCount = `SELECT COALESCE(SUM(i.price * p.quantity * (1 - i.discount/100)), 0)
					FROM unnest($1::text[], $2::int[]) AS p(product_id, quantity)
					JOIN inventory i ON i.product_id = p.product_id AND i.warehouse_id = $3
					WHERE i.quantity >= p.quantity`
	quantityCheck = `SELECT quantity FROM Inventory 
						WHERE warehouse_id = $1 AND product_id = $2`
	purchaseProduct = `UPDATE Inventory 
						SET quantity = quantity - $1 
						WHERE warehouse_id = $2 AND product_id = $3`
	purchaseProductAnalytics = `
								INSERT INTO analytics (warehouse_id, product_id, sold_goods, total_sum)
								SELECT 
									$1::text, 
									$2::text, 
									$3::integer,
									$3::integer * (
										SELECT price * (1 - COALESCE(discount, 0)) 
										FROM Inventory 
										WHERE product_id = $2::text
									)
								ON CONFLICT (warehouse_id, product_id) 
								DO UPDATE SET
									sold_goods = analytics.sold_goods + EXCLUDED.sold_goods,
									total_sum = analytics.total_sum + EXCLUDED.total_sum`
)

type DBService struct {
	db *pgx.Conn
}

func New(db *databaseConfig.PostgreSQL) *DBService {
	return &DBService{db: db.Db}
}
