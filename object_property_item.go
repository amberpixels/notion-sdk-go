package notion

import (
	"encoding/json"
	"fmt"
	"time"
)

// PropertyType is a type for Notion property types.
type PropertyType string

// nolint:revive
const (
	PropertyTypeTitle          PropertyType = "title"
	PropertyTypeRichText       PropertyType = "rich_text"
	PropertyTypeText           PropertyType = "text"
	PropertyTypeNumber         PropertyType = "number"
	PropertyTypeSelect         PropertyType = "select"
	PropertyTypeMultiSelect    PropertyType = "multi_select"
	PropertyTypeDate           PropertyType = "date"
	PropertyTypeFormula        PropertyType = "formula"
	PropertyTypeRelation       PropertyType = "relation"
	PropertyTypeRollup         PropertyType = "rollup"
	PropertyTypePeople         PropertyType = "people"
	PropertyTypeFiles          PropertyType = "files"
	PropertyTypeCheckbox       PropertyType = "checkbox"
	PropertyTypeURL            PropertyType = "url"
	PropertyTypeEmail          PropertyType = "email"
	PropertyTypePhoneNumber    PropertyType = "phone_number"
	PropertyTypeCreatedTime    PropertyType = "created_time"
	PropertyTypeCreatedBy      PropertyType = "created_by"
	PropertyTypeLastEditedTime PropertyType = "last_edited_time"
	PropertyTypeLastEditedBy   PropertyType = "last_edited_by"
	PropertyTypeStatus         PropertyType = "status"
	PropertyTypeUniqueID       PropertyType = "unique_id"
	PropertyTypeVerification   PropertyType = "verification"
	PropertyTypeButton         PropertyType = "button"
)

// Property is an interface for Notion properties.
type Property interface {
	GetID() string
	GetType() PropertyType
}

// PropertyArray is a slice of Notion properties.
type PropertyArray []Property

// UnmarshalJSON implements custom unmarshalling for PropertyArray
func (arr *PropertyArray) UnmarshalJSON(data []byte) error {
	var err error
	mapArr := make([]map[string]interface{}, 0)
	if err = json.Unmarshal(data, &mapArr); err != nil {
		return err
	}

	result := make([]Property, len(mapArr))
	for i, prop := range mapArr {
		if result[i], err = decodeProperty(prop); err != nil {
			return err
		}
	}

	if err = json.Unmarshal(data, &result); err != nil {
		return err
	}

	*arr = result
	return nil
}

// TitleProperty is a type for title property.
type TitleProperty struct {
	ID    PropertyID   `json:"id,omitempty"`
	Type  PropertyType `json:"type,omitempty"`
	Title RichTexts    `json:"title"`
}

// GetID returns the ID of the TitleProperty.
func (p TitleProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the TitleProperty.
func (p TitleProperty) GetType() PropertyType { return p.Type }

// RichTextProperty is a type for rich text property.
type RichTextProperty struct {
	ID       PropertyID   `json:"id,omitempty"`
	Type     PropertyType `json:"type,omitempty"`
	RichText RichTexts    `json:"rich_text"`
}

// GetID returns the ID of the RichTextProperty.
func (p RichTextProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the RichTextProperty.
func (p RichTextProperty) GetType() PropertyType { return p.Type }

// TextProperty is a type for text property.
type TextProperty struct {
	ID   PropertyID   `json:"id,omitempty"`
	Type PropertyType `json:"type,omitempty"`
	Text RichTexts    `json:"text"`
}

// GetID returns the ID of the TextProperty.
func (p TextProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the TextProperty.
func (p TextProperty) GetType() PropertyType { return p.Type }

// NumberProperty is a type for number property.
type NumberProperty struct {
	ID     PropertyID   `json:"id,omitempty"`
	Type   PropertyType `json:"type,omitempty"`
	Number float64      `json:"number"`
}

// GetID returns the ID of the NumberProperty.
func (p NumberProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the NumberProperty.
func (p NumberProperty) GetType() PropertyType { return p.Type }

// SelectProperty is a type for select property.
type SelectProperty struct {
	ID     ObjectID     `json:"id,omitempty"`
	Type   PropertyType `json:"type,omitempty"`
	Select Option       `json:"select"`
}

// GetID returns the ID of the SelectProperty.
func (p SelectProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the SelectProperty.
func (p SelectProperty) GetType() PropertyType { return p.Type }

// MultiSelectProperty is a type for multi-select property.
type MultiSelectProperty struct {
	ID          ObjectID     `json:"id,omitempty"`
	Type        PropertyType `json:"type,omitempty"`
	MultiSelect []Option     `json:"multi_select"`
}

// GetID returns the ID of the MultiSelectProperty.
func (p MultiSelectProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the MultiSelectProperty.
func (p MultiSelectProperty) GetType() PropertyType { return p.Type }

// Option is a type for option.
type Option struct {
	ID    PropertyID `json:"id,omitempty"`
	Name  string     `json:"name"`
	Color Color      `json:"color,omitempty"`
}

// Options is a slice of Option.
type Options []Option

// DateProperty is a type for date property.
type DateProperty struct {
	ID   ObjectID     `json:"id,omitempty"`
	Type PropertyType `json:"type,omitempty"`
	Date *DateObject  `json:"date"`
}

// DateObject is a type for date object.
type DateObject struct {
	Start *Date `json:"start"`
	End   *Date `json:"end"`
}

// GetID returns the ID of the DateProperty.
func (p DateProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the DateProperty.
func (p DateProperty) GetType() PropertyType { return p.Type }

// FormulaProperty is a type for formula property.
type FormulaProperty struct {
	ID      ObjectID     `json:"id,omitempty"`
	Type    PropertyType `json:"type,omitempty"`
	Formula Formula      `json:"formula"`
}

// FormulaType is a type for formula types.
type FormulaType string

// nolint:revive
const (
	FormulaTypeString  FormulaType = "string"
	FormulaTypeNumber  FormulaType = "number"
	FormulaTypeBoolean FormulaType = "boolean"
	FormulaTypeDate    FormulaType = "date"
)

// Formula is a type for formula.
type Formula struct {
	Type    FormulaType `json:"type,omitempty"`
	String  string      `json:"string,omitempty"`
	Number  float64     `json:"number,omitempty"`
	Boolean bool        `json:"boolean,omitempty"`
	Date    *DateObject `json:"date,omitempty"`
}

// GetID returns the ID of the FormulaProperty.
func (p FormulaProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the FormulaProperty.
func (p FormulaProperty) GetType() PropertyType { return p.Type }

// RelationProperty is a type for relation property.
type RelationProperty struct {
	ID       ObjectID     `json:"id,omitempty"`
	Type     PropertyType `json:"type,omitempty"`
	Relation []Relation   `json:"relation"`
}

// Relation is a type for relation.
type Relation struct {
	ID PageID `json:"id"`
}

// GetID returns the ID of the RelationProperty.
func (p RelationProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the RelationProperty.
func (p RelationProperty) GetType() PropertyType { return p.Type }

// RollupProperty is a type for rollup property.
type RollupProperty struct {
	ID     ObjectID     `json:"id,omitempty"`
	Type   PropertyType `json:"type,omitempty"`
	Rollup Rollup       `json:"rollup"`
}

// RollupType is a type for rollup types.
type RollupType string

// nolint:revive
const (
	RollupTypeNumber RollupType = "number"
	RollupTypeDate   RollupType = "date"
	RollupTypeArray  RollupType = "array"
)

// Rollup is a type for rollup.
type Rollup struct {
	Type   RollupType    `json:"type,omitempty"`
	Number float64       `json:"number,omitempty"`
	Date   *DateObject   `json:"date,omitempty"`
	Array  PropertyArray `json:"array,omitempty"`
}

// GetID returns the ID of the RollupProperty.
func (p RollupProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the RollupProperty.
func (p RollupProperty) GetType() PropertyType { return p.Type }

// PeopleProperty is a type for people property.
type PeopleProperty struct {
	ID     ObjectID     `json:"id,omitempty"`
	Type   PropertyType `json:"type,omitempty"`
	People Users        `json:"people"`
}

// GetID returns the ID of the PeopleProperty.
func (p PeopleProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the PeopleProperty.
func (p PeopleProperty) GetType() PropertyType { return p.Type }

// FilesProperty is a type for files property.
type FilesProperty struct {
	ID    ObjectID     `json:"id,omitempty"`
	Type  PropertyType `json:"type,omitempty"`
	Files Files        `json:"files"`
}

// GetID returns the ID of the FilesProperty.
func (p FilesProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the FilesProperty.
func (p FilesProperty) GetType() PropertyType { return p.Type }

// CheckboxProperty is a type for checkbox property.
type CheckboxProperty struct {
	ID       ObjectID     `json:"id,omitempty"`
	Type     PropertyType `json:"type,omitempty"`
	Checkbox bool         `json:"checkbox"`
}

// GetID returns the ID of the CheckboxProperty.
func (p CheckboxProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the CheckboxProperty.
func (p CheckboxProperty) GetType() PropertyType { return p.Type }

// URLProperty is a type for URL property.
type URLProperty struct {
	ID   ObjectID     `json:"id,omitempty"`
	Type PropertyType `json:"type,omitempty"`
	URL  string       `json:"url"`
}

// GetID returns the ID of the URLProperty.
func (p URLProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the URLProperty.
func (p URLProperty) GetType() PropertyType { return p.Type }

// EmailProperty is a type for email property.
type EmailProperty struct {
	ID    PropertyID   `json:"id,omitempty"`
	Type  PropertyType `json:"type,omitempty"`
	Email string       `json:"email"`
}

// GetID returns the ID of the EmailProperty.
func (p EmailProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the EmailProperty.
func (p EmailProperty) GetType() PropertyType { return p.Type }

// PhoneNumberProperty is a type for phone number property.
type PhoneNumberProperty struct {
	ID          ObjectID     `json:"id,omitempty"`
	Type        PropertyType `json:"type,omitempty"`
	PhoneNumber string       `json:"phone_number"`
}

// GetID returns the ID of the PhoneNumberProperty.
func (p PhoneNumberProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the PhoneNumberProperty.
func (p PhoneNumberProperty) GetType() PropertyType { return p.Type }

// CreatedTimeProperty is a type for created time property.
type CreatedTimeProperty struct {
	ID          ObjectID     `json:"id,omitempty"`
	Type        PropertyType `json:"type,omitempty"`
	CreatedTime time.Time    `json:"created_time"`
}

// GetID returns the ID of the CreatedTimeProperty.
func (p CreatedTimeProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the CreatedTimeProperty.
func (p CreatedTimeProperty) GetType() PropertyType { return p.Type }

// CreatedByProperty is a type for created by property.
type CreatedByProperty struct {
	ID        ObjectID     `json:"id,omitempty"`
	Type      PropertyType `json:"type,omitempty"`
	CreatedBy User         `json:"created_by"`
}

// GetID returns the ID of the CreatedByProperty.
func (p CreatedByProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the CreatedByProperty.
func (p CreatedByProperty) GetType() PropertyType { return p.Type }

// LastEditedTimeProperty is a type for last edited time property.
type LastEditedTimeProperty struct {
	ID             ObjectID     `json:"id,omitempty"`
	Type           PropertyType `json:"type,omitempty"`
	LastEditedTime time.Time    `json:"last_edited_time"`
}

// GetID returns the ID of the LastEditedTimeProperty.
func (p LastEditedTimeProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the LastEditedTimeProperty.
func (p LastEditedTimeProperty) GetType() PropertyType { return p.Type }

// LastEditedByProperty is a type for last edited by property.
type LastEditedByProperty struct {
	ID           ObjectID     `json:"id,omitempty"`
	Type         PropertyType `json:"type,omitempty"`
	LastEditedBy User         `json:"last_edited_by"`
}

// GetID returns the ID of the LastEditedByProperty.
func (p LastEditedByProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the LastEditedBy
func (p LastEditedByProperty) GetType() PropertyType { return p.Type }

// Status is a type for status.
type Status = Option

// StatusProperty is a type for status property.
type StatusProperty struct {
	ID     ObjectID     `json:"id,omitempty"`
	Type   PropertyType `json:"type,omitempty"`
	Status Status       `json:"status"`
}

// GetID returns the ID of the StatusProperty.
func (p StatusProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the StatusProperty.
func (p StatusProperty) GetType() PropertyType { return p.Type }

// UniqueID is a type for unique ID.
type UniqueID struct {
	Prefix *string `json:"prefix,omitempty"`
	Number int     `json:"number"`
}

// String returns the string representation of the UniqueID.
func (uID UniqueID) String() string {
	if uID.Prefix != nil {
		return fmt.Sprintf("%s-%d", *uID.Prefix, uID.Number)
	}
	return fmt.Sprintf("%d", uID.Number)
}

// UniqueIDProperty is a type for unique ID property.
type UniqueIDProperty struct {
	ID       ObjectID     `json:"id,omitempty"`
	Type     PropertyType `json:"type,omitempty"`
	UniqueID UniqueID     `json:"unique_id"`
}

// GetID returns the ID of the UniqueIDProperty.
func (p UniqueIDProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the UniqueIDProperty.
func (p UniqueIDProperty) GetType() PropertyType { return p.Type }

// VerificationProperty is a type for verification property.
type VerificationProperty struct {
	ID           ObjectID     `json:"id,omitempty"`
	Type         PropertyType `json:"type,omitempty"`
	Verification Verification `json:"verification"`
}

// GetID returns the ID of the VerificationProperty.
func (p VerificationProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the VerificationProperty.
func (p VerificationProperty) GetType() PropertyType { return p.Type }

// Button is a type for button.
type Button struct{}

// ButtonProperty is a type for button property.
type ButtonProperty struct {
	ID     ObjectID     `json:"id,omitempty"`
	Type   PropertyType `json:"type,omitempty"`
	Button Button       `json:"button"`
}

// GetID returns the ID of the ButtonProperty.
func (p ButtonProperty) GetID() string { return p.ID.String() }

// GetType returns the Type of the ButtonProperty.
func (p ButtonProperty) GetType() PropertyType { return p.Type }

// Properties is a map of property.
type Properties map[string]Property

// UnmarshalJSON implements custom unmarshalling for Properties
func (p *Properties) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	props, err := parsePageProperties(raw)
	if err != nil {
		return err
	}

	*p = props
	return nil
}

func parsePageProperties(raw map[string]interface{}) (map[string]Property, error) {
	result := make(map[string]Property)
	for k, v := range raw {
		switch rawProperty := v.(type) {
		case map[string]interface{}:
			p, err := decodeProperty(rawProperty)
			if err != nil {
				return nil, err
			}
			b, err := json.Marshal(rawProperty)
			if err != nil {
				return nil, err
			}

			if err = json.Unmarshal(b, &p); err != nil {
				return nil, err
			}

			result[k] = p
		default:
			return nil, fmt.Errorf("unsupported property format %T", v)
		}
	}

	return result, nil
}

func decodeProperty(raw map[string]interface{}) (Property, error) {
	var p Property
	switch PropertyType(raw["type"].(string)) {
	case PropertyTypeTitle:
		p = &TitleProperty{}
	case PropertyTypeRichText:
		p = &RichTextProperty{}
	case PropertyTypeText:
		p = &RichTextProperty{}
	case PropertyTypeNumber:
		p = &NumberProperty{}
	case PropertyTypeSelect:
		p = &SelectProperty{}
	case PropertyTypeMultiSelect:
		p = &MultiSelectProperty{}
	case PropertyTypeDate:
		p = &DateProperty{}
	case PropertyTypeFormula:
		p = &FormulaProperty{}
	case PropertyTypeRelation:
		p = &RelationProperty{}
	case PropertyTypeRollup:
		p = &RollupProperty{}
	case PropertyTypePeople:
		p = &PeopleProperty{}
	case PropertyTypeFiles:
		p = &FilesProperty{}
	case PropertyTypeCheckbox:
		p = &CheckboxProperty{}
	case PropertyTypeURL:
		p = &URLProperty{}
	case PropertyTypeEmail:
		p = &EmailProperty{}
	case PropertyTypePhoneNumber:
		p = &PhoneNumberProperty{}
	case PropertyTypeCreatedTime:
		p = &CreatedTimeProperty{}
	case PropertyTypeCreatedBy:
		p = &CreatedByProperty{}
	case PropertyTypeLastEditedTime:
		p = &LastEditedTimeProperty{}
	case PropertyTypeLastEditedBy:
		p = &LastEditedByProperty{}
	case PropertyTypeStatus:
		p = &StatusProperty{}
	case PropertyTypeUniqueID:
		p = &UniqueIDProperty{}
	case PropertyTypeVerification:
		p = &VerificationProperty{}
	case PropertyTypeButton:
		p = &ButtonProperty{}
	default:
		return nil, fmt.Errorf("unsupported property type: %s", raw["type"].(string))
	}

	return p, nil
}
