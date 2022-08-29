package service

import (
	"context"
	"istio-started/gsrc/pbfiles"
)

type ProdService struct {
}

func (p *ProdService) GetProd(ctx context.Context, request *pbfiles.ProdRequest) (*pbfiles.ProdResponse, error) {
	return &pbfiles.ProdResponse{
		Result: &pbfiles.ProdModel{
			Id:   11,
			Name: "zx",
		},
	}, nil
}

func NewProdService() pbfiles.ProdServiceServer {
	return &ProdService{}
}
