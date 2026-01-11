//  enumerated types

package main

import "fmt"

type OrderStatus int
type AccountStatus string

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

const (
	Active    AccountStatus = "active"
	Block     AccountStatus = "block"
	Suspended AccountStatus = "suspended"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func getAccountStatus(status AccountStatus) {
	fmt.Println("Your account status is:", status)
}

func main() {
	changeOrderStatus(Delivered)
	getAccountStatus(Active)
}
