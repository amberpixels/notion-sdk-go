package notion

import (
	"encoding/json"
)

// FilterOperator is a type for filter operators.
type FilterOperator string

// nolint:revive
const (
	FilterOperatorAND FilterOperator = "and"
	FilterOperatorOR  FilterOperator = "or"
)

// Filter is an interface for filter types.
// TODO: refactor probably
type Filter interface {
	filter()
}

// CompoundFilter is a type for compound filters.
type CompoundFilter map[FilterOperator][]PropertyFilter

// AndCompoundFilter is a type for `and` compound filters.
type AndCompoundFilter []Filter

// OrCompoundFilter is a type for `or` compound filters.
type OrCompoundFilter []Filter

func (f AndCompoundFilter) filter() {}
func (f OrCompoundFilter) filter()  {}

// MarshalJSON implements custom marshalling for AndCompoundFilter and OrCompoundFilter
func (f AndCompoundFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		And []Filter `json:"and"`
	}{
		And: f,
	})
}

// MarshalJSON implements custom marshalling for OrCompoundFilter
func (f OrCompoundFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Or []Filter `json:"or"`
	}{
		Or: f,
	})
}

// Condition is a type for filter conditions.
type Condition string

// nolint:revive
const (
	ConditionEquals         Condition = "equals"
	ConditionDoesNotEqual   Condition = "does_not_equal"
	ConditionContains       Condition = "contains"
	ConditionDoesNotContain Condition = "does_not_contain"
	ConditionDoesStartsWith Condition = "starts_with"
	ConditionDoesEndsWith   Condition = "ends_with"
	ConditionDoesIsEmpty    Condition = "is_empty"
	ConditionGreaterThan    Condition = "greater_than"
	ConditionLessThan       Condition = "less_than"

	ConditionGreaterThanOrEqualTo Condition = "greater_than_or_equal_to"
	ConditionLessThanOrEqualTo    Condition = "greater_than_or_equal_to"

	ConditionBefore     Condition = "before"
	ConditionAfter      Condition = "after"
	ConditionOnOrBefore Condition = "on_or_before"
	ConditionOnOrAfter  Condition = "on_or_after"
	ConditionPastWeek   Condition = "past_week"
	ConditionPastMonth  Condition = "past_month"
	ConditionPastYear   Condition = "past_year"
	ConditionNextWeek   Condition = "next_week"
	ConditionNextMonth  Condition = "next_month"
	ConditionNextYear   Condition = "next_year"

	ConditionText     Condition = "text"
	ConditionCheckbox Condition = "checkbox"
	ConditionNumber   Condition = "number"
	ConditionDate     Condition = "date"
)

// TimestampFilter is a type for timestamp filters.
type TimestampFilter struct {
	Timestamp      TimestampType        `json:"timestamp"`
	CreatedTime    *DateFilterCondition `json:"created_time,omitempty"`
	LastEditedTime *DateFilterCondition `json:"last_edited_time,omitempty"`
}

func (f TimestampFilter) filter() {}

// PropertyFilter is a type for property filters.
type PropertyFilter struct {
	Property    string                      `json:"property"`
	RichText    *TextFilterCondition        `json:"rich_text,omitempty"`
	Number      *NumberFilterCondition      `json:"number,omitempty"`
	Checkbox    *CheckboxFilterCondition    `json:"checkbox,omitempty"`
	Select      *SelectFilterCondition      `json:"select,omitempty"`
	MultiSelect *MultiSelectFilterCondition `json:"multi_select,omitempty"`
	Date        *DateFilterCondition        `json:"date,omitempty"`
	People      *PeopleFilterCondition      `json:"people,omitempty"`
	Files       *FilesFilterCondition       `json:"files,omitempty"`
	Relation    *RelationFilterCondition    `json:"relation,omitempty"`
	Formula     *FormulaFilterCondition     `json:"formula,omitempty"`
	Rollup      *RollupFilterCondition      `json:"rollup,omitempty"`
	Status      *StatusFilterCondition      `json:"status,omitempty"`
	UniqueID    *UniqueIDFilterCondition    `json:"unique_id,omitempty"`
}

func (f PropertyFilter) filter() {}

// SearchFilter is a type for search filters.
type SearchFilter struct {
	Value    string `json:"value"`
	Property string `json:"property"`
}

// TextFilterCondition is a type for text filter conditions.
type TextFilterCondition struct {
	Equals         string `json:"equals,omitempty"`
	DoesNotEqual   string `json:"does_not_equal,omitempty"`
	Contains       string `json:"contains,omitempty"`
	DoesNotContain string `json:"does_not_contain,omitempty"`
	StartsWith     string `json:"starts_with,omitempty"`
	EndsWith       string `json:"ends_with,omitempty"`
	IsEmpty        bool   `json:"is_empty,omitempty"`
	IsNotEmpty     bool   `json:"is_not_empty,omitempty"`
}

// NumberFilterCondition is a type for number filter conditions.
type NumberFilterCondition struct {
	Equals               *float64 `json:"equals,omitempty"`
	DoesNotEqual         *float64 `json:"does_not_equal,omitempty"`
	GreaterThan          *float64 `json:"greater_than,omitempty"`
	LessThan             *float64 `json:"less_than,omitempty"`
	GreaterThanOrEqualTo *float64 `json:"greater_than_or_equal_to,omitempty"`
	LessThanOrEqualTo    *float64 `json:"less_than_or_equal_to,omitempty"`
	IsEmpty              bool     `json:"is_empty,omitempty"`
	IsNotEmpty           bool     `json:"is_not_empty,omitempty"`
}

// CheckboxFilterCondition is a type for checkbox filter conditions.
type CheckboxFilterCondition struct {
	Equals       bool `json:"equals,omitempty"`
	DoesNotEqual bool `json:"does_not_equal,omitempty"`
}

// SelectFilterCondition is a type for select filter conditions.
type SelectFilterCondition struct {
	Equals       string `json:"equals,omitempty"`
	DoesNotEqual string `json:"does_not_equal,omitempty"`
	IsEmpty      bool   `json:"is_empty,omitempty"`
	IsNotEmpty   bool   `json:"is_not_empty,omitempty"`
}

// MultiSelectFilterCondition is a type for multi-select filter conditions.
type MultiSelectFilterCondition struct {
	Contains       string `json:"contains,omitempty"`
	DoesNotContain string `json:"does_not_contain,omitempty"`
	IsEmpty        bool   `json:"is_empty,omitempty"`
	IsNotEmpty     bool   `json:"is_not_empty,omitempty"`
}

// DateFilterCondition is a type for date
type DateFilterCondition struct {
	Equals     *Date     `json:"equals,omitempty"`
	Before     *Date     `json:"before,omitempty"`
	After      *Date     `json:"after,omitempty"`
	OnOrBefore *Date     `json:"on_or_before,omitempty"`
	OnOrAfter  *Date     `json:"on_or_after,omitempty"`
	PastWeek   *struct{} `json:"past_week,omitempty"`
	PastMonth  *struct{} `json:"past_month,omitempty"`
	PastYear   *struct{} `json:"past_year,omitempty"`
	NextWeek   *struct{} `json:"next_week,omitempty"`
	NextMonth  *struct{} `json:"next_month,omitempty"`
	NextYear   *struct{} `json:"next_year,omitempty"`
	IsEmpty    bool      `json:"is_empty,omitempty"`
	IsNotEmpty bool      `json:"is_not_empty,omitempty"`
}

// PeopleFilterCondition is a type for people filter conditions.
type PeopleFilterCondition struct {
	Contains       string `json:"contains,omitempty"`
	DoesNotContain string `json:"does_not_contain,omitempty"`
	IsEmpty        bool   `json:"is_empty,omitempty"`
	IsNotEmpty     bool   `json:"is_not_empty,omitempty"`
}

// FilesFilterCondition is a type for files filter conditions.
type FilesFilterCondition struct {
	IsEmpty    bool `json:"is_empty,omitempty"`
	IsNotEmpty bool `json:"is_not_empty,omitempty"`
}

// RelationFilterCondition is a type for relation filter conditions.
type RelationFilterCondition struct {
	Contains       string `json:"contains,omitempty"`
	DoesNotContain string `json:"does_not_contain,omitempty"`
	IsEmpty        bool   `json:"is_empty,omitempty"`
	IsNotEmpty     bool   `json:"is_not_empty,omitempty"`
}

// FormulaFilterCondition is a type for formula filter conditions.
type FormulaFilterCondition struct {
	// DEPRECATED use `String` instead
	Text     *TextFilterCondition     `json:"text,omitempty"`
	String   *TextFilterCondition     `json:"string,omitempty"`
	Checkbox *CheckboxFilterCondition `json:"checkbox,omitempty"`
	Number   *NumberFilterCondition   `json:"number,omitempty"`
	Date     *DateFilterCondition     `json:"date,omitempty"`
}

// RollupFilterCondition is a type for rollup filter conditions.
type RollupFilterCondition struct {
	Any    *RollupSubfilterCondition `json:"any,omitempty"`
	None   *RollupSubfilterCondition `json:"none,omitempty"`
	Every  *RollupSubfilterCondition `json:"every,omitempty"`
	Date   *DateFilterCondition      `json:"date,omitempty"`
	Number *NumberFilterCondition    `json:"number,omitempty"`
}

// RollupSubfilterCondition is a type for rollup subfilter conditions.
type RollupSubfilterCondition struct {
	RichText    *TextFilterCondition        `json:"rich_text,omitempty"`
	Number      *NumberFilterCondition      `json:"number,omitempty"`
	Checkbox    *CheckboxFilterCondition    `json:"checkbox,omitempty"`
	Select      *SelectFilterCondition      `json:"select,omitempty"`
	MultiSelect *MultiSelectFilterCondition `json:"multiSelect,omitempty"`
	Relation    *RelationFilterCondition    `json:"relation,omitempty"`
	Date        *DateFilterCondition        `json:"date,omitempty"`
	People      *PeopleFilterCondition      `json:"people,omitempty"`
	Files       *FilesFilterCondition       `json:"files,omitempty"`
}

// StatusFilterCondition is a type for status filter conditions.
type StatusFilterCondition struct {
	Equals       string `json:"equals,omitempty"`
	DoesNotEqual string `json:"does_not_equal,omitempty"`
	IsEmpty      bool   `json:"is_empty,omitempty"`
	IsNotEmpty   bool   `json:"is_not_empty,omitempty"`
}

// UniqueIDFilterCondition is a type for unique ID filter conditions.
type UniqueIDFilterCondition struct {
	Equals               *int `json:"equals,omitempty"`
	DoesNotEqual         *int `json:"does_not_equal,omitempty"`
	GreaterThan          *int `json:"greater_than,omitempty"`
	LessThan             *int `json:"less_than,omitempty"`
	GreaterThanOrEqualTo *int `json:"greater_than_or_equal_to,omitempty"`
	LessThanOrEqualTo    *int `json:"less_than_or_equal_to,omitempty"`
}
