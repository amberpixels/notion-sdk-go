package notion

// Reference: https://developers.notion.com/reference/block#to-do

// ToDo is a type for to-do blocks
type ToDo struct {
	AtomChildren
	RichText RichTexts `json:"rich_text"`
	Checked  bool      `json:"checked"`
	Color    string    `json:"color,omitempty"`
}

// ToDoBlock is a Notion block for to-do blocks
type ToDoBlock struct {
	BasicBlock
	ToDo ToDo `json:"to_do"`
}

// NewToDoBlock creates a new ToDoBlock
func NewToDoBlock(t ToDo) *ToDoBlock {
	return &ToDoBlock{
		BasicBlock: NewBasicBlock(BlockTypeToDo, t.ChildCount() > 0),
		ToDo:       t,
	}
}

// SetChildren calls inner .SetChildren + updates the HasChildren field
func (b *ToDoBlock) SetChildren(children Blocks) {
	b.ToDo.SetChildren(children)
	b.HasChildren = len(children) > 0
}

// AppendChildren calls inner .AppendChildren + updates the HasChildren field
func (b *ToDoBlock) AppendChildren(children ...Block) {
	b.ToDo.AppendChildren(children...)
	b.HasChildren = b.ToDo.ChildCount() > 0
}

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *ToDoBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

var (
	_ Block             = (*ToDoBlock)(nil)
	_ HierarchicalBlock = (*ToDoBlock)(nil)
	_ BasicBlockHolder  = (*ToDoBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeToDo, func() Block { return &ToDoBlock{} })
}
