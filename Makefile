build:
	protoc -I semanticContract/ semanticContract/semanticContract.proto --go_out=plugins=grpc:semanticContract
	protoc -I volumeContract/ volumeContract/volumeContract.proto --go_out=plugins=grpc:volumeContract
	protoc -I densityContract/ densityContract/densityContract.proto --go_out=plugins=grpc:densityContract
	protoc -I gatewayContract/ gatewayContract/gatewayContract.proto --go_out=plugins=grpc:gatewayContract

run:
	go run ./client_server.go