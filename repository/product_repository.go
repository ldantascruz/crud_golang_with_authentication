package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {

	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	err = rows.Close()
	if err != nil {
		return nil, err
	}

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (string, error) {

	var id string
	query, err := pr.connection.Prepare("INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	query.Close()
	return id, nil
}

func (pr *ProductRepository) GetProductById(idProduct string) (*model.Product, error) {

	query, err := pr.connection.Prepare("SELECT id, product_name, price FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var product model.Product

	err = query.QueryRow(idProduct).Scan(
		&product.ID,
		&product.Name,
		&product.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("product not found")
		}

		return nil, err
	}

	query.Close()

	return &product, nil
}

func (pr *ProductRepository) UpdateProduct(idProduct string, product model.Product) (string, error) {
	query, err := pr.connection.Prepare("UPDATE product SET product_name = $1, price = $2 WHERE id = $3 RETURNING id")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var id string
	err = query.QueryRow(product.Name, product.Price, idProduct).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	query.Close()

	return idProduct, nil
}

func (pr *ProductRepository) DeleteProduct(idProduct string) error {
	query, err := pr.connection.Prepare("DELETE FROM product WHERE id = $1 RETURNING id")
	if err != nil {
		fmt.Println(err)
		return err
	}

	var id string
	err = query.QueryRow(idProduct).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	query.Close()

	return nil
}
