package template

func ExampleMarioGame() {
	game := NewMarioGame()
	game.Play()

	// Output:
	// Mario Game Initialized! Start playing.
	// Mario Game Started. Enjoy the game!
	// Mario Game Finished!
}

func ExampleTankfightGame() {
	game := NewTankfightGame()
	game.Play()

	// Output:
	// Tankfight Game Initialized! Start playing.
	// Tankfight Game Started. Enjoy the game!
	// Tankfight Game Finished!
}
