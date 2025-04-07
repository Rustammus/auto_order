package kng

import (
	"auto_order/internal/config"
	"auto_order/internal/models"
	schema "auto_order/internal/schemas/planet"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const supplierID = 0

type APIAdapter struct {
	token   string
	baseUrl string
	client  *http.Client
}

func NewAPIAdapter() *APIAdapter {
	conf := config.GetConfig()
	return &APIAdapter{
		token:   conf.Planet.Token,
		baseUrl: conf.Planet.URL,
		client:  &http.Client{},
	}
}

// SearchByCode выполняет поиск товаров по коду
func (a *APIAdapter) searchByCode(code string) (*models.SearchResult, error) {
	// Формируем URL с query-параметрами
	reqUrl, err := url.Parse(a.baseUrl + "/v1/search/products")
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %v", err)
	}

	params := url.Values{}
	params.Add("product_code", code)
	params.Add("sort", "price")
	reqUrl.RawQuery = params.Encode()

	// Создаем запрос
	req, err := http.NewRequest("GET", reqUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Authorization", "Bearer "+a.token)
	req.Header.Set("Accept", "application/json")

	// Выполняем запрос
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус код
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Читаем тело ответа
	searchResult := &schema.SearchResponse{}

	err = json.NewDecoder(resp.Body).Decode(searchResult)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal search result: %v", err)
	}

	return searchResult.ToDTO(), nil
}

// SearchByText выполняет поиск товаров по тексту
func (a *APIAdapter) searchByText(text string) (*models.SearchResult, error) {
	// Формируем URL с query-параметрами
	reqUrl, err := url.Parse(a.baseUrl + "/v1/search/products")
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %v", err)
	}

	params := url.Values{}
	params.Add("query", text)
	params.Add("sort", "price")
	reqUrl.RawQuery = params.Encode()

	// Создаем запрос
	req, err := http.NewRequest("GET", reqUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Authorization", "Bearer "+a.token)
	req.Header.Set("Accept", "application/json")

	// Выполняем запрос
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус код
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Читаем тело ответа
	searchResult := &schema.SearchResponse{}

	err = json.NewDecoder(resp.Body).Decode(searchResult)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal search result: %v", err)
	}

	return searchResult.ToDTO(), nil
}

// ListBasket получает список товаров в корзине
func (a *APIAdapter) listBasket() (*models.BasketList, error) {
	// Формируем URL
	reqUrl := a.baseUrl + "/v1/basket/products"

	// Создаем запрос
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Authorization", "Bearer "+a.token)
	req.Header.Set("Accept", "application/json")

	// Выполняем запрос
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус код
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	basketListSchema := &schema.Basket{}

	err = json.NewDecoder(resp.Body).Decode(basketListSchema)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal search result: %v", err)
	}

	// Читаем тело ответа
	return basketListSchema.ToDTO(), nil
}

// addToBasket добавляет товары в корзину
func (a *APIAdapter) addToBasket(items *models.BasketAddProducts) (*models.BasketAddResult, error) {
	// Формируем URL
	reqUrl := a.baseUrl + "/v1/basket/products"

	// Преобразуем входные данные в тело запроса
	var requestBody schema.BasketAddRequest
	requestBody.FromDTO(items)

	// Подготавливаем тело запроса
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	// Создаем запрос
	req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Authorization", "Bearer "+a.token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Выполняем запрос
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус код
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Читаем тело ответа
	purchaseResult := &schema.BasketAddResponse{}
	err = json.NewDecoder(resp.Body).Decode(purchaseResult)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal purchase result: %v", err)
	}

	return purchaseResult.ToDTO(), nil
}

// removeFromBasket добавляет товары в корзину
func (a *APIAdapter) removeFromBasket(items *models.BasketRemove) (*models.BasketList, error) {
	// Формируем URL
	reqUrl := a.baseUrl + "/v1/basket/products"

	// Преобразуем входные данные в тело запроса
	var requestBody schema.BasketRemoveRequest
	requestBody.FromDTO(items)

	// Подготавливаем тело запроса
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	// Создаем запрос
	req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Authorization", "Bearer "+a.token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Выполняем запрос
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус код
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Читаем тело ответа
	purchaseResult := &schema.Basket{}
	err = json.NewDecoder(resp.Body).Decode(purchaseResult)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal purchase result: %v", err)
	}

	return purchaseResult.ToDTO(), nil
}
