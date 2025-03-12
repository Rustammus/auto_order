package models

type SearchItem struct {
	SupplierID        int64   `json:"supplier_id"`
	SupplierProductID string  `json:"supplier_product_id"`
	Code              string  `json:"code"`
	Article           string  `json:"article"`
	Brand             string  `json:"brand"`
	Title             string  `json:"title"`
	Price             float64 `json:"price"`
	Count             int64   `json:"count"`
	QuantityInPack    int64   `json:"quantity_in_pack"`
	MeasureRatio      int64   `json:"measure_ratio"`
	MeasureUnit       string  `json:"measure_unit"`
}

type SearchResult struct {
	Page           int64        `json:"page"`
	TotalPages     int64        `json:"total_pages"`
	ProductsOnPage int64        `json:"products_on_page"`
	ProductsTotal  int64        `json:"products_total_count"`
	SearchItems    []SearchItem `json:"search_items"`
}
