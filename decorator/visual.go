package main

import "fmt"

type Decorator interface {
	Draw()
}

type Window struct {
	draw      string
	component Decorator
}

func (w *Window) Draw() {
	fmt.Println(w.draw)
}

func (w *Window) SetComponent(c Decorator) {
	w.Draw()
	c.Draw()
}

func NewWindow() *Window {
	return &Window{draw: "Draw Window"}
}

type VisualComponent struct {
	draw      string
	component Decorator
}

func (c *VisualComponent) Draw() {
	fmt.Println(c.draw)
	c.component.Draw()
}

type BorderDecorator struct {
	VisualComponent
}

func NewBorderDecorator(c Decorator, width int) Decorator {
	return &BorderDecorator{
		VisualComponent{
			draw:      fmt.Sprintf("=> | Draw Border width %d |", width),
			component: c,
		},
	}
}

type ScrollDecorator struct {
	VisualComponent
}

func NewScrollDecorator(c Decorator) Decorator {
	return &ScrollDecorator{
		VisualComponent{
			draw:      "==> ^ Draw Scroll ^",
			component: c,
		},
	}
}

type TextView struct {
	draw string
}

func (tv *TextView) Draw() {
	fmt.Println(tv.draw)
}

func NewTextView() Decorator {
	return &TextView{"===> Draw Text View"}
}
