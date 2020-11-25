package equipment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEquipmentComposite(t *testing.T) {
	cabinet := NewCabinet("PC Cabinet") // price $9.75, discount $1.50, W10
	chassis := NewChassis("PC Chassis") // price $15.25, discount $0.00, W10

	cabinet.Add(chassis)

	bus := NewBus("MCA Bus")             // price $32.40, discount $4.10, W12
	bus.Add(NewCard("16Mbs Token Ring")) // price $26.83, discount $3.11, W14

	chassis.Add(bus)
	chassis.Add(NewFloppyDisk("3.5in Floppy")) // price $2.68, discount $0.43, W10

	assert.EqualValues(t, 86.91, cabinet.NetPrice())
	assert.EqualValues(t, 9.14, cabinet.Discount())
	assert.EqualValues(t, 56, cabinet.Power())
	assert.Equal(t, "$77.77", cabinet.Price())
}
