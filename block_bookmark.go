package notion

// Reference: https://developers.notion.com/reference/block#bookmark

// Bookmark is a block that represents a bookmark (URL + Caption).
type Bookmark struct {
	Caption RichTexts `json:"caption,omitempty"`
	URL     string    `json:"url"`
}

// BookmarkBlock is a Notion block for a bookmark.
type BookmarkBlock struct {
	BasicBlock
	Bookmark Bookmark `json:"bookmark"`
}

// NewBookmarkBlock creates a new BookmarkBlock.
func NewBookmarkBlock(b Bookmark) *BookmarkBlock {
	return &BookmarkBlock{
		BasicBlock: NewBasicBlock(BlockTypeBookmark),
		Bookmark:   b,
	}
}

var (
	_ Block             = (*BookmarkBlock)(nil)
	_ HierarchicalBlock = (*BookmarkBlock)(nil)
)

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *BookmarkBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

func init() {
	registerBlockDecoder(BlockTypeBookmark, func() Block { return &BookmarkBlock{} })
}
