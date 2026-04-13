package main

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jupitters/go-pizza-tracker/internal/models"
)

type OrderFormData struct {
	PizzaTypes []string
	PizzaSizes []string
}

type OrderRequest struct {
	Name         string `form:"name" binding:"required, min=2, max=100"`
	Phone        string `form:"phone" binding:"required, min=10, max=20"`
	Address      string `form:"address" binding:"required, min=10, max=100"`
	Sizes        string `form:"size" binding:"required, min=1, dive, valid_pizza_size"`
	Types        string `form:"type" binding:"required, min=1, dive, valid_pizza_type"`
	Instructions string `form:"instructions" binding:"max=200"`
}

func (h *Handler) ServeNewOrderPost(c *gin.Context) {
	c.HTML(http.StatusOK, "order.tmpl", OrderFormData{
		PizzaTypes: models.PizzaTypes,
		PizzaSizes: models.PizzaSizes,
	})
}

func (h *Handler) HandleNewOrderPost(c *gin.Context) {
	var form OrderRequest

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderItems := make([]models.OrderItem, len(form.Sizes))
	for i := range orderItems {
		orderItems[i] = models.OrderItem{
			Size:         form.Sizes[i],
			Pizza:        form.PizzaTypes[i],
			Instructions: form.Instructions[i],
		}
	}

	order := models.Order{
		CustomerName: form.Name,
		Phone:        form.Phone,
		Address:      form.Address,
		Status:       models.OrderStatus[0],
		Items:        orderItems,
	}

	if err := h.orders.CreateOrder(&order); err != nil {
		slog.Error("Failed to create order", "error", err)
		c.String(http.StatusInternalServerError, "Something went wrong!")
		return
	}
}
