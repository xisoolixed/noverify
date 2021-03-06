package scalar

import (
	"github.com/VKCOM/noverify/src/php/parser/freefloating"
	"github.com/VKCOM/noverify/src/php/parser/node"
	"github.com/VKCOM/noverify/src/php/parser/position"
	"github.com/VKCOM/noverify/src/php/parser/walker"
)

// Encapsed node
type Encapsed struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Parts        []node.Node
}

// NewEncapsed node constructor
func NewEncapsed(Parts []node.Node) *Encapsed {
	return &Encapsed{
		FreeFloating: nil,
		Parts:        Parts,
	}
}

// SetPosition sets node position
func (n *Encapsed) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Encapsed) GetPosition() *position.Position {
	return n.Position
}

func (n *Encapsed) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Encapsed) Walk(v walker.Visitor) {
	if !v.EnterNode(n) {
		return
	}

	if n.Parts != nil {
		for _, nn := range n.Parts {
			if nn != nil {
				nn.Walk(v)
			}
		}
	}

	v.LeaveNode(n)
}
