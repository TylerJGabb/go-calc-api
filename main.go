package main

import (
	"flag"
	"fmt"
	"go-calc-api/api"
	"go-calc-api/grpc"
	"go-calc-api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	grpcServerHost := flag.String("grpc-server-host", "localhost", "The gRPC server host")
	grpcServerPort := flag.String("grpc-server-port", "8080", "The gRPC server port")
	grpcConnectionTimeout := flag.Int("grpc-connection-timeout", 5, "The timeout for the gRPC client in seconds")
	httpPort := flag.String("http-port", "8080", "The port for the HTTP server")

	flag.Parse()
	fmt.Println("[STARTUP] grpcServerHost:", *grpcServerHost)
	fmt.Println("[STARTUP] grpcServerPort:", *grpcServerPort)
	fmt.Println("[STARTUP] grpcConnectionTimeout:", *grpcConnectionTimeout)
	fmt.Println("[STARTUP] httpPort:", *httpPort)

	serverAddr := fmt.Sprintf("%s:%s", *grpcServerHost, *grpcServerPort)
	client, err := grpc.NewClient(serverAddr, *grpcConnectionTimeout)
	if err != nil {
		panic(err)
	}
	app := gin.Default()
	app.Use(middleware.GrpcClientMiddleware(client))
	api.ApplyRoutes(app)
	panic(app.Run(":" + *httpPort))
}
