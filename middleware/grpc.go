package middleware

import (
	"github.com/TylerJGabb/go-calc-grpc-contract/contract"
	"github.com/gin-gonic/gin"
)

func GrpcClientMiddleware(client contract.CalculatorClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("client", client)
		ctx.Next()
	}
}
