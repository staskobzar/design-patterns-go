package visitor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewElement(t *testing.T) {
	e := NewElement("DDR4-3200 4G", 4, 16.99, 2.00)
	assert.Equal(t, "DDR4-3200 4G", e.Name)
	assert.EqualValues(t, 4, e.Power)
	assert.EqualValues(t, 1699, e.Price)
	assert.EqualValues(t, 200, e.Discount)
}

func TestNewComposite(t *testing.T) {
	c := NewComposite("ASUS-B460 PRO Intel", 12, 109.99, 0)
	assert.Equal(t, "ASUS-B460 PRO Intel", c.Name)
	assert.EqualValues(t, 12, c.Power)
	assert.EqualValues(t, 10999, c.Price)
	assert.EqualValues(t, 0, c.Discount)
	assert.Equal(t, 0, c.list.Count())
	assert.True(t, c.iter.IsDone())
}

func TestCompositeAdd(t *testing.T) {
	board := NewComposite("ASUS-B460 PRO Intel", 12, 109.99, 0)
	ram := NewElement("DDR4-3200 4G", 4, 16.99, 2.00)
	board.Add(ram)
	assert.Equal(t, 1, board.list.Count())

	assert.Equal(t, ram, board.iter.CurrentItem())
}

func stubPC() Equipment {
	mboard := NewMotherBoard("Asus B460 PRO Intel", 109.99)

	mboard.Add(NewCPU("Core i9-10900K 3.7GHz", 499.99))
	mboard.Add(NewRAM("Ripjaws V 4GB DDR4-3200", 14.99))
	mboard.Add(NewRAM("Ripjaws V 4GB DDR4-3200", 14.99))
	mboard.Add(NewRAM("Ripjaws V 4GB DDR4-3200", 14.99))
	mboard.Add(NewRAM("Ripjaws V 4GB DDR4-3200", 14.99))

	mboard.Add(NewHDD("BarraCuda 1TB", 47.49))
	mboard.Add(NewHDD("WD Black 2TB", 114.99))

	mboard.Add(NewVideoCard("Radeon 7750 2G", 179.99))

	pc := NewCabinet("eATX Full Tower", 79.99)
	pc.Add(mboard)
	return pc
}

func TestVisitorGPower(t *testing.T) {
	pc := stubPC()

	visitor := NewVisitorGPower()
	pc.Accept(visitor)

	assert.Equal(t, 174, visitor.GPower())
}

func Example_VisitorBill() {
	pc := stubPC()
	visitor := NewVisitorBill()
	pc.Accept(visitor)

	visitor.PrintBill()
	// Output:
	// Name                               Price
	// CPU: Core i9-10900K 3.7GHz         499.99$
	// RAM: Ripjaws V 4GB DDR4-3200       14.99$
	// RAM: Ripjaws V 4GB DDR4-3200       14.99$
	// RAM: Ripjaws V 4GB DDR4-3200       14.99$
	// RAM: Ripjaws V 4GB DDR4-3200       14.99$
	// HDD: BarraCuda 1TB                 47.49$
	// HDD: WD Black 2TB                  114.99$
	// Video Card: Radeon 7750 2G         179.99$
	// Mother Board: Asus B460 PRO Intel  109.99$
	// Cabinet: eATX Full Tower           79.99$
	//
	// Discount:    76.75$
	// Total:       1092.40$
	// Amount Due:  1015.65$
}
