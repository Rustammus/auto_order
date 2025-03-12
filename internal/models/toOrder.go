package models

import (
	"database/sql"
	"time"
)

type ToOrder struct {
	ID        int64          `json:"id"`
	ProductID int64          `json:"product_id"`
	SupID     int64          `json:"sup_id"`
	SupCode   sql.NullString `json:"sup_code"` // Указатель для nullable поля
	Count     int            `json:"count"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`

	// Опционально: связанные данные (можно добавить при необходимости)
	Product *Product `json:"product,omitempty"`
	//Supplier *Supplier `json:"supplier,omitempty"` // Если есть модель Supplier
}
