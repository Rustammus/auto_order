package service

import (
	"auto_order/internal/models"
	"auto_order/internal/repo"
	"auto_order/internal/sups"
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"
)

type ToOrder struct {
	toOrderRepo *repo.ToOrderRepo
	productRepo *repo.ProductRepo
	search      []models.ToOrder
}

func NewToOrderService(repo *repo.ToOrderRepo) *ToOrder {
	return &ToOrder{toOrderRepo: repo}
}

func (s *ToOrder) ColumnName(col int) (name string) {
	switch col {
	case 0:
		name = "Поставщик"
	case 1:
		name = "Код"
	case 2:
		name = "Артикул"
	case 3:
		name = "Название"
	case 4:
		name = "Кол-во"
	default:
		name = "Undefined"
	}
	return
}

func (s *ToOrder) Cell(row, col int) (cell string) {
	if row >= len(s.search) {
		return "RANGE"
	}

	o := s.search[row]

	switch col {
	case 0:
		cell = sups.SupNames[o.SupID]
	case 1:
		cell = o.SupCode.String
	case 2:
		if o.Product != nil {
			cell = o.Product.Article.String
		}
	case 3:
		if o.Product != nil {
			cell = o.Product.Name
		}
	case 4:
		cell = strconv.Itoa(o.Count)
	default:
		cell = "Undefined"
	}
	return
}

func (s *ToOrder) Size() (rows int, cols int) {
	rows = len(s.search)
	if rows != 0 {
		cols = 5
	}

	return
}

func (s *ToOrder) Search() {
	orders, err := s.toOrderRepo.ListOrdersWithProducts(context.Background())
	if err != nil {
		log.Printf("service.ToOrder.Search err %v", err)
		return
	}

	s.search = orders
}

func (s *ToOrder) AddToOrder(p *models.Product) {

	// supID = 1
	var supID int64 = sups.PlanetID

	// Scan code by supID
	code := ""
	switch supID {
	case 0:
		code = p.KngIDVal()
	case 1:
		code = p.PlanetIDVal()
	case 2:
		code = p.ProgressIDVal()
	case 3:
		code = p.EsscoIDVal()
	}
	if code == "" {
		log.Printf("service.ToOrder.AddToOrder product Sup:%d code is empty or null", supID)
		return
	}

	// Search existing product
	rOrder, err := s.toOrderRepo.FindByProductAndSup(context.Background(), p.ID, supID)
	if err != nil {
		log.Printf("service.ToOrder.FindByProductAndSup read err %v", err)
	}

	// Adding/creating toOrder
	if rOrder != nil {
		rOrder.Count++

		err = s.toOrderRepo.Update(context.Background(), rOrder)
		if err != nil {
			log.Printf("service.ToOrder.FindByProductAndSup update err %v", err)
		}

	} else {
		o := &models.ToOrder{
			ProductID: p.ID,
			SupID:     supID,
			SupCode:   sql.NullString{code, true},
			Count:     1,
		}

		_, err = s.toOrderRepo.Create(context.Background(), o)
		if err != nil {
			log.Printf("service.ToOrder.FindByProductAndSup create err %v", err)
		}
	}
}

// ========= Private Methods ==============

// getByID возвращает заказ по его ID
func (s *ToOrder) getByID(ctx context.Context, id int64) (*models.ToOrder, error) {
	if id <= 0 {
		return nil, errors.New("invalid ID")
	}

	order, err := s.toOrderRepo.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("order not found")
	}

	return order, nil
}

// getBySupplierID возвращает все заказы для указанного поставщика
func (s *ToOrder) getBySupplierID(ctx context.Context, supID int64) ([]models.ToOrder, error) {
	if supID <= 0 {
		return nil, errors.New("invalid supplier ID")
	}

	orders, err := s.toOrderRepo.FindBySupplierID(ctx, supID)
	if err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return nil, errors.New("no orders found for this supplier")
	}

	return orders, nil
}

// createOrder создает новый заказ
func (s *ToOrder) createOrder(ctx context.Context, order *models.ToOrder) (*models.ToOrder, error) {
	// Валидация
	if order.ProductID <= 0 {
		return nil, errors.New("product ID is required")
	}
	if order.SupID <= 0 {
		return nil, errors.New("supplier ID is required")
	}
	if order.Count <= 0 {
		return nil, errors.New("count must be positive")
	}

	// Создаем заказ
	id, err := s.toOrderRepo.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	// Получаем созданный заказ для возврата
	createdOrder, err := s.toOrderRepo.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	return createdOrder, nil
}

// updateCount обновляет количество товара в заказе
func (s *ToOrder) updateCount(ctx context.Context, orderID int64, newCount int) (*models.ToOrder, error) {
	if orderID <= 0 {
		return nil, errors.New("invalid order ID")
	}
	if newCount <= 0 {
		return nil, errors.New("count must be positive")
	}

	// Получаем текущий заказ
	order, err := s.toOrderRepo.Find(ctx, orderID)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("order not found")
	}

	// Обновляем количество
	order.Count = newCount
	err = s.toOrderRepo.Update(ctx, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// deleteOrder удаляет заказ
func (s *ToOrder) deleteOrder(ctx context.Context, orderID int64) error {
	if orderID <= 0 {
		return errors.New("invalid order ID")
	}

	// Проверяем существование заказа
	order, err := s.toOrderRepo.Find(ctx, orderID)
	if err != nil {
		return err
	}
	if order == nil {
		return errors.New("order not found")
	}

	return s.toOrderRepo.Delete(ctx, orderID)
}
