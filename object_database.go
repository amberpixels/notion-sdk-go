package notion

// DatabaseID stands for ID of Database object.
// As Database is an Object, then DatabaseID is just an alias for Object
type DatabaseID = ObjectID

// Database is a Notion object that represents a database.
type Database struct {
	AtomObject
	AtomID
	AtomParent
	AtomCreated
	AtomLastEdited
	AtomArchived
	AtomAppearance
	AtomURLs

	// Title of the page is represented via RichTexts(type==text)
	Title RichTexts `json:"title"`
	// Description of the page is represented via RichTexts(type==text)
	Description RichTexts `json:"description"`

	// Note on both Title & Description:
	// We do not have a separate limited RichText(type==text) type,
	// So we consider users to carefuly use it.
	// TODO: We can make it safer, at least via validation on Update/Create endpoints

	IsInline bool `json:"is_inline"`

	// Properties is a map of property configurations that defines what Page.Properties each page of the database can use
	Properties PropertyConfigs `json:"properties"`
}

var _ Object = (*Database)(nil)

// GetObject always returns ObjectTypeDatabase
func (db *Database) GetObject() ObjectType { return ObjectTypeDatabase }

// RelationObject TODO: should be part of PropertyConfigs, right?
type RelationObject struct {
	Database           DatabaseID `json:"database"`
	SyncedPropertyName string     `json:"synced_property_name"`
}
