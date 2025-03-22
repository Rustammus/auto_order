package repo

import (
	"auto_order/internal/models"
	"context"
	"database/sql"
	"time"
)

type ToOrderRepo struct {
	db *sql.DB
}

func NewToOrderRepo(db *sql.DB) *ToOrderRepo {
	return &ToOrderRepo{db: db}
}

func (r *ToOrderRepo) Create(ctx context.Context, order *models.ToOrder) (int64, error) {
	query := `
		INSERT INTO to_order (
			product_id, sup_id, sup_code, count, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?)
		RETURNING id
	`

	var id int64
	err := r.db.QueryRowContext(ctx, query,
		order.ProductID,
		order.SupID,
		order.SupCode,
		order.Count,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ToOrderRepo) Find(ctx context.Context, id int64) (*models.ToOrder, error) {
	query := `
		SELECT 
			id, product_id, sup_id, sup_code, count, created_at, updated_at
		FROM to_order
		WHERE id = ?
	`

	var order models.ToOrder
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&order.ID,
		&order.ProductID,
		&order.SupID,
		&order.SupCode,
		&order.Count,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &order, nil
}

func (r *ToOrderRepo) List(ctx context.Context) ([]models.ToOrder, error) {
	query := `
		SELECT 
			id, product_id, sup_id, sup_code, count, created_at, updated_at
		FROM to_order
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.ToOrder
	for rows.Next() {
		var order models.ToOrder
		err := rows.Scan(
			&order.ID,
			&order.ProductID,
			&order.SupID,
			&order.SupCode,
			&order.Count,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *ToOrderRepo) FindByProductAndSup(ctx context.Context, productID, supID int64) (*models.ToOrder, error) {
	query := `
		SELECT 
			id, product_id, sup_id, sup_code, count, created_at, updated_at
		FROM to_order
		WHERE product_id = ? AND sup_id = ?
	`

	var order models.ToOrder
	err := r.db.QueryRowContext(ctx, query, productID, supID).Scan(
		&order.ID,
		&order.ProductID,
		&order.SupID,
		&order.SupCode,
		&order.Count,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &order, nil
}

func (r *ToOrderRepo) FindByProductID(ctx context.Context, productID int64) ([]models.ToOrder, error) {
	query := `
		SELECT 
			id, product_id, sup_id, sup_code, count, created_at, updated_at
		FROM to_order
		WHERE product_id = ?
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.ToOrder
	for rows.Next() {
		var order models.ToOrder
		err := rows.Scan(
			&order.ID,
			&order.ProductID,
			&order.SupID,
			&order.SupCode,
			&order.Count,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *ToOrderRepo) FindBySupplierID(ctx context.Context, supID int64) ([]models.ToOrder, error) {
	query := `
		SELECT 
			id, product_id, sup_id, sup_code, count, created_at, updated_at
		FROM to_order
		WHERE sup_id = ?
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, supID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.ToOrder
	for rows.Next() {
		var order models.ToOrder
		err := rows.Scan(
			&order.ID,
			&order.ProductID,
			&order.SupID,
			&order.SupCode,
			&order.Count,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *ToOrderRepo) Update(ctx context.Context, order *models.ToOrder) error {
	query := `
		UPDATE to_order SET
			product_id = ?,
			sup_id = ?,
			sup_code = ?,
			count = ?,
			updated_at = ?
		WHERE id = ?
	`

	_, err := r.db.ExecContext(ctx, query,
		order.ProductID,
		order.SupID,
		order.SupCode,
		order.Count,
		time.Now(),
		order.ID,
	)

	return err
}

func (r *ToOrderRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM to_order WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// GetOrderWithProduct возвращает заказ с информацией о продукте по ID заказа
func (r *ToOrderRepo) GetOrderWithProduct(ctx context.Context, orderID int64) (*models.ToOrder, error) {
	query := `
        SELECT 
            o.id, o.product_id, o.sup_id, o.sup_code, o.count, 
            o.created_at AS order_created_at, o.updated_at AS order_updated_at,
            p.id AS product_id, p.article, p.name, p.description,
            p.created_at AS product_created_at, p.updated_at AS product_updated_at
        FROM to_order o
        JOIN products p ON o.product_id = p.id
        WHERE o.id = ?
    `

	var result models.ToOrder
	result.Product = &models.Product{}

	err := r.db.QueryRowContext(ctx, query, orderID).Scan(
		&result.ID,
		&result.ProductID,
		&result.SupID,
		&result.SupCode,
		&result.Count,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.Product.ID,
		&result.Product.Article,
		&result.Product.Name,
		&result.Product.Description,
		&result.Product.CreatedAt,
		&result.Product.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

// ListOrdersWithProducts возвращает список всех заказов с информацией о продуктах
func (r *ToOrderRepo) ListOrdersWithProducts(ctx context.Context) ([]models.ToOrder, error) {
	query := `
        SELECT 
            o.id, o.product_id, o.sup_id, o.sup_code, o.count, 
            o.created_at AS order_created_at, o.updated_at AS order_updated_at,
            p.id AS product_id, p.article, p.name, p.description,
            p.created_at AS product_created_at, p.updated_at AS product_updated_at
        FROM to_order o
        JOIN products p ON o.product_id = p.id
        ORDER BY o.created_at DESC
    `

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.ToOrder
	for rows.Next() {
		var order models.ToOrder
		order.Product = &models.Product{}
		err := rows.Scan(
			&order.ID,
			&order.ProductID,
			&order.SupID,
			&order.SupCode,
			&order.Count,
			&order.CreatedAt,
			&order.UpdatedAt,
			&order.Product.ID,
			&order.Product.Article,
			&order.Product.Name,
			&order.Product.Description,
			&order.Product.CreatedAt,
			&order.Product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
