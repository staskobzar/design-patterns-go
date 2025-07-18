package main

func Example_window_TextViewOnly() {
	window := NewWindow()
	textView := NewTextView()

	window.SetComponent(textView)

	// Output:
	// Draw Window
	// ===> Draw Text View
}

func Example_window_Decorators() {
	window := NewWindow()
	textView := NewTextView()

	window.SetComponent(
		NewBorderDecorator(
			NewScrollDecorator(textView), 2,
		),
	)

	// Output:
	// Draw Window
	// => | Draw Border width 2 |
	// ==> ^ Draw Scroll ^
	// ===> Draw Text View
}
