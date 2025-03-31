package _repo

import (
	"auto_order/internal/models"
	"context"
	"database/sql"
	"time"
)

type ProductRepo2 struct {
	db *sql.DB
}

func NewProductRepo2(db *sql.DB) *ProductRepo2 {
	return &ProductRepo2{db: db}
}

func (r *ProductRepo2) Create(ctx context.Context, product *models.Product) (int64, error) {
	query := `
		INSERT INTO products (
			kng_id, planet_id, progress_id, essco_id, 
			article, name, description, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		RETURNING id
	`

	var id int64
	err := r.db.QueryRowContext(ctx, query,
		product.KngID,
		product.PlanetID,
		product.ProgressID,
		product.EsscoID,
		product.Article,
		product.Name,
		product.Description,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ProductRepo2) Find(ctx context.Context, id int64) (*models.Product, error) {
	query := `
		SELECT 
			id, kng_id, planet_id, progress_id, essco_id,
			article, name, description, created_at, updated_at
		FROM products
		WHERE id = ?
	`

	var product models.Product
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&product.ID,
		&product.KngID,
		&product.PlanetID,
		&product.ProgressID,
		&product.EsscoID,
		&product.Article,
		&product.Name,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepo2) List(ctx context.Context) ([]models.Product, error) {
	query := `
		SELECT 
			id, kng_id, planet_id, progress_id, essco_id,
			article, name, description, created_at, updated_at
		FROM products
		ORDER BY name
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.KngID,
			&product.PlanetID,
			&product.ProgressID,
			&product.EsscoID,
			&product.Article,
			&product.Name,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepo2) FindByArticle(ctx context.Context, article string) ([]models.Product, error) {
	query := `
		SELECT 
			id, kng_id, planet_id, progress_id, essco_id,
			article, name, description, created_at, updated_at
		FROM products
		WHERE article LIKE ?
		ORDER BY name
	`

	rows, err := r.db.QueryContext(ctx, query, "%"+article+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.KngID,
			&product.PlanetID,
			&product.ProgressID,
			&product.EsscoID,
			&product.Article,
			&product.Name,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepo2) FindByName(ctx context.Context, name string) ([]models.Product, error) {
	query := `
		SELECT 
			id, kng_id, planet_id, progress_id, essco_id,
			article, name, description, created_at, updated_at
		FROM products
		WHERE name LIKE ?
		ORDER BY name
	`

	rows, err := r.db.QueryContext(ctx, query, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.KngID,
			&product.PlanetID,
			&product.ProgressID,
			&product.EsscoID,
			&product.Article,
			&product.Name,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// FindByNumber ищет продукты по числовому значению во всех int64 полях
// (kng_id, planet_id, progress_id, essco_id, id)
func (r *ProductRepo2) FindByNumber(ctx context.Context, number int64) ([]models.Product, error) {
	query := `
        SELECT 
            id, kng_id, planet_id, progress_id, essco_id,
            article, name, description, created_at, updated_at
        FROM products
        WHERE 
            kng_id = ? OR
            planet_id = ? OR
            progress_id = ? OR
            essco_id = ?
        ORDER BY name
    `

	// Передаем одно и то же число для всех полей
	rows, err := r.db.QueryContext(ctx, query, number, number, number, number)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.KngID,
			&product.PlanetID,
			&product.ProgressID,
			&product.EsscoID,
			&product.Article,
			&product.Name,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// FindByNumberP ищет продукты по числовому значению во всех int64 полях
// (kng_id, planet_id, progress_id, essco_id, id)
func (r *ProductRepo2) FindByNumberP(ctx context.Context, number int64) ([]models.Product, error) {
	query := `
        SELECT 
            id, kng_id, planet_id, progress_id, essco_id,
            article, name, description, created_at, updated_at
        FROM products
        WHERE 
            planet_id = ?
        ORDER BY name
    `

	// Передаем одно и то же число для всех полей
	rows, err := r.db.QueryContext(ctx, query, number)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.KngID,
			&product.PlanetID,
			&product.ProgressID,
			&product.EsscoID,
			&product.Article,
			&product.Name,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepo2) Search(ctx context.Context, key string) ([]models.Product, error) {
	query := `
		SELECT 
			id, kng_id, planet_id, progress_id, essco_id,
			article, name, description, created_at, updated_at
		FROM products
		WHERE 
			article LIKE ? OR 
			name LIKE ? OR 
			description LIKE ?
		ORDER BY name
	`

	searchKey := "%" + key + "%"
	rows, err := r.db.QueryContext(ctx, query, searchKey, searchKey, searchKey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.KngID,
			&product.PlanetID,
			&product.ProgressID,
			&product.EsscoID,
			&product.Article,
			&product.Name,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepo2) Update(ctx context.Context, product *models.Product) error {
	query := `
		UPDATE products SET
			kng_id = ?,
			planet_id = ?,
			progress_id = ?,
			essco_id = ?,
			article = ?,
			name = ?,
			description = ?,
			updated_at = ?
		WHERE id = ?
	`

	_, err := r.db.ExecContext(ctx, query,
		product.KngID,
		product.PlanetID,
		product.ProgressID,
		product.EsscoID,
		product.Article,
		product.Name,
		product.Description,
		time.Now(),
		product.ID,
	)

	return err
}

func (r *ProductRepo2) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM products WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
