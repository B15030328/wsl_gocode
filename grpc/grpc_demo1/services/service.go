package services

import (
	context "context"
	"fmt"
)

type ProducService struct {
}

func (service1 *ProducService) GetProductStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	fmt.Println("请求来了")
	resp := ProdResponse{
		ProdStock: 21,
	}
	return &resp, nil
}
