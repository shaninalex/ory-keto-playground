package api

import (
	"testketo/app/api/purchase_order"

	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = purchase_order.Module(c)

	_ = c.Provide(ProvideAPI)

	return nil
}
