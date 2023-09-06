package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	// Create channels for receiving, validating, and handling orders.
	var receivedOrderCh = make(chan Order)
	var validOrderCh = make(chan Order)
	var invalidOrderCh = make(chan invalidOrder)
	var wg sync.WaitGroup

	// Start goroutines to handle orders.
	go recievedOrders(receivedOrderCh)
	go validateOrders(receivedOrderCh, validOrderCh, invalidOrderCh)

	wg.Add(1)
	go func() {
		order := <-validOrderCh
		fmt.Printf("Valid Order received: %+v\n", order)
		wg.Done()
	}()
	go func() {
		order := <-invalidOrderCh
		fmt.Printf("Invalid Order received: %+v. Issue: %v\n", order.order, order.err)
		wg.Done()
	}()
	wg.Wait()

	elapsed := time.Since(startTime)
	fmt.Println("Time elapsed in seconds:", elapsed)
}

// validateOrders validates orders and sends them to either the valid or invalid channels.
func validateOrders(in, out chan Order, errCh chan invalidOrder) {
	order := <- in
	if order.Price <= 0 {
		errCh <- invalidOrder{order: order, err:
			errors.New("price can't be 0 or less")}
	} else {
		out <- order
	}
}

// receiverOrders simulates receiving orders and sends them to the input channel.
func recievedOrders(out chan Order) {
	for _, rawOrder := range rawOrders {
		var newOrder Order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil{
			log.Println(err)
			continue
		}
		out <- newOrder
	}
}

var rawOrders = []string {
    `{"ProductID":1, "Quantity":1, "Price":100, "Status": 1}`,
    `{"ProductID":2, "Quantity":2.5, "Price":300.55, "Status": 3}`,
    `{"ProductID":3, "Quantity":42.6, "Price":10.6, "Status": 1}`,
    `{"ProductID":4, "Quantity":10.5, "Price":256, "Status": 1}`,
    `{"ProductID":5, "Quantity":20.6, "Price":999, "Status": 1}`,
}