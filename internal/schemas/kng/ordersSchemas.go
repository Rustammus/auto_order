package planet

// ApiOrder представляет заказ
type ApiOrder struct {
	Orders []ApiOrderItem `json:"orders"`
}

// ApiOrderItem представляет позицию в заказе
type ApiOrderItem struct {
	ID           int64     `json:"id"`
	SupplierCode string    `json:"supplier_code"`
	SupplierID   int64     `json:"supplier_id"`
	ProductID    string    `json:"product_id"`
	Title        string    `json:"title"`
	Brand        string    `json:"brand"`
	Article      string    `json:"article"`
	Count        float64   `json:"count"`
	Price        float64   `json:"price"`
	Status       ApiStatus `json:"status"`
	DeliveryDate string    `json:"delivery_date"`
	Comment      string    `json:"comment"`
}

// ApiStatus представляет статус заказа
type ApiStatus struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}
