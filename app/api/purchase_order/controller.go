package purchase_order

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (s *Controller) Register(router *gin.RouterGroup) {
	group := router.Group("purchase-orders")

	// TODO: apply configurable permission middleware
	//group.Use(permission.PermissionMiddleware(permission.PermissionMiddlewareConfig{
	//	Role: "manager",
	//	Access: []string{"viewer"}
	//})
	// something like that...

	group.GET("", s.handlePurchaseOrders)
	group.GET("/:order", s.handlePurchaseOrder)
}
