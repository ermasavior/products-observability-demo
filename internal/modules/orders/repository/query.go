package repository

const (
	getProductStockByIDQuery = `SELECT stock FROM products WHERE id = $1`

	insertOrderQuery = `
		INSERT INTO orders (product_id, total, created_at)
		VALUES (:product_id, :total, :created_at)
	    RETURNING id
	`

	decreaseProductStockQuery = `
		UPDATE products
		SET stock = stock - $1, updated_at = $2
		WHERE id = $3 RETURNING id
	`
)
