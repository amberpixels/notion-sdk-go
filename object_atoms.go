package notion

import (
	"time"
)

//
// Atom NoChildren
//

// NoChildren blocks are declared via given atom
type AtomNoChildren struct{}

func (AtomNoChildren) GetChildren() Blocks { return nil }
func (AtomNoChildren) SetChildren(Blocks) {
	panic("it's not possible to SetChildren to childfree objects")
}
func (AtomNoChildren) AppendChildren(...Block) {
	panic("it's not possible to AppendCHildren to childfree objects")
}
func (AtomNoChildren) ChildCount() int { return 0 }

var _ HierarchicalBlock = (*AtomNoChildren)(nil)

//
// Atom Children
//

type AtomChildren struct {
	Children Blocks `json:"children"`
}

func (a AtomChildren) GetChildren() Blocks        { return a.Children }
func (a *AtomChildren) SetChildren(c Blocks)      { a.Children = c }
func (a *AtomChildren) AppendChildren(c ...Block) { a.Children = append(a.Children, c...) }
func (a AtomChildren) ChildCount() int            { return len(a.Children) }

var _ HierarchicalBlock = (*AtomChildren)(nil)

//
// Atom Object
//

type AtomObject struct {
	Object ObjectType `json:"object"`
}

func (b AtomObject) GetObject() ObjectType { return b.Object }

//
// Atom ID
//

// AtomID stands for ID + Parent reference
type AtomID struct {
	ID ObjectID `json:"id,omitempty"`
}

func (a AtomID) GetID() ObjectID { return a.ID }

//
// Atom Parent
//

type AtomParent struct {
	Parent Parent `json:"parent"`
}

func (a AtomParent) GetParent() Parent { return a.Parent }

//
// Atom URLs
//

type AtomURLs struct {
	URL       string `json:"url"`
	PublicURL string `json:"public_url"`
}

func (a AtomURLs) GetURL() string       { return a.URL }
func (a AtomURLs) GetPublicURL() string { return a.PublicURL }

//
// Atom Properties
//

type AtomProperties struct {
	Properties Properties `json:"properties"`
}

func (a AtomProperties) GetProperties() Properties { return a.Properties }

//
// Atom Archived
//

// AtomArchived stands for Archived + InTrash
type AtomArchived struct {
	Archived bool `json:"archived"`
	InTrash  bool `json:"in_trash"`
}

func (a AtomArchived) GetArchived() bool { return a.Archived }
func (a AtomArchived) GetInTrash() bool  { return a.InTrash }

//
// Atom Created (time+by)
//

// AtomCreated stands for CreatedTime + CreatedBy fields
type AtomCreated struct {
	CreatedTime *time.Time `json:"created_time,omitempty"`
	CreatedBy   *User      `json:"created_by,omitempty"`
}

func (a AtomCreated) GetCreatedTime() *time.Time { return a.CreatedTime }

func (a AtomCreated) GetCreatedBy() *User { return a.CreatedBy }

//
// Atom LastEdited (time+by)
//

type AtomLastEdited struct {
	LastEditedTime *time.Time `json:"last_edited_time"`
	LastEditedBy   *User      `json:"last_edited_by"`
}

func (a AtomLastEdited) GetLastEditedTime() *time.Time { return a.LastEditedTime }

func (a AtomLastEdited) GetLastEditedBy() *User { return a.LastEditedBy }
