package state

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ifToName(i interface{}) string {
	return reflect.TypeOf(i).Elem().Name()
}

func TestActiveOpen(t *testing.T) {
	c := NewTCPConnection()

	assert.Equal(t, "TCPClosed", ifToName(c.state))

	c.ActiveOpen()
	assert.Equal(t, "TCPEstablised", ifToName(c.state))

	c.Close()
	assert.Equal(t, "TCPClosed", ifToName(c.state))
}

func TestPassiveOpen(t *testing.T) {
	c := NewTCPConnection()

	assert.Equal(t, "TCPClosed", ifToName(c.state))

	c.PassiveOpen()
	assert.Equal(t, "TCPListen", ifToName(c.state))

	c.Close()
	assert.Equal(t, "TCPClosed", ifToName(c.state))
}
