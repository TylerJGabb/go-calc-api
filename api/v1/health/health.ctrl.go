package health

import (
	"github.com/TylerJGabb/go-calc-grpc-contract/contract"
	"github.com/gin-gonic/gin"
)

func getHealth(c *gin.Context) {
	_ = c.MustGet("client").(contract.CalculatorClient)
	c.JSON(200, gin.H{"status": ":thumbs-up:"})
}
