package service

import (
	"auto_order/internal/models"
	"auto_order/internal/repo"
	"log"
	"sync"
)

type BasketsService struct {
	prodRepo *repo.ProductRepo
	sups     *Suppliers
	baskets  []*models.BasketList
	od       sync.Once
}

func NewBaskets(r *repo.ProductRepo) *BasketsService {
	return &BasketsService{prodRepo: r, baskets: make([]*models.BasketList, 0), sups: NewSuppliers()}
}

func (c *BasketsService) Basket(row int) *models.BasketList {
	if row >= len(c.baskets) {
		return nil
	}

	return c.baskets[row]
}

func (c *BasketsService) Size() (rows int) {

	c.od.Do(func() {
		bList, err := c.sups.GetSupplierService(1).ListBasket()
		if err != nil {
			log.Printf("Error getting supplier list: %v", err)
		} else {
			log.Printf("Supplier list: %v", bList)
		}

		c.baskets = append(c.baskets, bList)

		c.baskets = append(c.baskets, &models.BasketList{
			SupplierID: 0,
			TotalSum:   12056.50,
			TotalCount: 23,
			Items:      make([]models.BasketItem, 16),
		})

	})

	for i := 0; i < len(c.baskets); i++ {
		if c.baskets[i] != nil {
			rows++
		}
	}
	return
}

//func (c *BasketsService) SearchByCode(code string) {
//	result, err := c.prodRepo.FindBySupplierCodeAny(context.TODO(), code)
//	if err != nil {
//		log.Print(err)
//	}
//	c.search = result
//}
//
//func (c *BasketsService) SearchByText(text string) {
//	result, err := c.prodRepo.FindByName(context.TODO(), text)
//	if err != nil {
//		log.Print(err)
//	}
//	c.search = result
//}
