build:
	protoc --proto_path=.:$GOPATH/src/github.com/RicoGo/try-go-micro/part-2/consignment-cli --go_out=. --micro_out=. proto/consignment/consignment.proto
	docker build -t consignment-cli .

run:
	docker run -e MICRO_REGISTRY=mdns consignment-cli