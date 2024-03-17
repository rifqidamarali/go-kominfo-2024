package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Order struct {
	Order_ID      uint   `json:"order_id"`
	Customer_Name string `json:"customer_name"`
	Ordered_At    int    `json:"ordered_at"`
}

var orders = []Order{}

func CreateOrder(ctx *gin.Context) {
	var newOrder Order
	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newOrder.Order_ID = uint(len(orders) + 1)
	orders = append(orders, newOrder)

	ctx.JSON(http.StatusCreated, gin.H{
		"order": newOrder,
	})


}