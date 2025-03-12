package service

import (
	"auto_order/internal/models"
	"auto_order/internal/repo"
	"context"
	"log"
)

type Catalog struct {
	prodRepo *repo.ProductRepo
	search   []models.Product
}

func NewCatalog(r *repo.ProductRepo) *Catalog {
	return &Catalog{prodRepo: r, search: make([]models.Product, 0)}
}

func (c *Catalog) ColumnName(col int) (name string) {
	switch col {
	case 0:
		name = "КНГ"
	case 1:
		name = "ПЛАНЕТА"
	case 2:
		name = "Прогресс"
	case 3:
		name = "ЭССКО"
	case 4:
		name = "Название"
	case 5:
		name = "Артикул"
	default:
		name = "Undefined"
	}
	return
}

func (c *Catalog) Cell(row, col int) (cell string) {
	if row >= len(c.search) {
		return "RANGE"
	}

	prod := c.search[row]

	switch col {
	case 0:
		cell = prod.KngIDVal()
	case 1:
		cell = prod.PlanetIDVal()
	case 2:
		cell = prod.ProgressIDVal()
	case 3:
		cell = prod.EsscoIDVal()
	case 4:
		cell = prod.Name
	case 5:
		cell = prod.Article.String
	default:
		cell = "Undefined"
	}
	return
}

func (c *Catalog) Size() (rows int, cols int) {
	rows = len(c.search)
	if rows != 0 {
		cols = 6
	}

	return
}

func (c *Catalog) SearchByCode(code string) {
	result, err := c.prodRepo.FindBySupplierCodeAny(context.TODO(), code)
	if err != nil {
		log.Print(err)
	}
	c.search = result
}

func (c *Catalog) SearchByText(text string) {
	result, err := c.prodRepo.FindByName(context.TODO(), text)
	if err != nil {
		log.Print(err)
	}
	c.search = result
}

func (c *Catalog) GetProduct(row int) *models.Product {
	if row >= len(c.search) {
		log.Printf("service.Catalog.GetProduct row out of range")
		return nil
	}

	prod := c.search[row]

	return &prod
}

func (c *Catalog) ListAll() {
	result, err := c.prodRepo.List(context.TODO())
	if err != nil {
		log.Print(err)
	}
	c.search = result
}
