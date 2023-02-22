package helper

import (
	"mfaisalh12/product-restful-api/model/domain"
	"mfaisalh12/product-restful-api/model/web"
)

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id: 					product.Id,
		Name: 				product.Name,
		Description:	product.Description,
		Category:			product.Category,
		Price:				product.Price,
	}
}

func ToAllProductResponse(products []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}