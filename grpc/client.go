package grpc

import (
	"context"
	"time"

	"github.com/TylerJGabb/go-calc-grpc-contract/contract"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(serverAddr string, timeout int) (contract.CalculatorClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.DialContext(ctx, serverAddr, opts...)
	if err != nil {
		return nil, err
	}
	client := contract.NewCalculatorClient(conn)
	return client, nil
}
