package notion

// Reference: https://developers.notion.com/reference/block#equation

// EquationBlock is a Notion block for Standalone equation blocks
// It differs from inline Equation: Inline equations are always part of paragraph (or other parent block)
// Standalone equations are always standalone blocks
type EquationBlock struct {
	BasicBlock
	Equation Equation `json:"equation"`
}

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *EquationBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

// NewEquationBlock creates a new Equation block with the given equation expression
func NewEquationBlock(expression string) *RichText {
	return &RichText{
		Type: RichTextTypeEquation,
		Equation: &Equation{
			Expression: expression,
		},
	}
}

var (
	_ Block             = (*EquationBlock)(nil)
	_ HierarchicalBlock = (*EquationBlock)(nil)
	_ BasicBlockHolder  = (*EquationBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeEquation, func() Block { return &EquationBlock{} })
}
