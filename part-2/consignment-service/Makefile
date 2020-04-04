build:
	protoc -I . --go_out=$(GOPATH)/src/github.com/RicoGo/try-go-micro/part-2/consignment-service/proto/consignment --micro_out=$(GOPATH)/src/github.com/RicoGo/try-go-micro/part-2/consignment-service/proto/consignment $(GOPATH)/src/github.com/RicoGo/try-go-micro/part-2/consignment-service/proto/consignment/consignment.proto

	docker build -t consignment-service .

run:
	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns consignment-service