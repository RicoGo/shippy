package main

import (
	"context"
	"encoding/json"
	pb "github.com/RicoGo/try-go-micro/part-2/consignment-cli/proto/consignment"
	//microClient "github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config/cmd"
	"io/ioutil"
	"log"
	"os"
)

const (
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
	cmd.Init()
	server := micro.NewService()
	server.Init()

	//client := pb.NewShippingService("go_micro_srv_consignment", microClient.DefaultClient)
	client := pb.NewShippingService("go_micro_srv_consignment", server.Client())

	file := defaultFileName
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v\n", err)
	}

	// context.Background() -> new(emptyCtx)
	//resp, err := client.CreateConsignment(context.Background(), consignment)
	resp, err := client.CreateConsignment(context.TODO(), consignment)
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
