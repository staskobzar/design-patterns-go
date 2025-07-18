package state

type TCPState interface {
	ActiveOpen(*TCPConnection)
	PassiveOpen(*TCPConnection)
	Close(*TCPConnection)
	ChangeState(*TCPConnection, TCPState)
}

type State struct{}

func (s *State) ChangeState(c *TCPConnection, state TCPState) {
	c.ChangeState(state)
}
func (s *State) ActiveOpen(c *TCPConnection)  {}
func (s *State) PassiveOpen(c *TCPConnection) {}
func (s *State) Close(c *TCPConnection) {
	s.ChangeState(c, NewTCPClosed())
}

// TCPConnection --------------------------------------------------
type TCPConnection struct {
	state TCPState
}

func NewTCPConnection() *TCPConnection {
	return &TCPConnection{state: NewTCPClosed()}
}

func (c *TCPConnection) ActiveOpen() {
	c.state.ActiveOpen(c)
	// network routin
}

func (c *TCPConnection) PassiveOpen() {
	c.state.PassiveOpen(c)
}

func (c *TCPConnection) Close() {
	c.state.Close(c)
}

func (c *TCPConnection) ChangeState(state TCPState) {
	c.state = state
}

// TCPClosed --------------------------------------------------
type TCPClosed struct{ State }

func NewTCPClosed() TCPState {
	return &TCPClosed{}
}

func (s *TCPClosed) ActiveOpen(c *TCPConnection) {
	s.ChangeState(c, NewTCPEstablished())
}

func (s *TCPClosed) PassiveOpen(c *TCPConnection) {
	s.ChangeState(c, NewTCPListen())
}

// TCPEstablised --------------------------------------------------
type TCPEstablised struct{ State }

func NewTCPEstablished() TCPState {
	return &TCPEstablised{}
}

// TCPListen --------------------------------------------------
type TCPListen struct{ State }

func NewTCPListen() TCPState {
	return &TCPListen{}
}
