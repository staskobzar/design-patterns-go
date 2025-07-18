package strategy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBilling(t *testing.T) {
	alice := NewCustomerBill(&NormalStrategy{})
	alice.Add(5.0, 1)
	alice.ChangeStrategy(&HappyHourStrategy{})
	alice.Add(3.0, 2)

	assert.Equal(t, "8.00$", alice.Bill())

	bob := NewCustomerBill(&HappyHourStrategy{})
	bob.Add(0.8, 1)
	bob.ChangeStrategy(&NormalStrategy{})
	bob.Add(1.3, 2)
	bob.Add(2.5, 1)
	assert.Equal(t, "5.50$", bob.Bill())
}
