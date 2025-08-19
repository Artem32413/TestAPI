package postgreSQL

import (
	"analytics/internal/config/databaseConfig"

	"github.com/jackc/pgx/v5"
)

var (
	// Аналитика
	existsWarehouse  = `SELECT EXISTS(SELECT 1 FROM Inventory WHERE warehouseId = $1)`
	productAnalytics = `SELECT 
							warehouseId,
							productId,
							SUM(sold_goods) AS total_sold,
							SUM(total_sum) AS total_revenue
						FROM 
							Analytics
						WHERE 
							warehouseId = $1
						GROUP BY 
							warehouseId, productId 
						ORDER BY 
							total_revenue DESC;`
	topWarehouses = `
						SELECT 
						w.warehouseId,
						w.addr,
						SUM(a.total_sum) AS total_revenue
					FROM 
						WarehousesTable w
					JOIN 
						Analytics a ON w.warehouseId = a.warehouseId
					GROUP BY 
						w.warehouseId, w.addr
					ORDER BY 
						total_revenue DESC
					LIMIT 10;
						`
)

type DBService struct {
    db *pgx.Conn
}

func New(db *databaseConfig.PostgreSQL) *DBService {
    return &DBService{db: db.Db}
}



