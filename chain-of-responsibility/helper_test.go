package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplication(t *testing.T) {
	a := NewApplication(NO_HELP_TOPIC)

	assert.Equal(t, "", a.HandleHelp())

	a = NewApplication(TOPIC_APPLICATION)
	assert.Equal(t, "Application Help", a.HandleHelp())
}

func TestDialog(t *testing.T) {
	a := NewApplication(TOPIC_APPLICATION)
	d := NewDialog(a, TOPIC_DIALOG)

	assert.True(t, d.HasHelp())
	assert.Equal(t, "Dialog Help", d.HandleHelp())
}

func TestButton(t *testing.T) {
	a := NewApplication(TOPIC_APPLICATION)
	b := NewButton(a, TOPIC_BUTTON)

	assert.True(t, b.HasHelp())
	assert.Equal(t, "Button Help", b.HandleHelp())
}

func TestChainToButton(t *testing.T) {
	a := NewApplication(TOPIC_APPLICATION)
	d := NewDialog(a, TOPIC_DIALOG)
	b := NewButton(d, TOPIC_BUTTON)
	assert.Equal(t, "Button Help", b.HandleHelp())
}

func TestChainToDialog(t *testing.T) {
	a := NewApplication(TOPIC_APPLICATION)
	d := NewDialog(a, TOPIC_DIALOG)
	b := NewButton(d, NO_HELP_TOPIC)
	assert.Equal(t, "Dialog Help", b.HandleHelp())
}

func TestChainToApplication(t *testing.T) {
	a := NewApplication(TOPIC_APPLICATION)
	d := NewDialog(a, NO_HELP_TOPIC)
	b := NewButton(d, NO_HELP_TOPIC)
	assert.Equal(t, "Application Help", b.HandleHelp())
}
