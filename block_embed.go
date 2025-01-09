package notion

type Embed struct {
	Caption RichTexts `json:"caption,omitempty"`
	URL     string    `json:"url"`
}

type EmbedBlock struct {
	BaseBlock
	Embed Embed `json:"embed"`
}

func NewEmbedBlock(embed Embed) *EmbedBlock {
	return &EmbedBlock{
		BaseBlock: NewBaseBlock(BlockTypeEmbed),
		Embed:     embed,
	}
}

var (
	_ Block             = (*EmbedBlock)(nil)
	_ HierarchicalBlock = (*EmbedBlock)(nil)
)
