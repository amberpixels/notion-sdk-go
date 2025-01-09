package notion

// Ref: https://developers.notion.com/reference/block#synced-block

type Synced struct {
	AtomChildren

	// SyncedFrom is nil for the original block.
	SyncedFrom *SyncedFrom `json:"synced_from"`
}

type SyncedFrom struct {
	BlockID BlockID `json:"block_id"`
}

type SyncedBlock struct {
	BaseBlock
	Synced Synced `json:"synced_block"`
}

func NewSyncedBlock(s Synced) *SyncedBlock {
	return &SyncedBlock{
		BaseBlock: NewBaseBlock(BlockTypeSyncedBlock, s.ChildCount() > 0),
		Synced:    s,
	}
}

// SetChildren calls inner .SetChildren + updates the HasChildren field
func (b *SyncedBlock) SetChildren(children Blocks) {
	b.Synced.SetChildren(children)
	b.HasChildren = len(children) > 0
}

// AppendChildren calls inner .AppendChildren + updates the HasChildren field
func (b *SyncedBlock) AppendChildren(children ...Block) {
	b.Synced.AppendChildren(children...)
	b.HasChildren = b.Synced.ChildCount() > 0
}

var (
	_ Block             = (*SyncedBlock)(nil)
	_ HierarchicalBlock = (*SyncedBlock)(nil)
)
