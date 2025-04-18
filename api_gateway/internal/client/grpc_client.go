package client

import (
	"api_gateway/config"
	"api_gateway/proto/barang"
	"api_gateway/proto/category"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	barangConn *grpc.ClientConn
}

func NewGrpcClient(cfg config.Config) (*GrpcClient, error) {
	barangConn, err := grpc.NewClient(cfg.GRPCAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &GrpcClient{
		barangConn: barangConn,
	}, nil
}

func (g *GrpcClient) BarangClient() barang.BarangServiceClient {
	return barang.NewBarangServiceClient(g.barangConn)
}

func (g *GrpcClient) CategoryClient() category.CategoryServiceClient {
	return category.NewCategoryServiceClient(g.barangConn)
}

func (g *GrpcClient) Close() {
	if g.barangConn != nil {
		g.barangConn.Close()
	}
}
