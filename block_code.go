package notion

type Code struct {
	RichText RichTexts `json:"rich_text"`
	Caption  RichTexts `json:"caption,omitempty"`
	Language string    `json:"language"`
}

type CodeBlock struct {
	BaseBlock
	Code Code `json:"code"`
}

func NewCodeBlock(code Code) *CodeBlock {
	return &CodeBlock{
		BaseBlock: NewBaseBlock(BlockTypeCode),
		Code:      code,
	}
}

var (
	_ Block             = (*CodeBlock)(nil)
	_ HierarchicalBlock = (*CodeBlock)(nil)
)
