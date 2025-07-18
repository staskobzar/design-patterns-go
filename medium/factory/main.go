package main

import "fmt"

// ============================ Notify: START
// Notifier interface
type Notifier interface {
	send(msg string)
}

// concrete Product
type PushNotifier struct{}

func (p *PushNotifier) send(msg string) {
	fmt.Printf("PUSH: %q\n", msg)
}

type SMSNotifier struct{}

func (p *SMSNotifier) send(msg string) {
	fmt.Printf("SMS: %q\n", msg)
}

type EmailNotifier struct{}

func (p *EmailNotifier) send(msg string) {
	fmt.Printf("EMAIL: %q\n", msg)
}

// Creator Interface
type Creator interface {
	createNotifier() Notifier
}

// concrete Creators
type PushCreator struct{}

func (*PushCreator) createNotifier() Notifier { return &PushNotifier{} }

type SMSCreator struct{}

func (*SMSCreator) createNotifier() Notifier { return &SMSNotifier{} }

type EmailCreator struct{}

func (*EmailCreator) createNotifier() Notifier { return &EmailNotifier{} }

// ============================ Notify: END

// ============================ Logistic: END
// product interface
type Transport interface {
	deliver()
}

// concrete product
type Ship struct{}

func (*Ship) deliver() { fmt.Println("deliver by sea") }

type Truck struct{}

func (*Truck) deliver() { fmt.Println("deliver by land") }

type Logistic interface {
	createTransport() Transport
}

type RoadLogistic struct{}

func (*RoadLogistic) createTransport() Transport { return &Truck{} }

type SeaLogistic struct{}

func (*SeaLogistic) createTransport() Transport { return &Ship{} }

// ============================ Logistic: END

func main() {
	fmt.Println("====== Factory Notifier =======")
	notifiers := []Creator{&PushCreator{}, &SMSCreator{}, &EmailCreator{}}
	for _, notifier := range notifiers {
		notifier.createNotifier().send("new message")
	}

	fmt.Println("====== Factory Transport =======")
	logistics := []Logistic{&RoadLogistic{}, &SeaLogistic{}}

	for _, l := range logistics {
		l.createTransport().deliver()
	}
}
