package models

type PurchaseResult struct {
	Success    bool    `json:"success"`
	TotalPrice float64 `json:"total_price"`
	Items      []BasketAddResultItem
}
