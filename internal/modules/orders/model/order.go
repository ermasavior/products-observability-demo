package model

type Order struct {
	ID        string `db:"id" json:"id"`
	ProductID int64  `db:"product_id" json:"product_id"`
	Total     int64  `db:"total" json:"total"`
	CreatedAt string `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt string `db:"updated_at" json:"updated_at,omitempty"`
}
