package api

import (
	"github.com/gin-gonic/gin"

	"project/internal/auth"
	"project/internal/orders"
	"project/internal/products"
	"project/internal/users"
)

type APIHandler struct {
	router         *gin.Engine
	authService    *auth.AuthService
	orderService   orders.Service
	productService products.Service
	userService    users.Service
}

func NewAPIHandler(router *gin.Engine, orderService orders.Service, userService users.Service, productService products.Service) *APIHandler {
	return &APIHandler{
		router:         router,
		orderService:   orderService,
		userService:    userService,
		productService: productService,
	}
}

func (ah *APIHandler) RegisterRoutes() {
	ah.router.POST("api/login", ah.LoginHandler)
	ah.router.POST("/api/register", ah.CreateUserHandler)

	ah.router.POST("/api/user/profile", ah.GetProfileHandler)
	ah.router.POST("/api/user/order-history", ah.GetOrderHistoryHandler)

	ah.router.POST("/api/order/create", ah.CreateOrderHandler)
	ah.router.POST("/api/order/cancel", ah.CancelOrderHandler)
	ah.router.POST("/api/order/detail", ah.OrderDetailHandler)

	ah.router.POST("/api/product/create", ah.CreateProductHandler)
	ah.router.POST("/api/product/list", ah.ProductListHandler)
	ah.router.POST("/api/product/detail", ah.ProductDetailHandler)

	ah.router.Use(auth.JWTMiddleware())
}
