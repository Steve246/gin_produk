package model

type Product struct {
	ProductId   string `json: "productId"`
	ProductName string `json: "productName"`
	IsStatus    bool   `json: "isStatus"`
}
