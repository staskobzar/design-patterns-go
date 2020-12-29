package observer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestObserver(t *testing.T) {
	timer := NewClockTimer(time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC))
	analogClock := NewAnalogClock(timer)
	digitalClock := NewDigitalClock(timer)

	assert.Equal(t, 12, timer.time.Hour())
	assert.Equal(t, 0, timer.time.Minute())
	assert.Equal(t, 0, timer.time.Second())

	assert.Equal(t, 12, analogClock.Hour)
	assert.Equal(t, 0, analogClock.Minute)
	assert.Equal(t, 0, analogClock.Second)
	assert.Equal(t, "ANALOG(12:00:00)", analogClock.String())

	assert.Equal(t, 12, digitalClock.Hour)
	assert.Equal(t, 0, digitalClock.Minute)
	assert.Equal(t, 0, digitalClock.Second)
	assert.Equal(t, "DIGITAL[ 12:00:00 PM ]", digitalClock.String())

	timer.Tick(100)
	assert.Equal(t, 12, timer.time.Hour())
	assert.Equal(t, 1, timer.time.Minute())
	assert.Equal(t, 40, timer.time.Second())
	assert.Equal(t, "ANALOG(12:01:40)", analogClock.String())
	assert.Equal(t, "DIGITAL[ 12:01:40 PM ]", digitalClock.String())

	analogClock.Stop()
	timer.Tick(100)
	assert.Equal(t, 12, timer.time.Hour())
	assert.Equal(t, 3, timer.time.Minute())
	assert.Equal(t, 20, timer.time.Second())
	assert.Equal(t, "ANALOG(12:01:40)", analogClock.String())
	assert.Equal(t, "DIGITAL[ 12:03:20 PM ]", digitalClock.String())

	digitalClock.Stop()
	timer.Tick(100)
	assert.Equal(t, 12, timer.time.Hour())
	assert.Equal(t, 5, timer.time.Minute())
	assert.Equal(t, 0, timer.time.Second())
	assert.Equal(t, "ANALOG(12:01:40)", analogClock.String())
	assert.Equal(t, "DIGITAL[ 12:03:20 PM ]", digitalClock.String())
}
