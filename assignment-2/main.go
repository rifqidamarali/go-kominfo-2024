package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	Item_ID     uint   `json:"id"`
	Item_Code   uint   `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}

type Order struct {
	Order_ID      uint   `json:"order_id"`
	Customer_Name string `json:"customer_name"`
	Ordered_At    int    `json:"ordered_at"`
}

func main() {
	engine := gin.new()
	engine.LoadHTMLGlob("template/*")
	engine.GET("/template/index/:name", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", map[string]any{
			"title": ctx.Param("name"),
		})
	})
}