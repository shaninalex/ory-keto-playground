package mock_db

import (
	"context"
	"errors"
	"maps"
	"slices"
	"testketo/app/models"
)

type Database interface {
	GetOrders(ctx context.Context) ([]*models.PurchaseOrder, error)
	GetOrder(ctx context.Context, id uint64) (*models.PurchaseOrder, error)
}

var _ Database = (*mockDatabase)(nil)

type mockDatabase struct {
	orders map[uint64]*models.PurchaseOrder
}

func (m mockDatabase) GetOrders(ctx context.Context) ([]*models.PurchaseOrder, error) {
	return slices.Collect(maps.Values(m.orders)), nil
}

func (m mockDatabase) GetOrder(ctx context.Context, id uint64) (*models.PurchaseOrder, error) {
	if order, ok := m.orders[id]; ok {
		return order, nil
	}
	return nil, errors.New("order not found")
}
