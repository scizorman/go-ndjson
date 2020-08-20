package ndjson

import (
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	type user struct {
		ID  int64  `json:"id"`
		Bio string `json:"bio"`
	}
	p := []byte(`{"id":1,"bio":"I am John."}
{"id":2,"bio":"I am Paul."}
`)

	type args struct {
		data []byte
		v    interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Interface",
			args: args{
				data: p,
			},
			want: []interface{}{
				[]map[string]interface{}{
					{"id": 1, "bio": "I am a dog."},
					{"id": 2, "bio": "I am a cat."},
				},
			},
			wantErr: false,
		},
		{
			name: "Map",
			args: args{
				data: p,
				v:    &[]map[string]interface{}{},
			},
			want: &[]map[string]interface{}{
				{"id": 1, "bio": "I am a dog."},
				{"id": 2, "bio": "I am a cat."},
			},
			wantErr: false,
		},
		{
			name: "Struct",
			args: args{
				data: p,
				v:    &[]user{},
			},
			want: &[]user{
				{1, "I am a dog."},
				{2, "I am a cat."},
			},
			wantErr: false,
		},
		{
			name: "PointerToStruct",
			args: args{
				data: p,
				v:    &[]*user{},
			},
			want: &[]*user{
				{1, "I am a dog."},
				{2, "I am a cat."},
			},
		},
		{
			name: "NotPointerToSlice",
			args: args{
				data: p,
				v:    []user{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.args.v, Unmarshal(tt.args.data, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
