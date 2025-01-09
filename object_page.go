package notion

// PageID stands for ID of Page object.
// As Page is an Object, then PageID is just an alias for Object
type PageID = ObjectID

// The Page object contains the page property values of a single Notion page.
//
// See https://developers.notion.com/reference/page
type Page struct {
	AtomObject
	AtomID
	AtomParent
	AtomCreated
	AtomLastEdited
	AtomArchived
	AtomAppearance
	AtomURLs
	AtomProperties
}

// Pages is a slice of Page objects.
type Pages []*Page

// GetObject always returns ObjectTypePage
func (p *Page) GetObject() ObjectType { return ObjectTypePage }
