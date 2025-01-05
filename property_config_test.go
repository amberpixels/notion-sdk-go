package notion_test

import (
	"encoding/json"
	"reflect"
	"testing"

	notion "github.com/amberpixels/notion-sdk-go"
)

func TestNumberPropertyConfig_MarshalJSON(t *testing.T) {
	type fields struct {
		Type   notion.PropertyConfigType
		Format notion.FormatType
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "returns correct json",
			fields: fields{
				Type:   notion.PropertyConfigTypeNumber,
				Format: notion.FormatDollar,
			},
			want:    []byte(`{"type":"number","number":{"format":"dollar"}}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := notion.NumberPropertyConfig{
				Type:   tt.fields.Type,
				Number: notion.NumberFormat{Format: tt.fields.Format},
			}
			got, err := json.Marshal(p)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
