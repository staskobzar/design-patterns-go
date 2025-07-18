package observer

import (
	"fmt"
	"time"
)

type IObserver interface {
	Update(ISubject)
}

type ISubject interface {
	Attach(IObserver)
	Detach(IObserver)
	Notify()
}

type Subject struct {
	observers []IObserver
}

func (s *Subject) Attach(o IObserver) {
	s.observers = append(s.observers, o)
}

func (s *Subject) Detach(o IObserver) {
	for i := 0; i < len(s.observers); i++ {
		if o == s.observers[i] {
			copy(s.observers[i:], s.observers[i+1:])
			s.observers = s.observers[:len(s.observers)-1]
			return
		}
	}
}

type ClockTimer struct {
	*Subject
	time time.Time
}

func NewClockTimer(t time.Time) *ClockTimer {
	sub := &Subject{observers: make([]IObserver, 0)}
	return &ClockTimer{sub, t}
}

func (t *ClockTimer) Notify() {
	for i := 0; i < len(t.observers); i++ {
		t.observers[i].Update(t)
	}
}

func (t *ClockTimer) Hour() int { return t.time.Hour() }
func (t *ClockTimer) Min() int  { return t.time.Minute() }
func (t *ClockTimer) Sec() int  { return t.time.Second() }
func (t *ClockTimer) Tick(sec int) {
	t.time = t.time.Add(time.Second * time.Duration(sec))
	t.Notify()
}

type Clock struct {
	subject ISubject
	Hour    int
	Minute  int
	Second  int
}

func (c *Clock) Stop() {
	c.subject.Detach(c)
}

func (c *Clock) Update(s ISubject) {
	timer := s.(*ClockTimer)
	c.Hour = timer.Hour()
	c.Minute = timer.Min()
	c.Second = timer.Sec()
}

type AnalogClock struct{ *Clock }

func (ac *AnalogClock) String() string {
	return fmt.Sprintf("ANALOG(%02d:%02d:%02d)", ac.Hour, ac.Minute, ac.Second)
}

type DigitalClock struct{ *Clock }

func (dc *DigitalClock) String() string {
	suffix := "AM"
	if dc.Hour >= 12 {
		suffix = "PM"
	}
	return fmt.Sprintf("DIGITAL[ %02d:%02d:%02d %s ]", dc.Hour, dc.Minute, dc.Second, suffix)
}

func NewAnalogClock(t *ClockTimer) *AnalogClock {
	clock := &Clock{t, t.Hour(), t.Min(), t.Sec()}
	clock.subject.Attach(clock)
	return &AnalogClock{clock}
}

func NewDigitalClock(t *ClockTimer) *DigitalClock {
	clock := &Clock{t, t.Hour(), t.Min(), t.Sec()}
	clock.subject.Attach(clock)
	return &DigitalClock{clock}
}
