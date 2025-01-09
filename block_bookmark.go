package notion

type Bookmark struct {
	Caption RichTexts `json:"caption,omitempty"`
	URL     string    `json:"url"`
}

type BookmarkBlock struct {
	BaseBlock
	Bookmark Bookmark `json:"bookmark"`
}

func NewBookmarkBlock(bookmark Bookmark) *BookmarkBlock {
	return &BookmarkBlock{
		BaseBlock: NewBaseBlock(BlockTypeBookmark),
		Bookmark:  bookmark,
	}
}

var (
	_ Block             = (*BookmarkBlock)(nil)
	_ HierarchicalBlock = (*BookmarkBlock)(nil)
)
