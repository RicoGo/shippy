package main

import (
	"context"
	"log"
	// 导入生成的consignment.pb.go
	pb "github.com/RicoGo/try-go-micro/part-2/consignment-service/proto/consignment"
	"github.com/micro/go-micro/v2"
)

// 仓库接口
type IRepository interface {
	// 存放新货物
	Create(*pb.Consignment) (*pb.Consignment, error)
	// 查询所有货物
	GetAll() []*pb.Consignment
}

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

// 定义微服务
// 实现 consignment.pb.micro.go ShippingServiceHandler 接口
type service struct {
	repo Repository
}

// 实现 consignment.pb.micro.go ShippingServiceHandler 接口
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	resp = &pb.Response{
		Created:     true,
		Consignment: consignment,
	}

	return nil
}

func (s *service) GetConsignments(ctx context.Context, rsq *pb.GetRequest, resp *pb.Response) error {
	allConsignments := s.repo.GetAll()
	resp = &pb.Response{Consignments: allConsignments}
	return nil
}

func main() {
	server := micro.NewService(
		micro.Name("go_micro_srv_consignment"),
		micro.Version("latest"),
	)

	server.Init()
	repo := Repository{}
	pb.RegisterShippingServiceHandler(server.Server(), &service{repo})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve : %v\n", err)
	}

}
