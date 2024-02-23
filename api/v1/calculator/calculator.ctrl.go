package calculator

import (
	"github.com/TylerJGabb/go-calc-grpc-contract/contract"
	"github.com/gin-gonic/gin"
)

type sumRequest struct {
	Numbers []int64 `json:"numbers" binding:"required"`
}

type divideRequest struct {
	A int64 `json:"a"`
	B int64 `json:"b"`
}

type addRequest struct {
	A int64 `json:"a"`
	B int64 `json:"b"`
}

func sum(ctx *gin.Context) {
	grpcClient := ctx.MustGet("client").(contract.CalculatorClient)
	var req sumRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err := grpcClient.Sum(ctx, &contract.NumbersRequest{Numbers: req.Numbers})
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"result": resp.Result})
}

func divide(ctx *gin.Context) {
	grpcClient := ctx.MustGet("client").(contract.CalculatorClient)
	var req divideRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err := grpcClient.Divide(ctx, &contract.CalculationRequest{A: req.A, B: req.B})
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"result": resp.Result})
}

func add(ctx *gin.Context) {
	grpcClient := ctx.MustGet("client").(contract.CalculatorClient)
	var req addRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err := grpcClient.Add(ctx, &contract.CalculationRequest{A: req.A, B: req.B})
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"result": resp.Result})
}
