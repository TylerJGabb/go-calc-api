go build -o bin/exe
./bin/exe \
    --grpc-server-host=localhost \
    --grpc-server-port=9091 \
    --grpc-connection-timeout=5 \
    --http-port=8080


