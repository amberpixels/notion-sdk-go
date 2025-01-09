package notion

// Ref: https://developers.notion.com/reference/block#to-do

type ToDo struct {
	AtomChildren
	RichText RichTexts `json:"rich_text"`
	Checked  bool      `json:"checked"`
	Color    string    `json:"color,omitempty"`
}

type ToDoBlock struct {
	BaseBlock
	ToDo ToDo `json:"to_do"`
}

func NewToDoBlock(t ToDo) *ToDoBlock {
	return &ToDoBlock{
		BaseBlock: NewBaseBlock(BlockTypeToDo, t.ChildCount() > 0),
		ToDo:      t,
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

var (
	_ Block             = (*ToDoBlock)(nil)
	_ HierarchicalBlock = (*ToDoBlock)(nil)
)
