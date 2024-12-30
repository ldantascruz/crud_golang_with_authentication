package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-api/controller"
	"go-api/db"
	"go-api/middleware"
	"go-api/repository"
	"go-api/usecase"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatalln("SECRET_KEY is not set. Check your .env file.")
	}

	log.Println("SECRET_KEY loaded successfully")

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	ProductController := controller.NewProductController(ProductUsecase)

	AuthUsecase := usecase.NewAuthUsecase()
	AuthController := controller.NewAuthController(AuthUsecase)

	server.POST("/login", AuthController.Login)

	server.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	protected := server.Group("/products", middleware.AuthMiddleware())
	protected.GET("", ProductController.GetProducts)
	protected.POST("", ProductController.CreateProduct)
	protected.GET("/:id", ProductController.GetProductById)
	protected.PUT("/:id", ProductController.UpdateProduct)
	protected.DELETE("/:id", ProductController.DeleteProduct)

	server.Run(":8000")

}
