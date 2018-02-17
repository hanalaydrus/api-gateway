build:
	protoc -I semanticContract/ semanticContract/semanticContract.proto --go_out=plugins=grpc:semanticContract
	protoc -I volumeContract/ volumeContract/volumeContract.proto --go_out=plugins=grpc:volumeContract

run:
	go run ./client_server.go