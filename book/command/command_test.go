package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandOpen(t *testing.T) {
	app := NewApp()
	assert.Empty(t, app.doc)

	NewCommandOpen(app, "New document")

	app.Open()
	assert.Equal(t, "New document", app.doc)
}

func TestCommandPaste(t *testing.T) {
	app := NewApp()
	assert.Empty(t, app.doc)

	NewCommandOpen(app, "New document")
	NewCommandPaste(app, "more text")

	app.Open()
	app.Paste()

	assert.Equal(t, "New document more text", app.doc)
}

func TestCommandMacro(t *testing.T) {
	app := NewApp()

	NewCommandOpen(app, "New document")
	NewCommandPaste(app, "more text")

	app.Add(&CmdDelWord{app, 0})
	app.Add(&CmdAppend{app, "foo"})
	app.Add(&CmdPrepend{app, "Bar"})

	app.Open()
	app.Paste()
	app.Macro()
	assert.Equal(t, "Bar document more text foo", app.doc)
}
