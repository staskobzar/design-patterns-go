package interpreter

/*
BooleanExp  ::= VariableExp | Constant | OrExp | AndExp | NotExp |
               '(' BooleanExp ')'
AndExp      ::= BooleanExp 'and' BooleanExp
OrExp       ::= BooleanExp 'or' BooleanExp
NotExp      ::= 'not' BooleanExp
Constant    ::= 'true' | 'false'
VariableExp ::= 'A' | 'B' | ... | 'X' | 'Y' | 'Z'
*/

type BooleanExp interface {
	Evaluate() bool
}

type Context map[string]bool

func NewContext() Context {
	return make(Context)
}

func (c Context) Assign(name string, value bool) {
	c[name] = value
}

func (c Context) Lookup(name string) bool {
	if val, ok := c[name]; ok {
		return val
	}
	return false
}

type Variable struct {
	ctx  Context
	name string
}

func NewVariable(name string, ctx Context) BooleanExp {
	return &Variable{ctx, name}
}

func (v Variable) Evaluate() bool {
	return v.ctx.Lookup(v.name)
}

type Constant bool

func NewConstant(val bool) BooleanExp {
	return Constant(val)
}

func (c Constant) Evaluate() bool {
	return bool(c)
}

type Operator uint8

const (
	OpAND = iota + 100
	OpOR
	OpNOT
)

type Expr struct {
	left  BooleanExp
	right BooleanExp
	op    Operator
}

func NewAndExp(left, right BooleanExp) BooleanExp {
	return &Expr{left, right, OpAND}
}

func NewOrExp(left, right BooleanExp) BooleanExp {
	return &Expr{left, right, OpOR}
}

func NewNotExp(op BooleanExp) BooleanExp {
	return &Expr{op, nil, OpNOT}
}

func (e Expr) Evaluate() bool {
	switch e.op {
	case OpAND:
		return e.left.Evaluate() && e.right.Evaluate()
	case OpOR:
		return e.left.Evaluate() || e.right.Evaluate()
	default:
		return !e.left.Evaluate()
	}
}
