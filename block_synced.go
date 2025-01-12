package notion

// Reference: https://developers.notion.com/reference/block#synced-block

// Synced is a type for synced blocks
type Synced struct {
	AtomChildren

	// SyncedFrom is nil for the original block.
	SyncedFrom *SyncedFrom `json:"synced_from"`
}

// SyncedFrom holds the ID of the original block
type SyncedFrom struct {
	BlockID BlockID `json:"block_id"`
}

// SyncedBlock is a Notion block for synced blocks
type SyncedBlock struct {
	BasicBlock
	Synced Synced `json:"synced_block"`
}

// NewSyncedBlock creates a new SyncedBlock
func NewSyncedBlock(s Synced) *SyncedBlock {
	return &SyncedBlock{
		BasicBlock: NewBasicBlock(BlockTypeSyncedBlock, s.ChildCount() > 0),
		Synced:     s,
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

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *SyncedBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

var (
	_ Block             = (*SyncedBlock)(nil)
	_ HierarchicalBlock = (*SyncedBlock)(nil)
	_ BasicBlockHolder  = (*SyncedBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeSyncedBlock, func() Block { return &SyncedBlock{} })
}
