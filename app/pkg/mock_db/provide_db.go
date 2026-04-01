package mock_db

import (
	"testketo/app/models"
)

func ProvideDatabase() Database {
	s := &mockDatabase{
		orders: make(map[uint64]*models.PurchaseOrder),
	}

	if err := s.importData(); err != nil {
		panic(err)
	}
	return s
}
