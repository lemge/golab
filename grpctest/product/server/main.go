package main

import (
	"context"
	"grpctest/product"
	"log"
	"net"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":50051"
)

type productInfoServer struct {
	product.UnimplementedProductInfoServer
	productMap map[string]*product.Product
}

//添加商品
func (s *productInfoServer) AddProduct(ctx context.Context, req *product.Product) (resp *product.ProductId, err error) {
	resp = &product.ProductId{}
	out, err := uuid.NewV4()
	if err != nil {
		return resp, status.Errorf(codes.Internal, "err while generate the uuid ", err)
	}

	req.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*product.Product)
	}

	s.productMap[req.Id] = req
	resp.Value = req.Id
	return
}

//获取商品
func (s *productInfoServer) GetProduct(ctx context.Context, req *product.ProductId) (resp *product.Product, err error) {
	if s.productMap == nil {
		s.productMap = make(map[string]*product.Product)
	}
	resp = s.productMap[req.Value]
	return
}

func newServer() *productInfoServer {
	s := &productInfoServer{}
	return s
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("net listen err ", err)
		return
	}

	s := grpc.NewServer()
	product.RegisterProductInfoServer(s, &productInfoServer{})
	log.Println("start gRPC listen on port " + port)
	if err := s.Serve(listener); err != nil {
		log.Println("failed to serve...", err)
		return
	}
}
