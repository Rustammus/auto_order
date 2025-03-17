package service

import (
	"auto_order/internal/models"
	"auto_order/internal/repo"
	"errors"
)

type Purchaser struct {
	sups *Suppliers
}

func NewPurchaser(r *repo.ProductRepo) *Searcher {
	return &Searcher{sups: NewSuppliers(), prodRepo: r}
}

func (p *Purchaser) PurchaseFromBasket(b *models.BasketList) (*models.PurchaseResult, error) {
	supID := b.SupplierID

	supService := p.sups.GetSupplierService(int(supID))
	if supService == nil {
		return nil, errors.New("supplier not found")
	}

	result, err := supService.Purchase()
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("result is nil")
	}
}
