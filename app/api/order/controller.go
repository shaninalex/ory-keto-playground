package order

import (
	"net/http"
	"testketo/app/pkg/mock_db"

	"github.com/gin-gonic/gin"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type Controller struct {
	db   mock_db.Database
	keto rts.CheckServiceClient
}

func NewController(db mock_db.Database, keto rts.CheckServiceClient) *Controller {
	return &Controller{
		db:   db,
		keto: keto,
	}
}

func (s *Controller) Register(router *gin.RouterGroup) {
	group := router.Group("orders")
	group.GET("", s.requirePermission("viewer"), s.handleOrders)
	group.GET("/:order", s.requirePermission("viewer"), s.handleOrder)
	group.POST("", s.requirePermission("creator"), s.handleOrderCreate)
}

func (s *Controller) requirePermission(relation string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetHeader("X-USER")
		if userId == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		res, err := s.keto.Check(c.Request.Context(), &rts.CheckRequest{
			Tuple: &rts.RelationTuple{
				Namespace: "PurchaseOrder",
				Object:    "global",
				Relation:  relation,
				Subject:   rts.NewSubjectID(userId),
			},
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if !res.Allowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		c.Next()
	}
}
