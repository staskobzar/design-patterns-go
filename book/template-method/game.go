package template

import "fmt"

// example from wikipedia article
// https://en.wikipedia.org/wiki/Template_method_pattern

type Template interface {
	Initialize()
	StartPlay()
	EndPlay()
}
type IGame interface {
	Template
	Play()
}

type Game struct{ Template }

func (g *Game) Play() {
	g.Initialize()

	g.StartPlay()

	g.EndPlay()
}

type MarioGame struct{}

func NewMarioGame() IGame {
	return &Game{&MarioGame{}}
}

func (*MarioGame) Initialize() {
	fmt.Println("Mario Game Initialized! Start playing.")
}

func (*MarioGame) StartPlay() {
	fmt.Println("Mario Game Started. Enjoy the game!")
}

func (*MarioGame) EndPlay() {
	fmt.Println("Mario Game Finished!")
}

type TankfightGame struct{}

func NewTankfightGame() IGame {
	return &Game{&TankfightGame{}}
}

func (*TankfightGame) Initialize() {
	fmt.Println("Tankfight Game Initialized! Start playing.")
}

func (*TankfightGame) StartPlay() {
	fmt.Println("Tankfight Game Started. Enjoy the game!")
}

func (*TankfightGame) EndPlay() {
	fmt.Println("Tankfight Game Finished!")
}
