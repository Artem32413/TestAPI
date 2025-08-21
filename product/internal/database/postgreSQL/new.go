package postgreSQL

import (
	"product/internal/config/databaseConfig"

	"github.com/jackc/pgx/v5"
)

var (
	displayProducts = `
	SELECT 
    p.productId,
    p.name,
    p.description,
    p.weight,
    p.barcode
FROM 
    products p
`
	displayKeyValue = `SELECT * FROM ProductKeyValues WHERE productId = $1`
	addingAProducts = `INSERT INTO Products (productId, name, description, weight, barcode) VALUES ($1, $2, $3, $4, $5)`
	addingKeyValue  = `INSERT INTO ProductKeyValues (productId, key, value) VALUES ($1, $2, $3)`
	deleteKeyValue  = `DELETE FROM ProductKeyValues WHERE productId = $1`
	updateAProducts = `UPDATE Products SET description = $1 WHERE productId = $2`
	updateValue     = `UPDATE ProductKeyValues SET value = $1 WHERE productId = $2 AND key = $3`
	productCheckUpd = `SELECT EXISTS(SELECT 1 FROM Products WHERE productId = $1)`
	productCheckIns = `SELECT EXISTS(SELECT 1 FROM Products WHERE name = $1)`
)

type DBService struct {
	db *pgx.Conn
}

func New(db *databaseConfig.PostgreSQL) *DBService {
	return &DBService{db: db.Db}
}
