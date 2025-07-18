package mediator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMediator(t *testing.T) {
	btnOk := &Button{true, "OK", WBtnOK, nil}
	btnCancel := &Button{false, "Cancel", WBtnCancel, nil}
	list := NewListBox()
	field := &EntryField{"", false, WEntryField, nil}

	dialog := NewDialog(btnOk, btnCancel, list, field)

	assert.False(t, btnCancel.Enabled)
	assert.Empty(t, field.Text)
	assert.False(t, field.Active)
	assert.Empty(t, list.Selected)

	dialog.SetList(2)
	assert.True(t, btnCancel.Enabled)
	assert.Equal(t, ListVals[2], field.Text)
	assert.True(t, field.Active)
	assert.Equal(t, ListVals[2], list.Selected)
}
