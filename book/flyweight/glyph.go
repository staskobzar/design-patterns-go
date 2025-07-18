package flyweight

import "fmt"

type Char struct {
	c      rune
	font   string
	weight int
	size   int
}

type GlyphFactory struct {
	chars map[rune]*Char
}

func CreateFactory() *GlyphFactory {
	g := &GlyphFactory{}
	g.chars = make(map[rune]*Char)
	for c := 'a'; c <= 'z'; c++ {
		char := &Char{c, "d", 100, 10}
		g.chars[c] = char
	}
	return g
}

func (g *GlyphFactory) SetSize(s int) {
	for _, char := range g.chars {
		char.size = s
	}
}

func (g *GlyphFactory) SetWeight(w int) {
	for _, char := range g.chars {
		char.weight = w
	}
}

func (g *GlyphFactory) SetFont(font string) {
	for _, char := range g.chars {
		char.font = font
	}
}

func (g *GlyphFactory) Print(input string) string {
	var output string
	for _, r := range input {
		c := g.chars[r]
		output = fmt.Sprintf("%s[-%s-%d-%d]%c",
			output, c.font, c.weight, c.size, c.c)
	}
	return output
}
