package notion

import (
	"encoding/json"
	"fmt"
	"time"
)

// BlockID stands for ID of Block object.
// As Block is an Object, then BlockID is just an alias for Object
type BlockID = ObjectID

// BlockType is a type of a Notion block.
type BlockType string

func (bt BlockType) String() string { return string(bt) }

// See https://developers.notion.com/reference/block
const (
	BlockTypeParagraph BlockType = "paragraph"
	BlockTypeHeading1  BlockType = "heading_1"
	BlockTypeHeading2  BlockType = "heading_2"
	BlockTypeHeading3  BlockType = "heading_3"

	BlockTypeBulletedListItem BlockType = "bulleted_list_item"
	BlockTypeNumberedListItem BlockType = "numbered_list_item"

	BlockTypeToDo          BlockType = "to_do"
	BlockTypeToggle        BlockType = "toggle"
	BlockTypeChildPage     BlockType = "child_page"
	BlockTypeChildDatabase BlockType = "child_database"

	BlockTypeEmbed           BlockType = "embed"
	BlockTypeImage           BlockType = "image"
	BlockTypeAudio           BlockType = "audio"
	BlockTypeVideo           BlockType = "video"
	BlockTypeFile            BlockType = "file"
	BlockTypePdf             BlockType = "pdf"
	BlockTypeBookmark        BlockType = "bookmark"
	BlockTypeCode            BlockType = "code"
	BlockTypeDivider         BlockType = "divider"
	BlockTypeCallout         BlockType = "callout"
	BlockTypeQuote           BlockType = "quote"
	BlockTypeTableOfContents BlockType = "table_of_contents"
	BlockTypeEquation        BlockType = "equation"
	BlockTypeBreadcrumb      BlockType = "breadcrumb"
	BlockTypeColumn          BlockType = "column"
	BlockTypeColumnList      BlockType = "column_list"
	BlockTypeLinkPreview     BlockType = "link_preview"
	BlockTypeLinkToPage      BlockType = "link_to_page"
	BlockTypeSyncedBlock     BlockType = "synced_block"
	BlockTypeTable           BlockType = "table"
	BlockTypeTableRow        BlockType = "table_row"
	BlockTypeUnsupported     BlockType = "unsupported"
)

const (
	// Deprecated
	// See https://developers.notion.com/reference/block#template
	BlockTypeTemplate BlockType = "template"
)

// Block is a general interface for ALL types of notion Blocks.
type Block interface {
	Object // Every block is an Object by default

	GetID() BlockID
	GetParent() Parent
	GetType() BlockType

	GetCreatedTime() *time.Time
	GetCreatedBy() *User

	GetLastEditedTime() *time.Time
	GetLastEditedBy() *User

	GetArchived() bool
	GetInTrash() bool

	GetHasChildren() bool
}

// HierarchicalBlock is a block that can handle its GetChildren
// Even childfree blocks should implement it (see AtomChildfree)
type HierarchicalBlock interface {
	GetChildren() Blocks
	SetChildren(Blocks)
	AppendChildren(...Block)
	ChildCount() int
}

// BasicBlock is a block that can handle its GetBasicBlock
// It's not supposed to be widely used, but it's useful for testing,
// and other tasks where you need reflect-type Block access
type BasicBlockHolder interface {
	GetBasicBlock() BasicBlock
	SetBasicBlock(BasicBlock) Block
}

// Blocks is a slice of (generic) Block objects.
type Blocks []Block

// UnmarshalJSON implements custom unmarshalling for Blocks
// It's required because each type of Block will have its own type-based field.
// see decodeBlock() for more details
func (b *Blocks) UnmarshalJSON(data []byte) error {
	var err error
	mapArr := make([]map[string]any, 0)
	if err = json.Unmarshal(data, &mapArr); err != nil {
		return err
	}

	result := make([]Block, len(mapArr))
	for i, prop := range mapArr {
		if result[i], err = decodeBlock(prop); err != nil {
			return err
		}
	}

	*b = result
	return nil
}

// BasicBlock defines the common fields of all Notion block types.
// See https://developers.notion.com/reference/block for the list.
// BasicBlock implements the Block interface.
type BasicBlock struct {
	AtomID
	AtomParent
	AtomObject
	AtomCreated
	AtomLastEdited
	AtomArchived
	AtomAppearance

	// blocks by default have NO children unless it's specified on the custom block implementation
	AtomNoChildren

	Type        BlockType `json:"type"`
	HasChildren bool      `json:"has_children,omitempty"`
}

// NewBasicBlock returns a new BasicBlock with the ObjectTypeBlock and given block type.
// It's used as a basic block for all other blocks.
func NewBasicBlock(blockType BlockType, hasChildrenArg ...bool) BasicBlock {
	var hasChildren bool
	if len(hasChildrenArg) > 0 {
		hasChildren = hasChildrenArg[0]
	}

	return BasicBlock{
		AtomObject: AtomObject{
			Object: ObjectTypeBlock,
		},
		Type:        blockType,
		HasChildren: hasChildren,
	}
}

func (b BasicBlock) GetType() BlockType    { return b.Type }
func (b BasicBlock) GetObject() ObjectType { return ObjectTypeBlock }
func (b BasicBlock) GetHasChildren() bool  { return b.HasChildren }

var _ Block = (*BasicBlock)(nil)
var _ HierarchicalBlock = (*BasicBlock)(nil)

var _ Block = (*BasicBlock)(nil)

// Note: we do not care about the order of registration and obout concurrency
// TODO do we?
var blockConstructors = map[BlockType]func() Block{}

func registerBlockDecoder(blockType BlockType, blockConstructor func() Block) {
	blockConstructors[blockType] = blockConstructor
}

func decodeBlock(raw map[string]any) (Block, error) {
	blockType, ok := raw["type"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid block type")
	}

	constructor, found := blockConstructors[BlockType(blockType)]
	if !found {
		constructor = func() Block { return &UnsupportedBlock{} } // Default to UnsupportedBlock
	}

	// Create the block
	block := constructor()

	j, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(j, block); err != nil {
		return nil, err
	}

	return block, nil
}
