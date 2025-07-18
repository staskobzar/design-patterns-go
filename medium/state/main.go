package main

import "fmt"

type PhoneState interface {
	onReceiveCall(context *Phone)
}

// concrete states
type NormalState struct{}

func (ns *NormalState) onReceiveCall(context *Phone) {
	fmt.Println("[S] Incoming call: Ring ring! (Normal mode)")
}

type VibrateState struct {
	vibrations int
}

func (vs *VibrateState) onReceiveCall(context *Phone) {
	vs.vibrations++
	fmt.Println("[S] Incoming call: Bzzt bzzt... (Vibrate mode)")

	if vs.vibrations >= 3 {
		context.setState(&SilentState{})
		fmt.Printf("[!] No answer after %d vibrations, switch to Silent mode.\n", vs.vibrations)
	}
}

type SilentState struct{}

func (ss *SilentState) onReceiveCall(context *Phone) {
	fmt.Println("[S] Incoming call: (Silent mode, no sound)")
	fmt.Println("[!] The phone stays silent. You might see a missed call later.")
}
func (SilentState) String() string { return "SilentState" }

// context
type Phone struct {
	state PhoneState
}

func (p *Phone) receiveCall() {
	p.state.onReceiveCall(p)
}

func (p *Phone) setState(state PhoneState) {
	p.state = state
}

func (p *Phone) modeName() string { return fmt.Sprintf("%v", p.state) }

func main() {
	fmt.Println("==== PHONE ====")
	fmt.Println("===============")

	phone := &Phone{}
	phone.setState(&NormalState{})

	fmt.Printf("Phone is now in %s.\n", phone.modeName())
	phone.receiveCall()

	fmt.Println()
	phone.setState(&VibrateState{})
	fmt.Printf("Phone is now in %s.\n", phone.modeName())
	phone.receiveCall()
	phone.receiveCall()
	phone.receiveCall()

	fmt.Println()
	fmt.Printf("Phone is now in %s.\n", phone.modeName())
	phone.receiveCall()

	fmt.Println()
	phone.setState(&NormalState{})
	fmt.Printf("Phone is now in %s.\n", phone.modeName())
	phone.receiveCall()
}
