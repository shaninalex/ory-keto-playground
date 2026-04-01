package api

import (
	"testketo/app/api/order"

	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = order.Module(c)

	_ = c.Provide(ProvideAPI)

	return nil
}
