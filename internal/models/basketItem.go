package models

type BasketItem struct {
	ID                  int64   `json:"basket_item_id"`
	SupplierID          int64   `json:"supplier_id"`
	SupplierProductID   string  `json:"supplier_product_id"`
	ProductCode         string  `json:"product_code"`
	Article             string  `json:"article"`
	Brand               string  `json:"brand"`
	Title               string  `json:"title"`
	Price               float64 `json:"price"`
	Count               int64   `json:"count"`
	Comment             string  `json:"comment"`
	DeliveryAddressGUID string  `json:"delivery_address_guid"`
}

//
// =============================================================
//

type BasketAddProducts struct {
	Items []BasketItem `json:"items"`
}

type BasketAddResult struct {
	Items []BasketAddResultItem `json:"items"`
}

type BasketAddResultItem struct {
	ID                 int64   `json:"id"`
	ActualPrice        float64 `json:"actual_price"`
	Count              int     `json:"count"`
	Price              float64 `json:"price"`
	SupplierID         int64   `json:"supplier_id"`
	SupplierNic        string  `json:"supplier_nic"`
	StorageCount       int     `json:"storage_count"`
	StorageMeasureUnit string  `json:"storage_measure_unit"`
	ProductCode        string  `json:"nomenclature_product_code"`
	Title              string  `json:"nomenclature_title"`
	Brand              string  `json:"nomenclature_brand"`
	Article            string  `json:"nomenclature_article"`
	MultiBasketID      string  `json:"multi_basket_id"`
}

//
// =============================================================
//

type BasketRemove struct {
	Items []BasketRemoveItem
}

type BasketRemoveItem struct {
	BasketItemID int64
}

//
// =============================================================
//
