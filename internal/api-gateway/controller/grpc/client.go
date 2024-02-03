package grpc

import (
	"Clean-sweep-solution/internal/genproto/product/pb"
	product_pb "Clean-sweep-solution/internal/genproto/product/pb"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn    *grpc.ClientConn
	productClient product_pb.ProductServiceClient
}

func NewClient(port string) (*Client, error) {
	var conn *grpc.ClientConn
	addr := fmt.Sprintf("Clean-sweep-solution_product-service_app_1%v", port)

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:          conn,
		productClient: pb.NewProductServiceClient(conn),
		
	}, nil
}


