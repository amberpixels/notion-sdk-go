package notion

// Reference: https://developers.notion.com/reference/block#embed

// Embed stores the caption and URL of the embed block
type Embed struct {
	Caption RichTexts `json:"caption,omitempty"`
	URL     string    `json:"url"`
}

// EmbedBlock is a Notion block for Embed
type EmbedBlock struct {
	BasicBlock
	Embed Embed `json:"embed"`
}

// NewEmbedBlock returns a new EmbedBlock with the given embed
func NewEmbedBlock(embed Embed) *EmbedBlock {
	return &EmbedBlock{
		BasicBlock: NewBasicBlock(BlockTypeEmbed),
		Embed:      embed,
	}
}

var (
	_ Block             = (*EmbedBlock)(nil)
	_ HierarchicalBlock = (*EmbedBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeEmbed, func() Block { return &EmbedBlock{} })
}
