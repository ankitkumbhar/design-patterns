package main

import "fmt"

type PaymentGatewayType int

const (
	GatewayTypeStripe PaymentGatewayType = 1
	GatewayTypePaypal PaymentGatewayType = 2
)

type PaymentGateway interface {
	ProcessPayment(amount float64)
}

type StripeGateway struct{}

func (*StripeGateway) ProcessPayment(amount float64) {
	fmt.Println("Processing payment using Stripe Gateway : ", amount)
}

type PaypalGateway struct{}

func (*PaypalGateway) ProcessPayment(amount float64) {
	fmt.Println("Processing payment using PayPal Gateway : ", amount)
}

// factoryMethod - instantiate a class bsaed on the passed parameter
func FactoryMethod(gatewayType PaymentGatewayType) PaymentGateway {
	switch gatewayType {
	case GatewayTypeStripe:
		return &StripeGateway{}

	case GatewayTypePaypal:
		return &PaypalGateway{}
	}

	return nil
}

func main() {
	// object of stripe is created
	paymentObj := FactoryMethod(GatewayTypeStripe)
	paymentObj.ProcessPayment(10.10)

	// object of payPal is created
	paymentObj = FactoryMethod(GatewayTypePaypal)
	paymentObj.ProcessPayment(10.20)
}
