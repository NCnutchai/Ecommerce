package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"

	"project/internal/auth"
	"project/internal/products"
)

type ProductCreateRequest struct {
	Code  string          `json:"code" binding:"required"`
	Name  string          `json:"name" binding:"required"`
	Price decimal.Decimal `json:"price" binding:"required"`
}

type ProductDetailRequest struct {
	Code string `json:"code" binding:"required"`
}

func (ah *APIHandler) CreateProductHandler(c *gin.Context) {
	c.Request.Header.Set("Authorization", c.GetHeader("Authorization"))
	auth.JWTMiddleware()(c)

	var req ProductCreateRequest
	var product products.Product
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.Code = req.Code
	product.ProductName = req.Name
	product.Price = req.Price
	err := ah.productService.ProductCreate(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product " + product.ProductName + " created successfully"})
}

func (ah *APIHandler) ProductListHandler(c *gin.Context) {
	c.Request.Header.Set("Authorization", c.GetHeader("Authorization"))
	auth.JWTMiddleware()(c)

	products, err := ah.productService.ProductList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (ah *APIHandler) ProductDetailHandler(c *gin.Context) {
	c.Request.Header.Set("Authorization", c.GetHeader("Authorization"))
	auth.JWTMiddleware()(c)

	var req ProductDetailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := ah.productService.ProductDetail(req.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}
