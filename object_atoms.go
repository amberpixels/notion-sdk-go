package notion

import (
	"time"
)

//
// Atom NoChildren
//

// AtomNoChildren should be embedded in all blocks that do not have children.
type AtomNoChildren struct{}

// GetChildren returns nil for the AtomNoChildren.
func (AtomNoChildren) GetChildren() Blocks { return nil }

// SetChildren panics for the AtomNoChildren.
func (AtomNoChildren) SetChildren(Blocks) {
	panic("it's not possible to SetChildren to childfree objects")
}

// AppendChildren panics for the AtomNoChildren.
func (AtomNoChildren) AppendChildren(...Block) {
	panic("it's not possible to AppendCHildren to childfree objects")
}

// ChildCount returns 0 for the AtomNoChildren.
func (AtomNoChildren) ChildCount() int { return 0 }

var _ HierarchicalBlock = (*AtomNoChildren)(nil)

//
// Atom Children
//

// AtomChildren stands for atom that holds children.
type AtomChildren struct {
	Children Blocks `json:"children"`
}

// GetChildren returns the children of the AtomChildren.
func (a AtomChildren) GetChildren() Blocks { return a.Children }

// SetChildren sets the children of the AtomChildren.
func (a *AtomChildren) SetChildren(c Blocks) { a.Children = c }

// AppendChildren appends the children of the AtomChildren.
func (a *AtomChildren) AppendChildren(c ...Block) { a.Children = append(a.Children, c...) }

// ChildCount returns the number
func (a AtomChildren) ChildCount() int { return len(a.Children) }

var _ HierarchicalBlock = (*AtomChildren)(nil)

//
// Atom Object
//

// AtomObject stands for atom that holds object (ObjctType).
type AtomObject struct {
	Object ObjectType `json:"object"`
}

// GetObject returns the object (ObjectType) of the AtomObject.
func (b AtomObject) GetObject() ObjectType { return b.Object }

//
// Atom ID
//

// AtomID stands for ID + Parent reference
type AtomID struct {
	ID ObjectID `json:"id,omitempty"`
}

// GetID returns the ID of the AtomID.
func (a AtomID) GetID() ObjectID { return a.ID }

//
// Atom Parent
//

// AtomParent stands for atom that holds Notion's parent object.
type AtomParent struct {
	Parent Parent `json:"parent"`
}

// GetParent returns the parent of the AtomParent.
func (a AtomParent) GetParent() Parent { return a.Parent }

//
// Atom URLs
//

// AtomURLs stands for atom that holds Notion's URLs.
type AtomURLs struct {
	URL       string `json:"url"`
	PublicURL string `json:"public_url"`
}

// GetURL returns the URL of the AtomURLs.
func (a AtomURLs) GetURL() string { return a.URL }

// GetPublicURL returns the PublicURL of the AtomURLs.
func (a AtomURLs) GetPublicURL() string { return a.PublicURL }

//
// Atom Properties
//

// AtomProperties stands for atom that holds Notion's properties.
type AtomProperties struct {
	Properties Properties `json:"properties"`
}

// GetProperties returns the properties of the AtomProperties.
func (a AtomProperties) GetProperties() Properties { return a.Properties }

//
// Atom Archived
//

// AtomArchived stands for Archived + InTrash
type AtomArchived struct {
	Archived bool `json:"archived"`
	InTrash  bool `json:"in_trash"`
}

// GetArchived returns the Archived bool of the AtomArchived.
func (a AtomArchived) GetArchived() bool { return a.Archived }

// GetInTrash returns the InTrash bool of the AtomArchived.
func (a AtomArchived) GetInTrash() bool { return a.InTrash }

//
// Atom Created (time+by)
//

// AtomCreated stands for CreatedTime + CreatedBy fields
type AtomCreated struct {
	CreatedTime *time.Time `json:"created_time,omitempty"`
	CreatedBy   *User      `json:"created_by,omitempty"`
}

// GetCreatedTime returns the CreatedTime of the AtomCreated.
func (a AtomCreated) GetCreatedTime() *time.Time { return a.CreatedTime }

// GetCreatedBy returns the CreatedBy of the AtomCreated.
func (a AtomCreated) GetCreatedBy() *User { return a.CreatedBy }

//
// Atom LastEdited (time+by)
//

// AtomLastEdited stands for atom that holds Notion's last edited time and by.
type AtomLastEdited struct {
	LastEditedTime *time.Time `json:"last_edited_time"`
	LastEditedBy   *User      `json:"last_edited_by"`
}

// GetLastEditedTime returns the LastEditedTime of the AtomLastEdited.
func (a AtomLastEdited) GetLastEditedTime() *time.Time { return a.LastEditedTime }

// GetLastEditedBy returns the LastEditedBy of the AtomLastEdited.
func (a AtomLastEdited) GetLastEditedBy() *User { return a.LastEditedBy }
