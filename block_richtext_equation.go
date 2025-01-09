package notion

// See: https://developers.notion.com/reference/block#equation
type Equation struct {
	Expression string `json:"expression"`
}

type EquationBlock struct {
	BaseBlock
	Equation Equation `json:"equation"`
}

func NewEquationBlock(eq Equation) *EquationBlock {
	return &EquationBlock{
		BaseBlock: NewBaseBlock(BlockTypeEquation),
		Equation:  eq,
	}
}
