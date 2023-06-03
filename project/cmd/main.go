package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"project/api"
	"project/internal/configs"
	"project/internal/database"
	"project/internal/orders"
	"project/internal/products"
	"project/internal/users"
)

func main() {
	err := configs.LoadEnvVariables()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	db, err := database.ConnectPostgres()
	if err != nil {
		panic(err)
	}

	orderRepo := orders.NewOrderRepository(db)
	orderService := orders.NewOrderService(orderRepo)
	userRepo := users.NewUserRepository(db)
	userService := users.NewUserService(userRepo)
	productRepo := products.NewProductRepository(db)
	productService := products.NewProductService(productRepo)

	apiHandler := api.NewAPIHandler(router, orderService, userService, productService)
	apiHandler.RegisterRoutes()

	router.Run("0.0.0.0:8000")
}
