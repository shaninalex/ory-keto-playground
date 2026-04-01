package order

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Controller) handleOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("order"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := s.db.GetOrder(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"order": order})
}
