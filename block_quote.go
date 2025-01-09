package notion

type Quote struct {
	AtomChildren

	RichText RichTexts `json:"rich_text"`
	Color    string    `json:"color,omitempty"`
}

type QuoteBlock struct {
	BaseBlock
	Quote Quote `json:"quote"`
}

func NewQuoteBlock(q Quote) *QuoteBlock {
	return &QuoteBlock{
		BaseBlock: NewBaseBlock(BlockTypeQuote, q.ChildCount() > 0),
		Quote:     q,
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

var (
	_ Block             = (*QuoteBlock)(nil)
	_ HierarchicalBlock = (*QuoteBlock)(nil)
)
