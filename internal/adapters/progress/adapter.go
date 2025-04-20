package kng

import (
	"auto_order/internal/models"
)

func (a *APIAdapter) SupplierID() int64 {
	return supplierID
}

func (a *APIAdapter) SearchByCode(code string) (*models.SearchResult, error) {
	// TODO errors
	panic("implement me")
}

func (a *APIAdapter) SearchByText(text string) (*models.SearchResult, error) {
	// TODO errors
	panic("implement me")
}

func (a *APIAdapter) ListBasket() (*models.BasketList, error) {
	// TODO errors

	panic("implement me")
}

func (a *APIAdapter) AddProduct(prods *models.BasketAddProducts) (*models.BasketAddResult, error) {
	// TODO errors
	panic("implement me")
}

func (a *APIAdapter) RemoveProduct(items *models.BasketRemove) (*models.BasketList, error) {
	// TODO errors
	panic("implement me")
}

func (a *APIAdapter) Purchase() (*models.PurchaseResult, error) {
	return nil, nil
}
