package notion

// Reference: https://developers.notion.com/reference/rich-text#equation
//            https://developers.notion.com/reference/block#equation

// Equation holds the equation expression
type Equation struct {
	Expression string `json:"expression"`
}

// TODO? REally?
type EquationBlock struct {
	BasicBlock
	Equation Equation `json:"equation"`
}

// NewEquationRichText creates a new RichText with the given equation expression
func NewEquationRichText(expression string) *RichText {
	return &RichText{
		Type: RichTextTypeEquation,
		Equation: &Equation{
			Expression: expression,
		},
	}
}

func init() {
	registerBlockDecoder(BlockTypeEquation, func() Block { return &EquationBlock{} })
}
