package planet

import (
	"auto_order/internal/models"
	"auto_order/internal/sups"
)

// Basket представляет корзину товаров
type Basket struct {
	Items      []BasketItem `json:"items"`
	TotalSum   float64      `json:"total_sum"`
	TotalCount int64        `json:"total_count"`
}

func (b *Basket) ToDTO() *models.BasketList {
	dto := &models.BasketList{}

	dto.TotalSum = b.TotalSum
	dto.TotalCount = b.TotalCount

	for _, item := range b.Items {
		i := models.BasketItem{}

		i.ID = item.ID
		i.ProductCode = item.ProductCode
		i.SupplierID = sups.PlanetID
		i.SupplierProductID = item.ProductID
		i.Article = item.Article
		i.Brand = item.Brand
		i.Title = item.Title
		i.Price = item.Price
		i.Count = item.Count
		i.Comment = item.Comment
		i.DeliveryAddressGUID = item.DeliveryAddressGUID

		dto.Items = append(dto.Items, i)
	}

	return dto
}

// BasketItem представляет позицию в корзине
type BasketItem struct {
	ID                         int64    `json:"id"`
	Comment                    string   `json:"comment"`
	Count                      int64    `json:"count"`
	Price                      float64  `json:"price"`
	DeliveryAddressGUID        string   `json:"delivery_address_guid"`
	ProductID                  string   `json:"product_id"`
	Title                      string   `json:"title"`
	Article                    string   `json:"article"`
	Brand                      string   `json:"brand"`
	ProductCode                string   `json:"product_code"`
	SupplierID                 int64    `json:"supplier_id"`
	IsExpressDeliveryAvailable bool     `json:"is_express_delivery_available"`
	ExpressDeliveryCost        int64    `json:"express_delivery_cost"`
	Images                     []string `json:"images"`
	CreatedAt                  int64    `json:"created_at"`
	UpdatedAt                  int64    `json:"updated_at"`
}

type BasketItemShort struct {
	ID          *int64 `json:"id,omitempty"`
	Count       *int64 `json:"count,omitempty"`
	ProductID   string `json:"product_id"`
	ProductCode string `json:"product_code"`
	SupplierID  int64  `json:"supplier_id"`
}

//
// =============================================================
//

// SummaryResponseModel представляет сводную информацию о корзине
type SummaryResponseModel struct {
	TotalCount float64 `json:"total_count"`
	TotalSum   float64 `json:"total_sum"`
	RowsCount  string  `json:"rows_count"`
}

//
// =============================================================
//

type BasketRemoveRequest struct {
	BasketIDs         []int64 `json:"basket_ids"`
	isExpressDelivery bool
}

func (b *BasketRemoveRequest) FromDTO(items *models.BasketRemove) {
	ids := make([]int64, 0, len(items.Items))

	for _, i := range items.Items {

		var id int64

		id = i.BasketItemID

		ids = append(ids, id)
	}

}

//
// =============================================================
//

// BasketItems представляет запрос для работы с позициями корзины
type BasketItems struct {
	BasketIDs         []int64 `json:"basket_ids"`
	IsExpressDelivery bool    `json:"is_express_delivery"`
}
