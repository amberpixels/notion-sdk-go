package notion

// Reference: https://developers.notion.com/reference/block#code

// Code stores the code, caption, and language of the code block
type Code struct {
	RichText RichTexts `json:"rich_text"`
	Caption  RichTexts `json:"caption,omitempty"`
	Language string    `json:"language"`
}

// CodeBlock is a Notion block for Code
type CodeBlock struct {
	BasicBlock
	Code Code `json:"code"`
}

// NewCodeBlock returns a new CodeBlock with the given code
func NewCodeBlock(code Code) *CodeBlock {
	return &CodeBlock{
		BasicBlock: NewBasicBlock(BlockTypeCode),
		Code:       code,
	}
}

var (
	_ Block             = (*CodeBlock)(nil)
	_ HierarchicalBlock = (*CodeBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeCode, func() Block { return &CodeBlock{} })
}
