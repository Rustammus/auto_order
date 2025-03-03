package repo

import (
	"auto_order/internal/models"
	"context"
	"database/sql"
	"errors"
	"time"
)

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) Create(ctx context.Context, product *models.Product) (int64, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	// Создаем основной продукт
	productQuery := `
		INSERT INTO products (
			article, name, description, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?)
		RETURNING id
	`

	var productID int64
	err = tx.QueryRowContext(ctx, productQuery,
		product.Article,
		product.Name,
		product.Description,
		time.Now(),
		time.Now(),
	).Scan(&productID)
	if err != nil {
		return 0, err
	}

	// Создаем коды поставщиков
	if err := r.createProductCodes(ctx, tx, productID, product); err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return productID, nil
}

func (r *ProductRepo) createProductCodes(ctx context.Context, tx *sql.Tx, productID int64, product *models.Product) error {
	codes := map[string]*models.ProductCode{
		"kng":      product.KngCode,
		"planet":   product.PlanetCode,
		"progress": product.ProgressCode,
		"essco":    product.EsscoCode,
	}

	codeQuery := `
		INSERT INTO product_suppliers (
			product_id, sup_id, sup_product_id, sup_code, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?)
	`

	for _, code := range codes {
		if code == nil {
			continue
		}

		_, err := tx.ExecContext(ctx, codeQuery,
			productID,
			code.SupID,
			code.SupProductID,
			code.SupCode,
			time.Now(),
			time.Now(),
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *ProductRepo) Find(ctx context.Context, id int64) (*models.Product, error) {
	// Загружаем основной продукт
	productQuery := `
		SELECT 
			id, article, name, description, created_at, updated_at
		FROM products
		WHERE id = ?
	`

	var product models.Product
	err := r.db.QueryRowContext(ctx, productQuery, id).Scan(
		&product.ID,
		&product.Article,
		&product.Name,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	// Загружаем коды поставщиков
	if err := r.loadProductCodes(ctx, &product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepo) loadProductCodes(ctx context.Context, product *models.Product) error {
	codesQuery := `
		SELECT 
			id, product_id, sup_id, sup_product_id, sup_code, created_at, updated_at
		FROM product_suppliers
		WHERE product_id = ?
	`

	rows, err := r.db.QueryContext(ctx, codesQuery, product.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	supplierCodes := make(map[int64]*models.ProductCode)
	for rows.Next() {
		var code models.ProductCode
		err := rows.Scan(
			&code.ID,
			&code.ProductID,
			&code.SupID,
			&code.SupProductID,
			&code.SupCode,
			&code.CreatedAt,
			&code.UpdatedAt,
		)
		if err != nil {
			return err
		}
		supplierCodes[code.SupID] = &code
	}

	if err = rows.Err(); err != nil {
		return err
	}

	// Предполагаем, что мы знаем ID поставщиков для каждого типа кода
	// Это можно вынести в конфигурацию или отдельную таблицу
	product.KngCode = supplierCodes[0]      // например, KNG имеет sup_id = 1
	product.PlanetCode = supplierCodes[1]   // Planet имеет sup_id = 2
	product.ProgressCode = supplierCodes[2] // Progress имеет sup_id = 3
	product.EsscoCode = supplierCodes[3]    // Essco имеет sup_id = 4

	return nil
}

func (r *ProductRepo) List(ctx context.Context) ([]models.Product, error) {
	productQuery := `
		SELECT 
			id, article, name, description, created_at, updated_at
		FROM products
		ORDER BY name
	`

	rows, err := r.db.QueryContext(ctx, productQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
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

	// Загружаем коды поставщиков для каждого продукта
	for i := range products {
		if err := r.loadProductCodes(ctx, &products[i]); err != nil {
			return nil, err
		}
	}

	return products, nil
}

func (r *ProductRepo) FindByArticle(ctx context.Context, article string) ([]models.Product, error) {
	query := `
		SELECT 
			id, article, name, description, created_at, updated_at
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

	// Загружаем коды поставщиков
	for i := range products {
		if err := r.loadProductCodes(ctx, &products[i]); err != nil {
			return nil, err
		}
	}

	return products, nil
}

func (r *ProductRepo) FindByName(ctx context.Context, name string) ([]models.Product, error) {
	query := `
		SELECT 
			id, article, name, description, created_at, updated_at
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

	// Загружаем коды поставщиков
	for i := range products {
		if err := r.loadProductCodes(ctx, &products[i]); err != nil {
			return nil, err
		}
	}

	return products, nil
}

func (r *ProductRepo) FindBySupplierCode(ctx context.Context, supID int64, code string) (*models.Product, error) {
	query := `
		SELECT 
			p.id, p.article, p.name, p.description, p.created_at, p.updated_at
		FROM products p
		JOIN product_suppliers ps ON p.id = ps.product_id
		WHERE ps.sup_id = ? AND ps.sup_code = ?
		LIMIT 1
	`

	var product models.Product
	err := r.db.QueryRowContext(ctx, query, supID, code).Scan(
		&product.ID,
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

	// Загружаем коды поставщиков
	if err := r.loadProductCodes(ctx, &product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepo) FindBySupplierCodeAny(ctx context.Context, code string) ([]models.Product, error) {
	query := `
        SELECT 
            p.id, p.article, p.name, p.description, p.created_at, p.updated_at
        FROM products p
        JOIN product_suppliers ps ON p.id = ps.product_id
        WHERE ps.sup_code = ?
        ORDER BY p.name
    `

	rows, err := r.db.QueryContext(ctx, query, code)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
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

	// Загружаем коды поставщиков для найденных продуктов
	for i := range products {
		if err := r.loadProductCodes(ctx, &products[i]); err != nil {
			return nil, err
		}
	}

	return products, nil
}

func (r *ProductRepo) Update(ctx context.Context, product *models.Product) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Обновляем основной продукт
	productQuery := `
		UPDATE products SET
			article = ?,
			name = ?,
			description = ?,
			updated_at = ?
		WHERE id = ?
	`

	_, err = tx.ExecContext(ctx, productQuery,
		product.Article,
		product.Name,
		product.Description,
		time.Now(),
		product.ID,
	)
	if err != nil {
		return err
	}

	// Удаляем старые коды поставщиков
	_, err = tx.ExecContext(ctx, "DELETE FROM product_suppliers WHERE product_id = ?", product.ID)
	if err != nil {
		return err
	}

	// Добавляем новые коды поставщиков
	if err := r.createProductCodes(ctx, tx, product.ID, product); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *ProductRepo) Delete(ctx context.Context, id int64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Сначала удаляем коды поставщиков
	_, err = tx.ExecContext(ctx, "DELETE FROM product_suppliers WHERE product_id = ?", id)
	if err != nil {
		return err
	}

	// Затем удаляем сам продукт
	_, err = tx.ExecContext(ctx, "DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
