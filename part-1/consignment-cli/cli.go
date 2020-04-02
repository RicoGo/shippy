package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"context"
	pb "github.com/RicoGo/try-go-micro/part-1/consignment-service/proto/consignment"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultFileName = "consignment.json"
)

// 读取 consignment.json 中记录的货物信息
func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &consignment); err != nil {
		return nil, err
	}
	return consignment, nil
}

func main() {
	// 连接到gRPC服务器
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	// 初始化 gRPC 客户端
	client := pb.NewShippingServiceClient(conn)

	file := defaultFileName
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	// 解析货物信息
	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	// context.Background() -> new(emptyCtx)
	// 调用 RPC,将货物存储到我们自己的仓库里
	resp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("cli.main.CreateConsignment: %v", err)
	}
	// 新货物是否托运成功
	log.Printf("Created: %t", resp.Created)

	// 列出目前所有托运的货物
	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("cli.main.GetConsignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
