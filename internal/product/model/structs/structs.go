package model

// Products структура данных товара
// @Description Основная информация о товаре в системе
type Products struct {
	Product_id  string              `json:"product_id" example:"PRD-1001"`                        // Уникальный идентификатор товара
	Name        string              `json:"name" example:"Смартфон"`                              // Наименование товара
	Description string              `json:"description" example:"Флагманский смартфон 2023 года"` // Описание товара
	KeyValue    []map[string]string `json:"keyvalue,omitempty"`                                   // Характеристики товара (ключ-значение)
	Weight      string              `json:"weight,omitempty" example:"0.2"`                       // Вес товара
	Barcode     string              `json:"barcode,omitempty" example:"123456789012"`             // Штрих-код товара
}
