package controller

import (
	"github.com/gin-gonic/gin"
	"go-api/model"
	"go-api/usecase"
	"net/http"
)

type ProductController struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) ProductController {
	return ProductController{
		ProductUsecase: usecase,
	}
}
func (p *ProductController) GetProducts(context *gin.Context) {
	products, err := p.ProductUsecase.GetProducts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(context *gin.Context) {

	var product model.Product
	err := context.BindJSON(&product)
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.ProductUsecase.CreateProduct(product)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusCreated, insertedProduct)
}

func (p *ProductController) GetProductById(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		response := model.Response{
			Message: "id is required",
		}
		context.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.ProductUsecase.GetProductById(id)
	if err != nil {
		if product == nil {
			response := model.Response{
				Message: "product not found",
			}
			context.JSON(http.StatusNotFound, response)
			return
		}
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, product)
}

func (p *ProductController) UpdateProduct(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		response := model.Response{
			Message: "id is required",
		}
		context.JSON(http.StatusBadRequest, response)
		return
	}

	var product model.Product
	err := context.BindJSON(&product)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	updatedProduct, err := p.ProductUsecase.UpdateProduct(id, product)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, updatedProduct)
}

func (p *ProductController) DeleteProduct(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		response := model.Response{
			Message: "id is required",
		}
		context.JSON(http.StatusBadRequest, response)
		return
	}

	err := p.ProductUsecase.DeleteProduct(id)
	if err != nil {

		context.JSON(http.StatusInternalServerError, err)
		return
	}

	response := model.Response{
		Message: "product deleted successfully",
	}
	context.JSON(http.StatusOK, response)
}
