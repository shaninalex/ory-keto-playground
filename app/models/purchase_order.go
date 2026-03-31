package models

import "time"

type PurchaseOrder struct {
	Id          uint64               `json:"id"`
	Code        string               `json:"code"`
	Name        string               `json:"name"`
	Count       float64              `json:"count"`
	TotalAmount float64              `json:"total_amount"`
	Items       []*PurchaseOrderItem `json:"items"`
	ProcessedAt time.Time            `json:"processed_at"`
	CreatedAt   time.Time            `json:"created_at"`
}

type PurchaseOrderItem struct {
	Id   uint64 `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
