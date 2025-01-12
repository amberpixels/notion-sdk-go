package notion

// Reference: https://developers.notion.com/reference/block#quote

// Quote is a type for quote blocks
type Quote struct {
	AtomChildren

	RichText RichTexts `json:"rich_text"`
	Color    string    `json:"color,omitempty"`
}

// QuoteBlock is a Notion block for quote blocks
type QuoteBlock struct {
	BasicBlock
	Quote Quote `json:"quote"`
}

// NewQuoteBlock creates a new QuoteBlock
func NewQuoteBlock(q Quote) *QuoteBlock {
	return &QuoteBlock{
		BasicBlock: NewBasicBlock(BlockTypeQuote, q.ChildCount() > 0),
		Quote:      q,
	}
}

// SetChildren calls inner .SetChildren + updates the HasChildren field
func (b *QuoteBlock) SetChildren(children Blocks) {
	b.Quote.SetChildren(children)
	b.HasChildren = len(children) > 0
}

// AppendChildren calls inner .AppendChildren + updates the HasChildren field
func (b *QuoteBlock) AppendChildren(children ...Block) {
	b.Quote.AppendChildren(children...)
	b.HasChildren = b.Quote.ChildCount() > 0
}

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *QuoteBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

var (
	_ Block             = (*QuoteBlock)(nil)
	_ HierarchicalBlock = (*QuoteBlock)(nil)
	_ BasicBlockHolder  = (*QuoteBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeQuote, func() Block { return &QuoteBlock{} })
}
