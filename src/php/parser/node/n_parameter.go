package node

import (
	"github.com/VKCOM/noverify/src/php/parser/freefloating"
	"github.com/VKCOM/noverify/src/php/parser/position"
	"github.com/VKCOM/noverify/src/php/parser/walker"
)

// Parameter node
type Parameter struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	ByRef        bool
	Variadic     bool
	VariableType Node
	Variable     *SimpleVar
	DefaultValue Node
}

// NewParameter node constructor
func NewParameter(VariableType Node, variable *SimpleVar, DefaultValue Node, ByRef bool, Variadic bool) *Parameter {
	return &Parameter{
		FreeFloating: nil,
		ByRef:        ByRef,
		Variadic:     Variadic,
		VariableType: VariableType,
		Variable:     variable,
		DefaultValue: DefaultValue,
	}
}

// SetPosition sets node position
func (n *Parameter) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Parameter) GetPosition() *position.Position {
	return n.Position
}

func (n *Parameter) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Parameter) Walk(v walker.Visitor) {
	if !v.EnterNode(n) {
		return
	}

	if n.VariableType != nil {
		n.VariableType.Walk(v)
	}

	if n.Variable != nil {
		n.Variable.Walk(v)
	}

	if n.DefaultValue != nil {
		n.DefaultValue.Walk(v)
	}

	v.LeaveNode(n)
}
