package api

import (
	"net/http"
	"testketo/app/api/order"
	"testketo/app/pkg/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func ProvideRouter() *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.DebugMode)

	router.RedirectTrailingSlash = false
	router.RedirectFixedPath = false

	router.GET("/_health", HealthRoute)

	return router
}

func HealthRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"app":     "ketotestapp",
		"version": "0.0.1",
	})
}

type ApiDeps struct {
	dig.In

	Config                  *config.Config
	PurchaseOrderController *order.Controller
}

func ProvideAPI(deps ApiDeps) *gin.Engine {
	router := ProvideRouter()
	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")
	deps.PurchaseOrderController.Register(v1)

	return router
}
