package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"

	"project/internal/auth"
	"project/internal/orders"
)

type OrderItemRequest struct {
	ProductCode string          `json:"product_code"`
	Quantity    int             `json:"quantity"`
	Price       decimal.Decimal `json:"price"`
}

type OrderCreateRequest struct {
	OrderNumber     string             `json:"order_number" binding:"required"`
	Status          string             `json:"status" binding:"required"`
	ShippingAddress string             `json:"shipping_address"`
	Total           decimal.Decimal    `json:"total"`
	TotalDiscount   decimal.Decimal    `json:"total_discount"`
	OrderItems      []OrderItemRequest `json:"order_items"`
}

type CancelCreateRequest struct {
	OrderNumber string `json:"order_number" binding:"required"`
}

type OrderDetailRequest struct {
	OrderNumber string `json:"order_number" binding:"required"`
}

func (ah *APIHandler) CreateOrderHandler(c *gin.Context) {
	c.Request.Header.Set("Authorization", c.GetHeader("Authorization"))
	auth.JWTMiddleware()(c)

	var req OrderCreateRequest
	var order orders.Order
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	username, _ := c.Get("username")
	user_id, _ := ah.userService.FindID(username.(string))

	order.OrderNumber = req.OrderNumber
	order.Status = req.Status
	order.ShippingAddress = req.ShippingAddress
	order.UserID = *user_id
	order.Total = req.Total
	order.TotalDiscount = req.TotalDiscount
	order.IsCancelled = false
	order.CreatedAt = time.Now().UTC()
	order.UpdatedAt = time.Now().UTC()

	for _, item := range req.OrderItems {
		var order_item orders.OrderItem
		product_id, err := ah.productService.FindProductID(item.ProductCode)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product code"})
			return
		}
		order_item = orders.OrderItem{
			ProductID: *product_id,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
		*order.OrderItems = append(*order.OrderItems, order_item)
	}

	_, err := ah.orderService.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order " + order.OrderNumber + " created successfully"})
}

func (ah *APIHandler) CancelOrderHandler(c *gin.Context) {
	c.Request.Header.Set("Authorization", c.GetHeader("Authorization"))
	auth.JWTMiddleware()(c)

	var req CancelCreateRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	username, _ := c.Get("username")
	user_id, _ := ah.userService.FindID(username.(string))

	err := ah.orderService.CancelOrder(req.OrderNumber, *user_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order " + req.OrderNumber + " cancelled successfully"})
}

func (ah *APIHandler) OrderDetailHandler(c *gin.Context) {
	c.Request.Header.Set("Authorization", c.GetHeader("Authorization"))
	auth.JWTMiddleware()(c)

	var req OrderDetailRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	username, _ := c.Get("username")
	user_id, _ := ah.userService.FindID(username.(string))

	order, err := ah.orderService.OrderDetail(req.OrderNumber, *user_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}
