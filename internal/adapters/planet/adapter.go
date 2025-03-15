package planet

import (
	"auto_order/internal/models"
)

func (a *APIAdapter) SupplierID() int64 {
	return supplierID
}

func (a *APIAdapter) SearchByCode(code string) (*models.SearchResult, error) {
	// TODO errors
	return a.searchByCode(code)
}

func (a *APIAdapter) SearchByText(text string) (*models.SearchResult, error) {
	// TODO errors
	return a.searchByText(text)
}

func (a *APIAdapter) ListBasket() (*models.BasketList, error) {
	// TODO errors

	list, err := a.listBasket()
	if err != nil {
		panic(err)
	}
	list.SupplierID = supplierID
	return list, err
}

func (a *APIAdapter) AddProduct(prods *models.BasketAddProducts) (*models.BasketAddResult, error) {
	// TODO errors
	return a.addToBasket(prods)
}

func (a *APIAdapter) RemoveProduct(items *models.BasketRemove) (*models.BasketList, error) {
	// TODO errors
	list, err := a.removeFromBasket(items)
	list.SupplierID = supplierID
	return list, err
}

func (a *APIAdapter) Purchase() (*models.PurchaseResult, error) {
	return nil, nil
}
