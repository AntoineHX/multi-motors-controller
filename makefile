BINARY_NAME=motorsim
SRC_PATH=./src

all: generate-protos build

build:
	go build -o ${BINARY_NAME} ${SRC_PATH}/main.go
 
generate-protos:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ${SRC_PATH}/motorsim.proto

clean:
	go clean
	rm ${BINARY_NAME} ./src/*.pb.go