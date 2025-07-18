package main

import (
	"fmt"
	"log"
)

// base class
type OrderProcessor struct {
	validateOrder    func()
	calculateTotal   func()
	processPayment   func()
	deliverOrder     func()
	sendConfirmation func()
}

func InitOrderProcessor() *OrderProcessor {
	return &OrderProcessor{
		validateOrder:    func() { log.Println("[v] Validating order...") },
		calculateTotal:   func() { log.Println("[v] Calculate total...") },
		processPayment:   func() { log.Println("[v] Process payment...") },
		deliverOrder:     func() { log.Println("[v] Deliver order...") },
		sendConfirmation: func() { log.Println("[v] Send confirmation...") },
	}
}

func (op *OrderProcessor) processOrder() {
	op.validateOrder()
	op.calculateTotal()
	op.processPayment()
	op.deliverOrder()
	op.sendConfirmation()
}

// subclasses
type PhisicalOrder struct {
	*OrderProcessor
}

func InitPhisicalOrder() *PhisicalOrder {
	po := &PhisicalOrder{InitOrderProcessor()}
	po.deliverOrder = func() { log.Println("[+] Shipping phisical item via courier...") }
	return po
}

type DigitalOrder struct {
	*OrderProcessor
}

func InitDigitalOrder() *DigitalOrder {
	do := &DigitalOrder{InitOrderProcessor()}
	do.deliverOrder = func() { log.Println("[*] Deliver digital content via download link...") }
	do.calculateTotal = func() { log.Println("[*] Calculate total price (no phisical shipping)...") }
	return do
}

func main() {
	fmt.Println("=> New Phisical order")
	InitPhisicalOrder().processOrder()

	fmt.Println("=> New Digital order")
	InitDigitalOrder().processOrder()
}
