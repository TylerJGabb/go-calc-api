package apiV1

import (
	"go-calc-api/api/v1/calculator"
	"go-calc-api/api/v1/health"

	"github.com/gin-gonic/gin"
)

type Translator struct {
}

func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	calculator.ApplyRoutes(v1)
	health.ApplyRoutes(v1)
}
