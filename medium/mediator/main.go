package main

// FrontDoorSensor: Motion detected!
// Hub: FrontDoorSensor reported motion. Notifying relevant devices...
// LivingRoomLight: Turning ON
// MainAlarm: Ringing alarm!

import (
	"fmt"
)

type Event string

const (
	Motion Event = "MOTION"
	Disarm       = "DISARM"
)

type SmartHomeMediator interface {
	notify(sender Device, event Event)
}

type Device interface {
	init(string, SmartHomeMediator)
	activate()
	desactivate()
	Name() string
}

type DeviceKernel struct {
	mediator SmartHomeMediator
	name     string
}

func (d *DeviceKernel) init(name string, mediator SmartHomeMediator) {
	d.mediator = mediator
	d.name = name
}
func (d *DeviceKernel) Name() string { return d.name }
func (d *DeviceKernel) activate()    {}
func (d *DeviceKernel) desactivate() {}

type MotionDetector struct {
	DeviceKernel
}

func (m *MotionDetector) motionDetected() {
	m.mediator.notify(m, Motion)
}

func (m *MotionDetector) disarm() {
	m.mediator.notify(m, Disarm)
}

type Alarm struct {
	DeviceKernel
}

func (a *Alarm) activate() {
	fmt.Printf("%q: Ringing alarm!\n", a.name)
}

func (a *Alarm) desactivate() {
	fmt.Printf("%q: Ringing stop!\n", a.name)
}

type Light struct {
	DeviceKernel
}

func (l *Light) activate() {
	fmt.Printf("%q: Turn ON!\n", l.name)
}

func (l *Light) desactivate() {
	fmt.Printf("%q: Turn OFF!\n", l.name)
}

type SmartHomeHub struct {
	devices []Device
}

func (h *SmartHomeHub) addDevice(d Device) {
	h.devices = append(h.devices, d)
}

func (h *SmartHomeHub) notify(d Device, e Event) {
	switch e {
	case Motion:
		fmt.Printf("Hub: %q reported motion. Notifying relevant devices...\n", d.Name())
		h.motionDetected()
	case Disarm:
		fmt.Printf("Hub: %q disarmed. Notifying device...\n", d.Name())
		h.disarm()
	default:
		fmt.Printf("invalid event: %#v\n", e)
	}
}

func (h *SmartHomeHub) motionDetected() {
	for _, d := range h.devices {
		switch d.(type) {
		case *Alarm, *Light:
			d.activate()
		}
	}
}

func (h *SmartHomeHub) disarm() {
	for _, d := range h.devices {
		switch d.(type) {
		case *Alarm, *Light:
			d.desactivate()
		}
	}
}

func main() {
	fmt.Println("==================================================")
	fmt.Println("================================================== MEDIATOR")

	mediator := &SmartHomeHub{}

	detector := &MotionDetector{}
	detector.init("motion detector", mediator)

	alarm := &Alarm{}
	alarm.init("alarm", mediator)

	light := &Light{}
	light.init("light", mediator)

	mediator.addDevice(detector)
	mediator.addDevice(alarm)
	mediator.addDevice(light)

	detector.motionDetected()
	fmt.Println("all checked, all good. disarm alarm")
	detector.disarm()
}
