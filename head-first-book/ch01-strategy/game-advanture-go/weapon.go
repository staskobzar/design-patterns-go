package main

import "fmt"

type Weapon interface {
	useWeapon()
}

type Sword struct{}

func (s *Sword) useWeapon() {
	fmt.Println("Swinging a sword!")
}

type Axe struct{}

func (a *Axe) useWeapon() {
	fmt.Println("Chopping with an axe!")
}

type Knife struct{}

func (k *Knife) useWeapon() {
	fmt.Println("Cutting with a knife!")
}

type Bow struct{}

func (b *Bow) useWeapon() {
	fmt.Println("Shooting an arrow with a bow!")
}
