package notion

// Reference: https://developers.notion.com/reference/rich-text#equation

// Equation holds the equation expression
// It's used both for inline equations and standalone equations.
type Equation struct {
	Expression string `json:"expression"`
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
