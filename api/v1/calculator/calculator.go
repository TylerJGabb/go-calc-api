package calculator

import "github.com/gin-gonic/gin"

func ApplyRoutes(r *gin.RouterGroup) {
	calculator := r.Group("/calculator")
	{
		calculator.GET("/sum", sum)
		calculator.GET("/divide", divide)
		calculator.GET("/add", add)
	}
}
