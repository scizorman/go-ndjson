package ndjson

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUnmarshal(t *testing.T) {
	type user struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Bio  string `json:"bio"`
	}
	data := []byte(`{"id":1,"name":"John Smith","bio":"I am John.\nI like football."}
{"id":2,"name":"Ashley Madison","bio":"I am Ashley.\nI like baseball."}
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
			name: "Map",
			args: args{
				data: data,
				v:    new([]map[string]interface{}),
			},
			want: &[]map[string]interface{}{
				{"id": float64(1), "name": "John Smith", "bio": "I am John.\nI like football."},
				{"id": float64(2), "name": "Ashley Madison", "bio": "I am Ashley.\nI like baseball."},
			},
			wantErr: false,
		},
		{
			name: "Struct",
			args: args{
				data: data,
				v:    new([]user),
			},
			want: &[]user{
				{1, "John Smith", "I am John.\nI like football."},
				{2, "Ashley Madison", "I am Ashley.\nI like baseball."},
			},
			wantErr: false,
		},
		{
			name: "PointerToStruct",
			args: args{
				data: data,
				v:    new([]*user),
			},
			want: &[]*user{
				{1, "John Smith", "I am John.\nI like football."},
				{2, "Ashley Madison", "I am Ashley.\nI like baseball."},
			},
			wantErr: false,
		},
		{
			name: "NotPointer",
			args: args{
				data: data,
				v:    []user{},
			},
			wantErr: true,
		},
		{
			name: "NotPointerToSlice",
			args: args{
				data: data,
				v:    new(int),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Unmarshal(tt.args.data, tt.args.v)
			t.Logf("err: %v", err)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if !cmp.Equal(tt.args.v, tt.want) {
					t.Errorf("Unmarshal() = %v, want %v\ndiff: %s", tt.args.v, &tt.want, cmp.Diff(tt.args.v, tt.want))
				}
			}
		})
	}
}
