package notion

import (
	"encoding/json"
	"fmt"
)

// PropertyConfigType is a type for property config types.
type PropertyConfigType string

// nolint:revive
const (
	PropertyConfigTypeTitle       PropertyConfigType = "title"
	PropertyConfigTypeRichText    PropertyConfigType = "rich_text"
	PropertyConfigTypeNumber      PropertyConfigType = "number"
	PropertyConfigTypeSelect      PropertyConfigType = "select"
	PropertyConfigTypeMultiSelect PropertyConfigType = "multi_select"
	PropertyConfigTypeDate        PropertyConfigType = "date"
	PropertyConfigTypePeople      PropertyConfigType = "people"
	PropertyConfigTypeFiles       PropertyConfigType = "files"
	PropertyConfigTypeCheckbox    PropertyConfigType = "checkbox"
	PropertyConfigTypeURL         PropertyConfigType = "url"
	PropertyConfigTypeEmail       PropertyConfigType = "email"
	PropertyConfigTypePhoneNumber PropertyConfigType = "phone_number"
	PropertyConfigTypeFormula     PropertyConfigType = "formula"
	PropertyConfigTypeRelation    PropertyConfigType = "relation"
	PropertyConfigTypeRollup      PropertyConfigType = "rollup"
	PropertyConfigCreatedTime     PropertyConfigType = "created_time"
	PropertyConfigCreatedBy       PropertyConfigType = "created_by"
	PropertyConfigLastEditedTime  PropertyConfigType = "last_edited_time"
	PropertyConfigLastEditedBy    PropertyConfigType = "last_edited_by"
	PropertyConfigStatus          PropertyConfigType = "status"
	PropertyConfigUniqueID        PropertyConfigType = "unique_id"
	PropertyConfigVerification    PropertyConfigType = "verification"
	PropertyConfigButton          PropertyConfigType = "button"
)

// PropertyID is a type for property IDs.
type PropertyID string

// String returns the string representation of the PropertyID.
func (pID PropertyID) String() string { return string(pID) }

// PropertyConfig is an interface for property configs.
type PropertyConfig interface {
	GetType() PropertyConfigType
	GetID() PropertyID
}

// TitlePropertyConfig is a type for title property configs.
type TitlePropertyConfig struct {
	ID    PropertyID         `json:"id,omitempty"`
	Type  PropertyConfigType `json:"type"`
	Title any                `json:"title"`
}

// GetType returns the Type of the TitlePropertyConfig.
func (p TitlePropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the TitlePropertyConfig.
func (p TitlePropertyConfig) GetID() PropertyID { return p.ID }

// RichTextPropertyConfig is a type for rich text property configs.
type RichTextPropertyConfig struct {
	ID       PropertyID         `json:"id,omitempty"`
	Type     PropertyConfigType `json:"type"`
	RichText any                `json:"rich_text"`
}

// GetType returns the Type of the RichTextPropertyConfig.
func (p RichTextPropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the RichTextPropertyConfig.
func (p RichTextPropertyConfig) GetID() PropertyID { return p.ID }

// NumberPropertyConfig is a type for number property configs.
type NumberPropertyConfig struct {
	ID     PropertyID         `json:"id,omitempty"`
	Type   PropertyConfigType `json:"type"`
	Number NumberFormat       `json:"number"`
}

// FormatType is a type for format types.
type FormatType string

// String returns the string representation of the FormatType.
func (ft FormatType) String() string { return string(ft) }

// nolint:revive
const (
	FormatNumber           FormatType = "number"
	FormatNumberWithCommas FormatType = "number_with_commas"
	FormatPercent          FormatType = "percent"
	FormatDollar           FormatType = "dollar"
	FormatCanadianDollar   FormatType = "canadian_dollar"
	FormatEuro             FormatType = "euro"
	FormatPound            FormatType = "pound"
	FormatYen              FormatType = "yen"
	FormatRuble            FormatType = "ruble"
	FormatRupee            FormatType = "rupee"
	FormatWon              FormatType = "won"
	FormatYuan             FormatType = "yuan"
	FormatReal             FormatType = "real"
	FormatLira             FormatType = "lira"
	FormatRupiah           FormatType = "rupiah"
	FormatFranc            FormatType = "franc"
	FormatHongKongDollar   FormatType = "hong_kong_dollar"
	FormatNewZealandDollar FormatType = "hong_kong_dollar"
	FormatKrona            FormatType = "krona"
	FormatNorwegianKrone   FormatType = "norwegian_krone"
	FormatMexicanPeso      FormatType = "mexican_peso"
	FormatRand             FormatType = "rand"
	FormatNewTaiwanDollar  FormatType = "new_taiwan_dollar"
	FormatDanishKrone      FormatType = "danish_krone"
	FormatZloty            FormatType = "zloty"
	FormatBath             FormatType = "baht"
	FormatForint           FormatType = "forint"
	FormatKoruna           FormatType = "koruna"
	FormatShekel           FormatType = "shekel"
	FormatChileanPeso      FormatType = "chilean_peso"
	FormatPhilippinePeso   FormatType = "philippine_peso"
	FormatDirham           FormatType = "dirham"
	FormatColombianPeso    FormatType = "colombian_peso"
	FormatRiyal            FormatType = "riyal"
	FormatRinggit          FormatType = "ringgit"
	FormatLeu              FormatType = "leu"
	FormatArgentinePeso    FormatType = "argentine_peso"
	FormatUruguayanPeso    FormatType = "uruguayan_peso"
	FormatSingaporeDollar  FormatType = "singapore_dollar"
)

// NumberFormat is a type for number formats.
type NumberFormat struct {
	Format FormatType `json:"format"`
}

// GetType returns the Type of the NumberPropertyConfig.
func (p NumberPropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the NumberPropertyConfig.
func (p NumberPropertyConfig) GetID() PropertyID { return p.ID }

// SelectPropertyConfig is a type for select property configs.
type SelectPropertyConfig struct {
	ID     PropertyID         `json:"id,omitempty"`
	Type   PropertyConfigType `json:"type"`
	Select Select             `json:"select"`
}

// GetType returns the Type of the SelectPropertyConfig.
func (p SelectPropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the SelectPropertyConfig.
func (p SelectPropertyConfig) GetID() PropertyID { return p.ID }

// MultiSelectPropertyConfig is a type for multi-select property configs.
type MultiSelectPropertyConfig struct {
	ID          PropertyID         `json:"id,omitempty"`
	Type        PropertyConfigType `json:"type"`
	MultiSelect Select             `json:"multi_select"`
}

// Select is a type for select configs.
type Select struct {
	Options Options `json:"options"`
}

// GetType returns the Type of the MultiSelectPropertyConfig.
func (p MultiSelectPropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the MultiSelectPropertyConfig.
func (p MultiSelectPropertyConfig) GetID() PropertyID { return p.ID }

// DatePropertyConfig is a type for date property configs.
type DatePropertyConfig struct {
	ID   PropertyID         `json:"id,omitempty"`
	Type PropertyConfigType `json:"type"`
	Date any                `json:"date"`
}

// GetType returns the Type of the DatePropertyConfig.
func (p DatePropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the DatePropertyConfig.
func (p DatePropertyConfig) GetID() PropertyID { return p.ID }

// PeoplePropertyConfig is a type for people property configs.
type PeoplePropertyConfig struct {
	ID     PropertyID         `json:"id,omitempty"`
	Type   PropertyConfigType `json:"type"`
	People any                `json:"people"`
}

// GetType returns the Type of the PeoplePropertyConfig.
func (p PeoplePropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the PeoplePropertyConfig.
func (p PeoplePropertyConfig) GetID() PropertyID { return p.ID }

// FilesPropertyConfig is a type for files property configs.
type FilesPropertyConfig struct {
	ID    PropertyID         `json:"id,omitempty"`
	Type  PropertyConfigType `json:"type"`
	Files any                `json:"files"`
}

// GetType returns the Type of the FilesPropertyConfig.
func (p FilesPropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the FilesPropertyConfig.
func (p FilesPropertyConfig) GetID() PropertyID { return p.ID }

// CheckboxPropertyConfig is a type for checkbox property configs.
type CheckboxPropertyConfig struct {
	ID       PropertyID         `json:"id,omitempty"`
	Type     PropertyConfigType `json:"type"`
	Checkbox any                `json:"checkbox"`
}

// GetType returns the Type of the CheckboxPropertyConfig.
func (p CheckboxPropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the CheckboxPropertyConfig.
func (p CheckboxPropertyConfig) GetID() PropertyID { return p.ID }

// URLPropertyConfig is a type for URL property configs.
type URLPropertyConfig struct {
	ID   PropertyID         `json:"id,omitempty"`
	Type PropertyConfigType `json:"type"`
	URL  any                `json:"url"`
}

// GetType returns the Type of the URLPropertyConfig.
func (p URLPropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the URLPropertyConfig.
func (p URLPropertyConfig) GetID() PropertyID { return p.ID }

// EmailPropertyConfig is a type for email property configs.
type EmailPropertyConfig struct {
	ID    PropertyID         `json:"id,omitempty"`
	Type  PropertyConfigType `json:"type"`
	Email any                `json:"email"`
}

// GetType returns the Type of the EmailPropertyConfig.
func (p EmailPropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the EmailPropertyConfig.
func (p EmailPropertyConfig) GetID() PropertyID { return p.ID }

// PhoneNumberPropertyConfig is a type for phone number property configs.
type PhoneNumberPropertyConfig struct {
	ID          PropertyID         `json:"id,omitempty"`
	Type        PropertyConfigType `json:"type"`
	PhoneNumber any                `json:"phone_number"`
}

// GetType returns the Type of the PhoneNumberPropertyConfig.s
func (p PhoneNumberPropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the PhoneNumberPropertyConfig.
func (p PhoneNumberPropertyConfig) GetID() PropertyID { return p.ID }

// FormulaPropertyConfig is a type for formula property configs.
type FormulaPropertyConfig struct {
	ID      PropertyID         `json:"id,omitempty"`
	Type    PropertyConfigType `json:"type"`
	Formula FormulaConfig      `json:"formula"`
}

// FormulaConfig is a type for formula configs.
type FormulaConfig struct {
	Expression string `json:"expression"`
}

// GetType returns the Type of the FormulaPropertyConfig.
func (p FormulaPropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the FormulaPropertyConfig.
func (p FormulaPropertyConfig) GetID() PropertyID { return p.ID }

// RelationPropertyConfig is a type for relation property configs.
type RelationPropertyConfig struct {
	Type     PropertyConfigType `json:"type"`
	Relation RelationConfig     `json:"relation"`
}

// RelationConfigType is a type for relation config types.
type RelationConfigType string

// nolint:revive
const (
	RelationSingleProperty RelationConfigType = "single_property"
	RelationDualProperty   RelationConfigType = "dual_property"
)

// String returns the string representation of the RelationConfigType.
func (rp RelationConfigType) String() string { return string(rp) }

// SingleProperty is a type for single properties.
type SingleProperty struct{}

// DualProperty is a type for dual properties.
type DualProperty struct{}

// RelationConfig is a type for relation configs.
type RelationConfig struct {
	DatabaseID         DatabaseID         `json:"database_id"`
	SyncedPropertyID   PropertyID         `json:"synced_property_id,omitempty"`
	SyncedPropertyName string             `json:"synced_property_name,omitempty"`
	Type               RelationConfigType `json:"type,omitempty"`
	SingleProperty     *SingleProperty    `json:"single_property,omitempty"`
	DualProperty       *DualProperty      `json:"dual_property,omitempty"`
}

// GetType returns the Type of the RelationPropertyConfig.
func (p RelationPropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the RelationPropertyConfig.
func (p RelationPropertyConfig) GetID() PropertyID { return "" }

// RollupPropertyConfig is a type for rollup property configs.
type RollupPropertyConfig struct {
	ID     PropertyID         `json:"id,omitempty"`
	Type   PropertyConfigType `json:"type"`
	Rollup RollupConfig       `json:"rollup"`
}

// RollupConfig is a type for rollup configs.
type RollupConfig struct {
	RelationPropertyName string       `json:"relation_property_name"`
	RelationPropertyID   PropertyID   `json:"relation_property_id"`
	RollupPropertyName   string       `json:"rollup_property_name"`
	RollupPropertyID     PropertyID   `json:"rollup_property_id"`
	Function             FunctionType `json:"function"`
}

// FunctionType is a type for function types.
type FunctionType string

// String returns the string representation of the FunctionType.
func (ft FunctionType) String() string { return string(ft) }

// nolint:revive
const (
	FunctionCountAll          FunctionType = "count_all"
	FunctionCountValues       FunctionType = "count_values"
	FunctionCountUniqueValues FunctionType = "count_unique_values"
	FunctionCountEmpty        FunctionType = "count_empty"
	FunctionCountNotEmpty     FunctionType = "count_not_empty"
	FunctionPercentEmpty      FunctionType = "percent_empty"
	FunctionPercentNotEmpty   FunctionType = "percent_not_empty"
	FunctionSum               FunctionType = "sum"
	FunctionAverage           FunctionType = "average"
	FunctionMedian            FunctionType = "median"
	FunctionMin               FunctionType = "min"
	FunctionMax               FunctionType = "max"
	FunctionRange             FunctionType = "range"
)

// GetType returns the Type of the RollupPropertyConfig.
func (p RollupPropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the RollupPropertyConfig.
func (p RollupPropertyConfig) GetID() PropertyID { return p.ID }

// CreatedTimePropertyConfig is a type for created time property configs.
type CreatedTimePropertyConfig struct {
	ID          PropertyID         `json:"id,omitempty"`
	Type        PropertyConfigType `json:"type"`
	CreatedTime any                `json:"created_time"`
}

// GetType returns the Type of the CreatedTimePropertyConfig.
func (p CreatedTimePropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the CreatedTimePropertyConfig.
func (p CreatedTimePropertyConfig) GetID() PropertyID { return p.ID }

// CreatedByPropertyConfig is a type for created by property configs.
type CreatedByPropertyConfig struct {
	ID        PropertyID         `json:"id"`
	Type      PropertyConfigType `json:"type"`
	CreatedBy any                `json:"created_by"`
}

// GetType returns the Type of the CreatedByPropertyConfig.
func (p CreatedByPropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the CreatedByPropertyConfig.
func (p CreatedByPropertyConfig) GetID() PropertyID { return p.ID }

// LastEditedTimePropertyConfig is a type for last edited time property configs.
type LastEditedTimePropertyConfig struct {
	ID             PropertyID         `json:"id"`
	Type           PropertyConfigType `json:"type"`
	LastEditedTime any                `json:"last_edited_time"`
}

// GetType returns the Type of the LastEditedTimePropertyConfig.
func (p LastEditedTimePropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the LastEditedTimePropertyConfig.
func (p LastEditedTimePropertyConfig) GetID() PropertyID { return p.ID }

// LastEditedByPropertyConfig is a type for last edited by property configs.
type LastEditedByPropertyConfig struct {
	ID           PropertyID         `json:"id"`
	Type         PropertyConfigType `json:"type"`
	LastEditedBy any                `json:"last_edited_by"`
}

// GetType returns the Type of the LastEditedByPropertyConfig.
func (p LastEditedByPropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the LastEditedByPropertyConfig.
func (p LastEditedByPropertyConfig) GetID() PropertyID { return p.ID }

// StatusPropertyConfig is a type for status property configs.
type StatusPropertyConfig struct {
	ID     PropertyID         `json:"id"`
	Type   PropertyConfigType `json:"type"`
	Status StatusConfig       `json:"status"`
}

// GetType returns the Type of the StatusPropertyConfig.
func (p StatusPropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the StatusPropertyConfig.
func (p StatusPropertyConfig) GetID() PropertyID { return p.ID }

// StatusConfig is a type for status configs.
type StatusConfig struct {
	Options Options       `json:"options"`
	Groups  []GroupConfig `json:"groups"`
}

// GroupConfig is a type for group configs.
type GroupConfig struct {
	ID        ObjectID   `json:"id"`
	Name      string     `json:"name"`
	Color     string     `json:"color"`
	OptionIDs []ObjectID `json:"option_ids"`
}

// UniqueIDPropertyConfig is a type for unique ID property configs.
type UniqueIDPropertyConfig struct {
	ID       PropertyID         `json:"id,omitempty"`
	Type     PropertyConfigType `json:"type"`
	UniqueID UniqueIDConfig     `json:"unique_id"`
}

// UniqueIDConfig is a type for unique ID configs.
type UniqueIDConfig struct {
	Prefix string `json:"prefix"`
}

// GetType returns the Type of the UniqueIDPropertyConfig.
func (p UniqueIDPropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the UniqueIDPropertyConfig.
func (p UniqueIDPropertyConfig) GetID() PropertyID { return p.ID }

// VerificationState is a type for verification states.
type VerificationState string

// nolint:revive
const (
	VerificationStateVerified   VerificationState = "verified"
	VerificationStateUnverified VerificationState = "unverified"
)

// String returns the string representation of the VerificationState.
func (vs VerificationState) String() string { return string(vs) }

// Verification documented here: https://developers.notion.com/reference/page-property-values#verification
type Verification struct {
	State      VerificationState `json:"state"`
	VerifiedBy *User             `json:"verified_by,omitempty"`
	Date       *DateObject       `json:"date,omitempty"`
}

// VerificationPropertyConfig is a type for verification property configs.
type VerificationPropertyConfig struct {
	ID           PropertyID         `json:"id,omitempty"`
	Type         PropertyConfigType `json:"type,omitempty"`
	Verification Verification       `json:"verification"`
}

// GetType returns the Type of the VerificationPropertyConfig.
func (p VerificationPropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the VerificationPropertyConfig.
func (p VerificationPropertyConfig) GetID() PropertyID { return p.ID }

// ButtonPropertyConfig is a type for button property configs.
type ButtonPropertyConfig struct {
	ID     PropertyID         `json:"id"`
	Type   PropertyConfigType `json:"type"`
	Button any                `json:"button"`
}

// GetType returns the Type of the ButtonPropertyConfig.
func (p ButtonPropertyConfig) GetType() PropertyConfigType { return p.Type }

// GetID returns the ID of the ButtonPropertyConfig.
func (p ButtonPropertyConfig) GetID() PropertyID { return p.ID }

// PropertyConfigs is a map of property configs.
type PropertyConfigs map[string]PropertyConfig

// UnmarshalJSON implements custom unmarshalling for PropertyConfigs
func (p *PropertyConfigs) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	props, err := parsePropertyConfigs(raw)
	if err != nil {
		return err
	}

	*p = props
	return nil
}

func parsePropertyConfigs(raw map[string]interface{}) (PropertyConfigs, error) {
	result := make(PropertyConfigs)
	for k, v := range raw {
		var p PropertyConfig
		switch rawProperty := v.(type) {
		case map[string]interface{}:
			switch PropertyConfigType(rawProperty["type"].(string)) {
			case PropertyConfigTypeTitle:
				p = &TitlePropertyConfig{}
			case PropertyConfigTypeRichText:
				p = &RichTextPropertyConfig{}
			case PropertyConfigTypeNumber:
				p = &NumberPropertyConfig{}
			case PropertyConfigTypeSelect:
				p = &SelectPropertyConfig{}
			case PropertyConfigTypeMultiSelect:
				p = &MultiSelectPropertyConfig{}
			case PropertyConfigTypeDate:
				p = &DatePropertyConfig{}
			case PropertyConfigTypePeople:
				p = &PeoplePropertyConfig{}
			case PropertyConfigTypeFiles:
				p = &FilesPropertyConfig{}
			case PropertyConfigTypeCheckbox:
				p = &CheckboxPropertyConfig{}
			case PropertyConfigTypeURL:
				p = &URLPropertyConfig{}
			case PropertyConfigTypeEmail:
				p = &EmailPropertyConfig{}
			case PropertyConfigTypePhoneNumber:
				p = &PhoneNumberPropertyConfig{}
			case PropertyConfigTypeFormula:
				p = &FormulaPropertyConfig{}
			case PropertyConfigTypeRelation:
				p = &RelationPropertyConfig{}
			case PropertyConfigTypeRollup:
				p = &RollupPropertyConfig{}
			case PropertyConfigCreatedTime:
				p = &CreatedTimePropertyConfig{}
			case PropertyConfigCreatedBy:
				p = &CreatedTimePropertyConfig{}
			case PropertyConfigLastEditedTime:
				p = &LastEditedTimePropertyConfig{}
			case PropertyConfigLastEditedBy:
				p = &LastEditedByPropertyConfig{}
			case PropertyConfigStatus:
				p = &StatusPropertyConfig{}
			case PropertyConfigUniqueID:
				p = &UniqueIDPropertyConfig{}
			case PropertyConfigVerification:
				p = &VerificationPropertyConfig{}
			case PropertyConfigButton:
				p = &ButtonPropertyConfig{}
			default:

				return nil, fmt.Errorf("unsupported property type: %s", rawProperty["type"].(string))
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
