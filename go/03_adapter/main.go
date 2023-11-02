package main

import "fmt"

type PaymentGateway interface {
	processPayment()
}

// stripe
type stripe struct{}

func (s *stripe) processPayment() {
	fmt.Println("stripe payment gateway - process payment")
}

// payPal
type payPal struct{}

func (w *payPal) makePayment() {
	fmt.Println("payPal payment gateway - make payment")
}

// paymentAdapter
type paymentAdapter struct {
	payPalAdapter *payPal
}

func (p *paymentAdapter) processPayment() {
	p.payPalAdapter.makePayment()
}

// client
type client struct{}

func (c *client) clientMakingPayment(pg PaymentGateway) {
	pg.processPayment()
}

func main() {
	client := client{}

	stripe := &stripe{}

	client.clientMakingPayment(stripe)

	payPal := &payPal{}

	paymentAdapter := &paymentAdapter{
		payPalAdapter: payPal,
	}

	// using client we can call processPayment for paypal
	client.clientMakingPayment(paymentAdapter)
}
