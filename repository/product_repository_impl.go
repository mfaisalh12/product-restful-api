package repository

import (
	"context"
	"database/sql"
	"errors"
	"mfaisalh12/product-restful-api/helper"
	"mfaisalh12/product-restful-api/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}


func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "INSERT INTO products(name, description, category, price) VALUES(?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Description, product.Category, product.Price)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)
	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "UPDATE products SET name = ?, description = ?, category = ?, price = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Description, product.Category, product.Price, product.Id)
	helper.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	SQL := "DELETE FROM products WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	SQL := "SELECT * FROM products WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	product := domain.Product{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Category, &product.Price)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("product is not found")
	}
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, filter map[string]string) []domain.Product {
	SQL := "SELECT * FROM products"
	if len(filter["category"]) != 0 {
		SQL += " WHERE category LIKE '%" + filter["category"] +"%'"
	}
	if len(filter["price"]) != 0 {
		SQL += " ORDER BY price " + filter["price"]
	}

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Category, &product.Price)
		helper.PanicIfError(err)
		products = append(products, product)
	}
	return products
}

func (repository *ProductRepositoryImpl) FindQueryByCategory(ctx context.Context, tx *sql.Tx, category string) []domain.Product {
	SQL := "SELECT * FROM products WHERE category = ?"
	rows, err := tx.QueryContext(ctx, SQL, category)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Category, &product.Price)
		helper.PanicIfError(err)
		products = append(products, product)
	}
	return products

}