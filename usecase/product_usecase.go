package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {

	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUsecase) GetProductById(idProduct string) (*model.Product, error) {
	product, err := pu.repository.GetProductById(idProduct)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductUsecase) UpdateProduct(idProduct string, product model.Product) (model.Product, error) {
	productId, err := pu.repository.UpdateProduct(idProduct, product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}
func (pu *ProductUsecase) DeleteProduct(idProduct string) error {
	err := pu.repository.DeleteProduct(idProduct)
	if err != nil {
		return err
	}

	return nil
}
