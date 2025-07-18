package strategy

import "fmt"

type Price float32
type BillingStratagy interface {
	GetActPrice(price Price) Price
}

type CustomerBill struct {
	drinks   []Price
	strategy BillingStratagy
}

func NewCustomerBill(s BillingStratagy) *CustomerBill {
	return &CustomerBill{drinks: make([]Price, 0), strategy: s}
}

func (b *CustomerBill) ChangeStrategy(s BillingStratagy) {
	b.strategy = s
}

func (b *CustomerBill) Add(price Price, quantity int) {
	cost := price * Price(quantity)
	b.drinks = append(b.drinks, b.strategy.GetActPrice(cost))
}

func (b *CustomerBill) Bill() string {
	var cost Price
	for _, price := range b.drinks {
		cost += price
	}
	return fmt.Sprintf("%.2f$", cost)
}

type NormalStrategy struct{}

func (s *NormalStrategy) GetActPrice(price Price) Price {
	return price
}

type HappyHourStrategy struct{}

func (s *HappyHourStrategy) GetActPrice(price Price) Price {
	return price / 2.0
}
