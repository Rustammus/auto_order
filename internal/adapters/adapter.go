package adapters

import (
	"auto_order/internal/adapters/planet"
	"auto_order/internal/config"
	"auto_order/internal/models"
	"log"
)

type Supplier interface {
	SupplierID() int64
	SearchByCode(code string) (*models.SearchResult, error)
	SearchByText(text string) (*models.SearchResult, error)
	ListBasket() (*models.BasketList, error)
	AddProduct(prods *models.BasketAddProducts) (*models.BasketAddResult, error)
	RemoveProduct(items *models.BasketRemove) (*models.BasketList, error)
	Purchase() (*models.PurchaseResult, error)
	//
}

func NewSuppliers() []Supplier {
	kng := Supplier(nil)
	//p := planet.NewAPIAdapter()

	var p Supplier
	c := config.GetConfig()
	if c.MockAdapters == "true" {
		log.Println("mock adapters is true")
		p = planet.NewMockAPIAdapter()
	} else {
		p = planet.NewAPIAdapter()
	}

	progress := Supplier(nil)
	essco := Supplier(nil)
	return []Supplier{kng, p, progress, essco}
}
