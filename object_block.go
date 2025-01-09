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

// BaseBlock defines the common fields of all Notion block types.
// See https://developers.notion.com/reference/block for the list.
// BaseBlock implements the Block interface.
type BaseBlock struct {
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

// NewBaseBlock returns a new BaseBlock with the ObjectTypeBlock and given block type.
// It's used as a basic block for all other blocks.
func NewBaseBlock(blockType BlockType, hasChildrenArg ...bool) BaseBlock {
	var hasChildren bool
	if len(hasChildrenArg) > 0 {
		hasChildren = hasChildrenArg[0]
	}

	return BaseBlock{
		AtomObject: AtomObject{
			Object: ObjectTypeBlock,
		},
		Type:        blockType,
		HasChildren: hasChildren,
	}
}

func (b BaseBlock) GetType() BlockType    { return b.Type }
func (b BaseBlock) GetObject() ObjectType { return ObjectTypeBlock }
func (b BaseBlock) GetHasChildren() bool  { return b.HasChildren }

var _ Block = (*BaseBlock)(nil)
var _ HierarchicalBlock = (*BaseBlock)(nil)

var _ Block = (*BaseBlock)(nil)

func decodeBlock(raw map[string]any) (Block, error) {
	blockConstructors := map[BlockType]func() Block{
		BlockTypeParagraph:        func() Block { return &ParagraphBlock{} },
		BlockTypeHeading1:         func() Block { return &Heading1Block{} },
		BlockTypeHeading2:         func() Block { return &Heading2Block{} },
		BlockTypeHeading3:         func() Block { return &Heading3Block{} },
		BlockTypeCallout:          func() Block { return &CalloutBlock{} },
		BlockTypeQuote:            func() Block { return &QuoteBlock{} },
		BlockTypeBulletedListItem: func() Block { return &BulletedListItemBlock{} },
		BlockTypeNumberedListItem: func() Block { return &NumberedListItemBlock{} },
		BlockTypeToDo:             func() Block { return &ToDoBlock{} },
		BlockTypeCode:             func() Block { return &CodeBlock{} },
		BlockTypeToggle:           func() Block { return &ToggleBlock{} },
		BlockTypeChildPage:        func() Block { return &ChildPageBlock{} },
		BlockTypeEmbed:            func() Block { return &EmbedBlock{} },
		BlockTypeImage:            func() Block { return &ImageBlock{} },
		BlockTypeAudio:            func() Block { return &AudioBlock{} },
		BlockTypeVideo:            func() Block { return &VideoBlock{} },
		BlockTypeFile:             func() Block { return &FileBlock{} },
		BlockTypePdf:              func() Block { return &PdfBlock{} },
		BlockTypeBookmark:         func() Block { return &BookmarkBlock{} },
		BlockTypeChildDatabase:    func() Block { return &ChildDatabaseBlock{} },
		BlockTypeTableOfContents:  func() Block { return &TableOfContentsBlock{} },
		BlockTypeDivider:          func() Block { return &DividerBlock{} },
		BlockTypeEquation:         func() Block { return &EquationBlock{} },
		BlockTypeBreadcrumb:       func() Block { return &BreadcrumbBlock{} },
		BlockTypeColumn:           func() Block { return &ColumnBlock{} },
		BlockTypeColumnList:       func() Block { return &ColumnListBlock{} },
		BlockTypeLinkPreview:      func() Block { return &LinkPreviewBlock{} },
		BlockTypeLinkToPage:       func() Block { return &LinkToPageBlock{} },
		BlockTypeTemplate:         func() Block { return &TemplateBlock{} },
		BlockTypeSyncedBlock:      func() Block { return &SyncedBlock{} },
		BlockTypeTable:            func() Block { return &TableBlock{} },
		BlockTypeTableRow:         func() Block { return &TableRowBlock{} },
		BlockTypeUnsupported:      func() Block { return &UnsupportedBlock{} },
	}

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
