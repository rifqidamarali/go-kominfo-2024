package controllers

type Item struct {
	Item_ID     uint   `json:"id"`
	Item_Code   uint   `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}