Iniciando com gRPC/Go

Criado comunicação entre o cliente e servidor usando gRPC/Go.

comando para compilar o arquivo .proto:
  protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb

comando para executar o servidor:
  go run cmd/server/server.go

comando para executar o cliente:
  go run cmd/client/client.go