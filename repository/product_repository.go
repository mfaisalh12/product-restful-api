package repository

import (
	"context"
	"database/sql"
	"mfaisalh12/product-restful-api/model/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, product domain.Product)
	FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx, filter map[string]string) []domain.Product
	FindQueryByCategory(ctx context.Context, tx *sql.Tx,  category string) []domain.Product
}