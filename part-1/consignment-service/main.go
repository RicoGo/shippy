package main

import (
	"log"
	"net"

	"context"
	// 导入生成的consignment.pb.go
	pb "github.com/RicoGo/try-go-micro/part-1/consignment-service/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// 仓库接口
type IRepository interface {
	// 存放新货物
	Create(*pb.Consignment) (*pb.Consignment, error)
	// 查询所有货物
	GetAll() []*pb.Consignment
}

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
// 存放多批货物的仓库，实现了 IRepository 接口
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
// 定义微服务
type service struct {
	repo IRepository
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
// service 实现了consignment.pb.go 的 ShippingServiceServer 接口
// 使 service 作为 grpc 的服务端
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {

	// Save our consignment
	// 托运新货物
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	// 返回与我们在protobuf定义中创建的' Response '消息匹配的结果。
	return &pb.Response{Created: true, Consignment: consignment}, nil
}

// 获取目前所有托运的货物
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.GetAll()
	return &pb.Response{Consignments: consignments}, nil
}

func main() {

	repo := &Repository{}

	// 设置gRPC服务器
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("net.Listen", port)
	s := grpc.NewServer()

	/*
	 * 向gRPC服务器注册我们的服务，这将把我们的实现 service 绑定到我们的protobuf定义的
	 * 自动生成的 ShippingServiceServer 接口代码中
	 */
	pb.RegisterShippingServiceServer(s, &service{repo})

	// 在gRPC服务器上注册反射服务
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
