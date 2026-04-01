package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Controller) handleOrders(c *gin.Context) {
	orders, err := s.db.GetOrders(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"orders": orders})
}
