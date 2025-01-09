package notion

// ObjectID is a unique identifier for a Notion object.
// It is set on the Notion side only (immutable and read-only).
// It will be ignored if you manually set its value for the new objects.
type ObjectID string

// String returns the string representation of the ObjectID.
func (oID ObjectID) String() string { return string(oID) }

// ObjectType is a type of a Notion object.
type ObjectType string

// String returns the string representation of the ObjectType.
func (ot ObjectType) String() string { return string(ot) }

const (
	ObjectTypeDatabase ObjectType = "database"
	ObjectTypeBlock    ObjectType = "block"
	ObjectTypePage     ObjectType = "page"
	ObjectTypeList     ObjectType = "list"
	ObjectTypeText     ObjectType = "text"
	ObjectTypeUser     ObjectType = "user"
	ObjectTypeError    ObjectType = "error"
	ObjectTypeComment  ObjectType = "comment"
)

// Object is an interface for all Notion objects.
// ObjectType ("object") field is the only shared field between all objects.
type Object interface {
	GetObject() ObjectType
}

// Objects is a slice of Object.
type Objects []Object
