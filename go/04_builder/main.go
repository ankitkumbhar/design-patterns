package main

import "fmt"

type Response struct {
	User    User
	Account Account
	Address Address
}

type (
	User struct {
		Email string
	}

	Account struct {
		Balance float64
	}

	Address struct {
		City string
	}
)

// Map the response data to response struct using builder pattern
type ResponseBuilder struct {
	Response Response
}

// NewBuilder - returns a pointer to a ResponseBuilder
func NewBuilder() *ResponseBuilder {
	return &ResponseBuilder{
		Response: Response{},
	}
}

// Add helper methods to the builder for constructing the repsonse
func (rb *ResponseBuilder) SetEmail(email string) {
	rb.Response.User.Email = email
}

func (rb *ResponseBuilder) SetBalance(balance float64) {
	rb.Response.Account.Balance = balance
}

func (rb *ResponseBuilder) SetCity(city string) {
	rb.Response.Address.City = city
}

func (rb *ResponseBuilder) Build() Response {
	return rb.Response
}

type Director struct {
	builder *ResponseBuilder
}

func (d *Director) constructResponse(email string, amount float64, city string) Response {
	d.builder.SetEmail(email)
	d.builder.SetBalance(amount)
	d.builder.SetCity(city)

	return d.builder.Build()
}

func main() {
	rb := NewBuilder()

	d := Director{builder: rb}

	resp := d.constructResponse("abcd@abcd.com", 100, "Pune")

	fmt.Println("Response : ", resp)
}
