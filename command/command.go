package command

import (
	"fmt"
	"strings"
)

type Command interface {
	Execute()
}

type App struct {
	doc      string
	cmdOpen  Command
	cmdPaste Command
	macro    []Command
}

func NewApp() *App {
	return &App{macro: make([]Command, 0)}
}

func (a *App) Open() {
	a.cmdOpen.Execute()
}

func (a *App) Paste() {
	a.cmdPaste.Execute()
}

func (a *App) Add(c Command) {
	a.macro = append(a.macro, c)
}

func (a *App) Macro() {
	for _, cmd := range a.macro {
		cmd.Execute()
	}
}

type CommandOpen struct {
	app  *App
	text string
}

func NewCommandOpen(app *App, t string) {
	app.cmdOpen = &CommandOpen{app, t}
}

func (c *CommandOpen) Execute() {
	c.app.doc = c.text
}

type CommandPaste struct {
	app  *App
	text string
}

func NewCommandPaste(app *App, t string) {
	app.cmdPaste = &CommandPaste{app, t}
}

func (c *CommandPaste) Execute() {
	c.app.doc = fmt.Sprintf("%s %s", c.app.doc, c.text)
}

type CmdDelWord struct {
	app   *App
	index int
}

func (c *CmdDelWord) Execute() {
	words := strings.Split(c.app.doc, " ")
	result := append(words[:c.index], words[c.index+1:]...)
	c.app.doc = strings.Join(result, " ")
}

type CmdAppend struct {
	app  *App
	word string
}

func (c *CmdAppend) Execute() {
	c.app.doc = fmt.Sprintf("%s %s", c.app.doc, c.word)
}

type CmdPrepend struct {
	app  *App
	word string
}

func (c *CmdPrepend) Execute() {
	c.app.doc = fmt.Sprintf("%s %s", c.word, c.app.doc)
}
