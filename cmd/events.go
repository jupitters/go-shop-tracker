package main

import "github.com/gin-gonic/gin"

func (h *Handler) NotificationHandler(c *gin.Context) {
	orderID := c.Query("orderId")
	if orderID == "" {
		c.String(400, "Invalid orderId")
		return
	}

	_, err := h.orders.GetOrder(orderID)
	if err != nil {
		c.String(404, "Order not found")
		return
	}

	key := "order:" + orderID
	client := make(chan string, 10)
}
