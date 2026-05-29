package models

// ProductOrder represents a line item within an order.
// ProductPrice captures the product price at the time the order was placed.
type ProductOrder struct {
	OrderID      int `json:"orderId" db:"order_id"`
	ProductID    int `json:"productId" db:"product"`
	Amount       int `json:"amount" db:"amount"`
	ProductPrice int `json:"productPrice" db:"product_price"`
}

// CreateProductOrder creates a new ProductOrder line item.
func CreateProductOrder(orderID, productID, amount, productPrice int) ProductOrder {
	return ProductOrder{
		OrderID:      orderID,
		ProductID:    productID,
		Amount:       amount,
		ProductPrice: productPrice,
	}
}
