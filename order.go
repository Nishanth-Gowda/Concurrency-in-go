package main

import "fmt"

//Container various attrributes of the order
type Order struct{
	ProductID int
	Quantity float64
	Price float64
	Status OrderStatus
}

type invalidOrder struct{
	order Order
	err error
}

type OrderStatus int


const (
	none OrderStatus = iota
	pending
	fulfilled
	rejected
	cancelled
)


func (o Order) String() string {
	return fmt.Sprintf("ProductID: %v, Quantity: %v, Price: %v, Status: %v\n", o.ProductID, o.Quantity, o.Price, orderStatusToText(o.Status))
}

func orderStatusToText(o OrderStatus) string {
	switch o {
	case none:
		return "none"
	case pending:
		return "pending"
	case fulfilled:
		return "fulfilled"
	case rejected:
		return "rejected"
	case cancelled:
		return "cancelled"
	default:
		return "unknown"
	}
}