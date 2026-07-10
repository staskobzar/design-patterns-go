package main

import "fmt"

func main() {
	fmt.Println("=> Advanture Game Start")

	king := createKing()
	queen := createQueen()
	knight := createKnight()
	troll := createTroll()

	king.Fight()
	troll.Fight()
	queen.Fight()
	troll.Fight()
	knight.Fight()
	fmt.Println("[!] Troll is defeated by the knight!")
}
