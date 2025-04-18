package client

import (
	"api_gateway/config"
	"api_gateway/proto/barang"
	"api_gateway/proto/category"
	"api_gateway/proto/role"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	barangConn *grpc.ClientConn
	userConn   *grpc.ClientConn
}

func NewGrpcClient(cfg config.Config) (*GrpcClient, error) {
	barangConn, err := grpc.NewClient(cfg.GRPCAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	userConn, err := grpc.NewClient(cfg.GRPCAddressUser, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &GrpcClient{
		barangConn: barangConn,
		userConn:   userConn,
	}, nil
}

func (g *GrpcClient) BarangClient() barang.BarangServiceClient {
	return barang.NewBarangServiceClient(g.barangConn)
}

func (g *GrpcClient) CategoryClient() category.CategoryServiceClient {
	return category.NewCategoryServiceClient(g.barangConn)
}

func (g *GrpcClient) RoleClient() role.RoleServiceClient {
	return role.NewRoleServiceClient(g.userConn)
}

func (g *GrpcClient) Close() {
	if g.barangConn != nil {
		g.barangConn.Close()
	}
}
