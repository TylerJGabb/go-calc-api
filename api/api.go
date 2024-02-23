package api

import (
	apiV1 "go-calc-api/api/v1"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		apiV1.ApplyRoutes(api)
	}

}
