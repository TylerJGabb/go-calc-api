package health

import "github.com/gin-gonic/gin"

func ApplyRoutes(r *gin.RouterGroup) {
	health := r.Group("/health")
	{
		health.GET("", getHealth)
	}
}
