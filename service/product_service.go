package service

import (
	"context"
	"mfaisalh12/product-restful-api/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse
	Delete(ctx context.Context, productId int)
	FindById(ctx context.Context, productId int) web.ProductResponse
	FindAll(ctx context.Context, query map[string]string) []web.ProductResponse
}