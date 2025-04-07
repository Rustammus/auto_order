package kng

import (
	"auto_order/internal/models"
)

// MockAPIAdapter реализует тот же интерфейс, что и APIAdapter
type MockAPIAdapter struct {
	Token          string
	BaseURL        string
	SearchByCodeFn func(code string) (*models.SearchResult, error)
	SearchByTextFn func(text string) (*models.SearchResult, error)
	ListBasketFn   func() (*models.BasketList, error)
}

func (m *MockAPIAdapter) AddProduct(prods *models.BasketAddProducts) (*models.BasketAddResult, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockAPIAdapter) RemoveProduct(items *models.BasketRemove) (*models.BasketList, error) {
	//TODO implement me
	panic("implement me")
}

func NewMockAPIAdapter() *MockAPIAdapter {
	return &MockAPIAdapter{
		Token:   "mock-token",
		BaseURL: "http://mock-api",
	}
}

func (m *MockAPIAdapter) SupplierID() int64 {
	return 1 // supplierID
}

func (m *MockAPIAdapter) SearchByCode(code string) (*models.SearchResult, error) {
	if m.SearchByCodeFn != nil {
		return m.SearchByCodeFn(code)
	}
	// Стандартный моковый ответ
	return &models.SearchResult{
		SearchItems: []models.SearchItem{
			{
				SupplierID:        1,
				SupplierProductID: "mock-" + code,
				Article:           "ART-" + code,
				Brand:             "MockBrand",
				Title:             "Mock Product " + code,
				Price:             100.0,
				Count:             10,
			},
		},
	}, nil
}

func (m *MockAPIAdapter) SearchByText(text string) (*models.SearchResult, error) {
	if m.SearchByTextFn != nil {
		return m.SearchByTextFn(text)
	}
	// Стандартный моковый ответ
	return &models.SearchResult{
		SearchItems: []models.SearchItem{
			{
				SupplierID:        1,
				SupplierProductID: "mock-" + text,
				Article:           "ART-" + text,
				Brand:             "MockBrand",
				Title:             "Mock Product for " + text,
				Price:             150.0,
				Count:             5,
			},
			{
				SupplierID:        1,
				SupplierProductID: "mock-" + text,
				Article:           "ART-" + text,
				Brand:             "MockBrand",
				Title:             "Mock Product for " + text,
				Price:             2000.0,
				Count:             200,
			},
			{
				SupplierID:        1,
				SupplierProductID: "mock-" + text,
				Article:           "ART-" + text,
				Brand:             "MockBrand",
				Title:             "Mock Product for " + text,
				Price:             50.0,
				Count:             2000,
			},
			{
				SupplierID:        1,
				SupplierProductID: "mock-" + text,
				Article:           "ART-" + text,
				Brand:             "MockBrand",
				Title:             "Mock Product for " + text,
				Price:             150.0,
				Count:             532,
			},
			{
				SupplierID:        1,
				SupplierProductID: "mock-" + text,
				Article:           "ART-" + text,
				Brand:             "MockBrand",
				Title:             "Mock Product for " + text,
				Price:             150.0,
				Count:             5222,
			},
		},
	}, nil
}

func (m *MockAPIAdapter) ListBasket() (*models.BasketList, error) {
	if m.ListBasketFn != nil {
		return m.ListBasketFn()
	}
	// Стандартный моковый ответ
	return &models.BasketList{
		TotalSum:   350.0,
		TotalCount: 3,
		Items: []models.BasketItem{
			{
				ID:                  1,
				SupplierID:          1,
				SupplierProductID:   "mock-001",
				Article:             "ART-001",
				Brand:               "MockBrand",
				Title:               "Mock Product 1",
				Price:               100.0,
				Count:               2,
				Comment:             "Test comment",
				DeliveryAddressGUID: "address-1",
			},
			{
				ID:                  2,
				SupplierID:          1,
				SupplierProductID:   "mock-002",
				Article:             "ART-002",
				Brand:               "MockBrand",
				Title:               "Mock Product 2",
				Price:               150.0,
				Count:               1,
				Comment:             "",
				DeliveryAddressGUID: "address-1",
			},
		},
	}, nil
}
