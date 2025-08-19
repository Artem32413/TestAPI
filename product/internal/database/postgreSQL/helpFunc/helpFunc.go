package helpfunc

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func ConvertMapToSlice(attrs map[string]string) []map[string]string {
	var result []map[string]string

	for key, value := range attrs {
		result = append(result, map[string]string{key: value})
	}

	return result
}

func GetAllAttributes(db *pgx.Conn, ctx context.Context, productIDs []string) (map[string]map[string]string, error) {
	if len(productIDs) == 0 {
		return make(map[string]map[string]string), nil
	}

	query := `SELECT productId, key, value FROM ProductKeyValues WHERE productId = ANY($1)`

	rows, err := db.Query(ctx, query, productIDs)
	if err != nil {
		return nil, fmt.Errorf("Ошибка в поиске клю-значения: %w", err)
	}

	defer rows.Close()

	attrs := make(map[string]map[string]string)

	for rows.Next() {
		var productID, key, value string
		if err := rows.Scan(&productID, &key, &value); err != nil {
			return nil, fmt.Errorf("scan attribute failed: %w", err)
		}

		if _, exists := attrs[productID]; !exists {
			attrs[productID] = make(map[string]string)
		}
		attrs[productID][key] = value
	}

	return attrs, nil
}