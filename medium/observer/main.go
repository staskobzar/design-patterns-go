package main

import (
	"fmt"
	"slices"
)

// Obserever interface
type BlogSubscriber interface {
	update(title string)
}

// Subject
type Blog struct {
	observers []BlogSubscriber
}

func NewBlog() *Blog {
	return &Blog{
		observers: make([]BlogSubscriber, 0),
	}
}

func (b *Blog) subscribe(observer BlogSubscriber) {
	b.observers = append(b.observers, observer)
}

func (b *Blog) unsubscribe(observer BlogSubscriber) {
	fmt.Printf("[-] unsubscribe observer")
	b.observers = slices.DeleteFunc(b.observers, func(user BlogSubscriber) bool {
		return user == observer
	})
}

func (b *Blog) publish(title string) {
	fmt.Printf("[+] publish new article: %q\n", title)
	for _, observer := range b.observers {
		observer.update(title)
	}
}

// Observers
type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{name: name}
}

func (r *Reader) update(title string) {
	fmt.Printf("[+] reader %s was notified about new article: %q\n", r.name, title)
}

func main() {
	fmt.Println("=== OBSERVER")
	blog := NewBlog()
	alice := NewReader("Alice")
	bob := NewReader("Bob")
	carl := NewReader("Carl")

	blog.subscribe(alice)
	blog.subscribe(bob)
	blog.subscribe(carl)

	blog.publish("Understanding the observer Pattern")

	fmt.Println()
	blog.unsubscribe(bob)
	blog.publish("Go Observer Pattern in Practice")

	fmt.Println()
	blog.unsubscribe(carl)
	blog.publish("Go patterns for dummies")
}
