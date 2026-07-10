package main

type Character struct {
	weapon Weapon
}

func (c *Character) Fight() {
	c.weapon.useWeapon()
}

func createKing() *Character {
	return &Character{weapon: &Bow{}}
}

func createQueen() *Character {
	return &Character{weapon: &Knife{}}
}

func createKnight() *Character {
	return &Character{weapon: &Sword{}}
}

func createTroll() *Character {
	return &Character{weapon: &Axe{}}
}
