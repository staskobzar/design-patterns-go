package main

import "fmt"

// strategy interface
type PaymentStrategy interface {
	pay(amount float32)
}

// concrete strategy implementation
type CreditCard struct {
	name   string
	number string
}

func CreateCreditCard(name, num string) *CreditCard {
	return &CreditCard{name: name, number: num}
}

func (cc *CreditCard) pay(amount float32) {
	fmt.Printf("Paying $%0.2f using Credit Card [%s] for %q\n",
		amount, cc.number, cc.name)
}

// concrete strategy implementation
type PayPall struct {
	email string
}

func CreatePayPal(email string) *PayPall {
	return &PayPall{email: email}
}

func (pp *PayPall) pay(amount float32) {
	fmt.Printf("Paying %.2f using PayPal account [%s]\n", amount, pp.email)
}

// strategy context
type Order struct {
	paymentStrategy PaymentStrategy
}

func NewOrder(ps PaymentStrategy) *Order {
	return &Order{paymentStrategy: ps}
}

func (o *Order) checkout(amount float32) {
	fmt.Printf("Processing order of $%.2f\n", amount)
	o.paymentStrategy.pay(amount)
	fmt.Println("Payment complete...")
}

func main() {
	NewOrder(CreateCreditCard("John Doe", "4111-1111-1111-1111")).checkout(249.99)
	NewOrder(CreatePayPal("john.doe@example.com")).checkout(89.5)
	fmt.Println("--------------------------------------------------")
}
