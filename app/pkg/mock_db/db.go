package mock_db

import (
	"context"
	"encoding/json"
	"errors"
	"maps"
	"math/rand"
	"os"
	"slices"
	"sync"
	"testketo/app/models"
	"time"
)

type Database interface {
	GetOrders(ctx context.Context) ([]*models.PurchaseOrder, error)
	GetOrder(ctx context.Context, id uint64) (*models.PurchaseOrder, error)
	AddOrder(ctx context.Context, order *models.PurchaseOrder) (*models.PurchaseOrder, error)
}

var _ Database = (*mockDatabase)(nil)

type mockDatabase struct {
	orders map[uint64]*models.PurchaseOrder

	mu sync.RWMutex
}

func (s *mockDatabase) GetOrders(ctx context.Context) ([]*models.PurchaseOrder, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return slices.Collect(maps.Values(s.orders)), nil
}

func (s *mockDatabase) GetOrder(ctx context.Context, id uint64) (*models.PurchaseOrder, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if order, ok := s.orders[id]; ok {
		return order, nil
	}
	return nil, errors.New("order not found")
}

func (s *mockDatabase) AddOrder(ctx context.Context, order *models.PurchaseOrder) (*models.PurchaseOrder, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := rand.Uint64()
	order.Id = id
	order.CreatedAt = time.Now()

	s.orders[id] = order

	return order, nil
}

func (s *mockDatabase) importData() error {
	b, err := os.ReadFile("resources/orders.json")
	if err != nil {
		return err
	}
	var orders []*models.PurchaseOrder
	if err = json.Unmarshal(b, &orders); err != nil {
		return err
	}
	s.mu.Lock()
	for _, order := range orders {
		s.orders[order.Id] = order
	}
	s.mu.Unlock()

	return nil
}
