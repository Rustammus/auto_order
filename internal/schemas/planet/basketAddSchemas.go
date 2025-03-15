package planet

import "auto_order/internal/models"

//
// =============================================================
//

// BasketAdd представляет запрос на добавление товаров в корзину
type BasketAddRequest struct {
	Items []BasketAddRequestItem `json:"items"`
}

// BasketAddItem представляет позицию для добавления в корзину
type BasketAddRequestItem struct {
	ProductID   *string `json:"product_id,omitempty"`
	ProductCode *string `json:"product_code,omitempty"`
	SupplierID  int64   `json:"supplier_id"`
	Count       float64 `json:"count"`
}

func (b *BasketAddRequest) FromDTO(items *models.BasketAddProducts) {
	reqItems := make([]BasketAddRequestItem, 0, len(items.Items))

	for _, item := range items.Items {
		pID := ""
		pCode := ""

		var reqItem BasketAddRequestItem
		reqItem.SupplierID = item.SupplierID
		reqItem.Count = float64(item.Count)

		if item.ProductCode != "" {
			pCode = item.ProductCode
			reqItem.ProductCode = &pCode
		}

		if item.SupplierProductID != "" {
			pID = item.SupplierProductID
			reqItem.ProductID = &pID
		}
		reqItems = append(reqItems, reqItem)
	}

}

//
// =============================================================
//

type BasketAddResponse struct {
	Items []BasketAddResponseItem `json:"items"`
}

type BasketAddResponseItem struct {
	ID                      int64   `json:"id"`
	ActualPrice             float64 `json:"actual_price"`
	Count                   int     `json:"count"`
	Price                   float64 `json:"price"`
	SupplierAlias           string  `json:"supplier_alias"`
	SupplierID              int64   `json:"supplier_id"`
	SupplierNic             string  `json:"supplier_nic"`
	StorageCount            int     `json:"storage_count"`
	StorageMeasureUnit      string  `json:"storage_measure_unit"`
	NomenclatureProductCode string  `json:"nomenclature_product_code"`
	NomenclatureTitle       string  `json:"nomenclature_title"`
	NomenclatureBrand       string  `json:"nomenclature_brand"`
	NomenclatureArticle     string  `json:"nomenclature_article"`
	MultiBasketID           string  `json:"multi_basket_id"`
}

func (b *BasketAddResponse) ToDTO() *models.BasketAddResult {
	var dto models.BasketAddResult
	dtoItems := make([]models.BasketAddResultItem, 0, len(b.Items))

	for _, i := range b.Items {
		dtoItems = append(dtoItems, models.BasketAddResultItem{
			ID:                 i.ID,
			ActualPrice:        i.ActualPrice,
			Count:              i.Count,
			Price:              i.Price,
			SupplierID:         i.SupplierID,
			SupplierNic:        i.SupplierNic,
			StorageCount:       i.StorageCount,
			StorageMeasureUnit: i.StorageMeasureUnit,
			ProductCode:        i.NomenclatureProductCode,
			Title:              i.NomenclatureTitle,
			Brand:              i.NomenclatureBrand,
			Article:            i.NomenclatureArticle,
			MultiBasketID:      i.MultiBasketID,
		})
	}

	dto.Items = dtoItems
	return &dto
}
