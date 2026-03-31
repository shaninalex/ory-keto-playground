package mock_db

import (
	"testketo/app/models"
	"time"
)

func ProvideDatabase() Database {
	return &mockDatabase{
		orders: map[uint64]*models.PurchaseOrder{
			123: {
				Id:          123,
				Code:        "ORD123",
				Name:        "test123",
				Count:       1,
				TotalAmount: 1,
				Items: []*models.PurchaseOrderItem{
					{
						Id:   123,
						Code: "ORD123",
						Name: "test123",
					},
				},
				ProcessedAt: time.Now(),
				CreatedAt:   time.Now(),
			},
			45345: {
				Id:          45345,
				Code:        "ORD45345",
				Name:        "test45345",
				Count:       1,
				TotalAmount: 1,
				Items: []*models.PurchaseOrderItem{
					{
						Id:   123,
						Code: "ORD123",
						Name: "test123",
					},
				},
				ProcessedAt: time.Now(),
				CreatedAt:   time.Now(),
			},
		},
	}
}
