package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/TylerJGabb/go-calc-grpc-contract/contract"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	serverAddr := flag.String(
		"server", "localhost:8080", "The server address in the format of host:port",
	)

	flag.Parse()
	fmt.Printf("Connecting to server at %s\n", *serverAddr)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.DialContext(ctx, *serverAddr, opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := contract.NewCalculatorClient(conn)

	result, err := client.Sum(ctx, &contract.NumbersRequest{
		Numbers: []int64{1, 2, 3, 4, 5},
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Sum result: %d\n", result.Result)

	result, err = client.Divide(ctx, &contract.CalculationRequest{
		A: 10,
		B: 5,
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Divide result: %d\n", result.Result)

	result, err = client.Add(ctx, &contract.CalculationRequest{
		A: 10,
		B: 5,
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Add result: %d\n", result.Result)

	result, err = client.Divide(ctx, &contract.CalculationRequest{
		A: 10,
		B: 0,
	})

	if err != nil {
		fmt.Printf("Divide error: %s\n", err.Error())
	} else {
		panic(fmt.Errorf("expected error, got %d", result.Result))
	}

}
