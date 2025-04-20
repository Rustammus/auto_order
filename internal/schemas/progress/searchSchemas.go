package planet

import (
	"auto_order/internal/models"
	"auto_order/internal/sups"
)

// SearchResponse представляет результат поиска товаров
type SearchResponse struct {
	Page           int64     `json:"page"`
	TotalPages     int64     `json:"total_pages"`
	ProductsOnPage int64     `json:"products_on_page"`
	ProductsTotal  int64     `json:"products_total_count"`
	Products       []Product `json:"products"`
}

func (r *SearchResponse) ToDTO() *models.SearchResult {
	dto := models.SearchResult{}

	for _, p := range r.Products {
		var i models.SearchItem

		i.SupplierID = sups.PlanetID
		i.SupplierProductID = p.ID
		i.Code = p.ProductCode
		i.Article = p.Article
		i.Brand = p.Brand
		i.Title = p.Title
		i.Price = p.Price
		i.Count = p.Count

		dto.SearchItems = append(dto.SearchItems, i)
	}
	return &dto
}
