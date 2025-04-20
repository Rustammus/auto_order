package kng

import (
	"auto_order/internal/config"
	"auto_order/internal/models"
	"net/http"
)

const supplierID = 2

type APIAdapter struct {
	token   string
	baseUrl string
	client  *http.Client
}

func NewAPIAdapter() *APIAdapter {
	conf := config.GetConfig()
	return &APIAdapter{
		token:   conf.Progress.Token,
		baseUrl: conf.Progress.URL,
		client:  &http.Client{},
	}
}

// SearchByCode выполняет поиск товаров по коду
func (a *APIAdapter) searchByCode(code string) (*models.SearchResult, error) {
	// Формируем URL с query-параметрами
	panic("implement me")
}

// SearchByText выполняет поиск товаров по тексту
func (a *APIAdapter) searchByText(text string) (*models.SearchResult, error) {
	// Формируем URL с query-параметрами
	panic("implement me")
}

// ListBasket получает список товаров в корзине
func (a *APIAdapter) listBasket() (*models.BasketList, error) {
	// Формируем URL
	panic("implement me")
}

// addToBasket добавляет товары в корзину
func (a *APIAdapter) addToBasket(items *models.BasketAddProducts) (*models.BasketAddResult, error) {
	// Формируем URL
	panic("implement me")
}

// removeFromBasket добавляет товары в корзину
func (a *APIAdapter) removeFromBasket(items *models.BasketRemove) (*models.BasketList, error) {
	// Формируем URL
	panic("implement me")
}
