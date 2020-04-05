build:
	protoc --proto_path=.:$GOPATH/src/github.com/RicoGo/try-go-micro/part-2/consignment-service --go_out=. --micro_out=. proto/consignment/consignment.proto

	docker build -t consignment-service .

run:
	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns consignment-service