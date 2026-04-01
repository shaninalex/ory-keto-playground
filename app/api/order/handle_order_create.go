package order

import (
	"net/http"
	"testketo/app/models"

	"github.com/gin-gonic/gin"
)

func (s *Controller) handleOrderCreate(c *gin.Context) {
	var payload models.PurchaseOrder
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := s.db.AddOrder(c.Request.Context(), &payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"order": order})
}
