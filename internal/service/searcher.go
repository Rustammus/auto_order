package service

import (
	"auto_order/internal/models"
	"auto_order/internal/repo"
	"auto_order/internal/sups"
	"log"
)

type Searcher struct {
	sups     *Suppliers
	result   *models.SearchResult
	prodRepo *repo.ProductRepo
}

func NewSearcher(r *repo.ProductRepo) *Searcher {
	return &Searcher{sups: NewSuppliers(), prodRepo: r}
}

func (s *Searcher) Size() int {
	if s.result != nil && s.result.SearchItems != nil {
		return len(s.result.SearchItems)
	}
	return 0
}

func (s *Searcher) Item(itemID int) models.SearchItem {
	if s.result != nil && s.result.SearchItems != nil {
		if len(s.result.SearchItems) > itemID {
			return s.result.SearchItems[itemID]
		}
	}

	return models.SearchItem{
		SupplierID:        -1,
		SupplierProductID: "RANGE",
		Article:           "RANGE",
		Brand:             "RANGE",
		Title:             "RANGE",
		Price:             -1,
		Count:             -1,
	}
}

// Search Run search method on
func (s *Searcher) Search(text string, searchType int, supFlag int) {
	if supFlag&sups.KngMask != 0 {
		_ = s.search(text, searchType, 0)
	}
	if supFlag&sups.PlanetMask != 0 {
		r1 := s.search(text, searchType, 1)
		s.result = r1
	}
	if supFlag&sups.ProgressMask != 0 {
		_ = s.search(text, searchType, 2)
	}
	if supFlag&sups.EsscoMask != 0 {
		_ = s.search(text, searchType, 3)
	}

	// catalog insertion
	//for _, i := range s.result.SearchItems {
	//	p, err := s.prodRepo.FindByNumberP(context.Background(), i.SupplierID)
	//	if err != nil {
	//		log.Println("error on FindByNumberP", err)
	//		continue
	//	}
	//	if len(p) > 1 {
	//		log.Println("error on FindByNumberP: len > 1 got: ", len(p))
	//		continue
	//	}
	//
	//	if len(p) > 0 {
	//		el := p[0]
	//		if el.PlanetID == i.SupplierID {
	//			continue
	//		}
	//	}
	//
	//	r := rand.Intn(2)
	//	kngCode := int64(0)
	//
	//	if r > 0 {
	//		kngCode = int64(rand.Intn(139000))
	//	}
	//
	//	t := time.Now()
	//	_, err = s.prodRepo.Create(context.Background(), &models.Product{
	//		KngID:       kngCode,
	//		PlanetID:    i.SupplierID,
	//		ProgressID:  0,
	//		EsscoID:     0,
	//		Article:     i.Article,
	//		Name:        i.Title,
	//		Description: "",
	//		CreatedAt:   t,
	//		UpdatedAt:   t,
	//	})
	//	if err != nil {
	//		log.Println("error on Create", err)
	//	}
	//}
	// catalog insertion
}

func (s *Searcher) search(text string, searchType int, supplier int) *models.SearchResult {
	if s.sups == nil && s.sups.suppliers[supplier] == nil {
		log.Print("service.Searcher: search: got null supplier")
		return nil
	}

	result, err := s.sups.suppliers[supplier].SearchByText(text)
	if err != nil {
		log.Print("service.Searcher: search: error on search: ", err)
		return nil
	}
	log.Printf("service.Searcher: search: supplier %v, result rows %v", supplier, len(result.SearchItems))

	return result
}
