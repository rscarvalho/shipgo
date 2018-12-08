package main

import (
  "github.com/micro/go-micro"
  pb "github.com/rscarvalho/shipgo/consignment-service/proto/consignment"
  "golang.org/x/net/context"
  "log"
)

const (
	port = ":50051"
)

type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

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

type service struct {
	repo IRepository
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, out *pb.Response) error {
  out.Consignments = s.repo.GetAll()
	return nil
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, out *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	out.Created = true
  out.Consignment = consignment
	return nil
}

func main() {
  repo := &Repository{}

	srv := micro.NewService(
	  micro.Name("go.micro.srv.consignment"),
	  micro.Version("latest"))

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
	  log.Fatal(err)
  }
}
