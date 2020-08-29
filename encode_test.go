package ndjson

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMarshal(t *testing.T) {
	type user struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Bio  string `json:"bio"`
	}

	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Map",
			args: args{
				v: []map[string]interface{}{
					{"id": float64(1), "name": "John Smith", "bio": "I am John.\nI like football."},
					{"id": float64(2), "name": "Ashley Madison", "bio": "I am Ashley.\nI like baseball."},
				},
			},
			want: []byte(`{"bio":"I am John.\nI like football.","id":1,"name":"John Smith"}
{"bio":"I am Ashley.\nI like baseball.","id":2,"name":"Ashley Madison"}
`),
			wantErr: false,
		},
		{
			name: "Struct",
			args: args{
				v: []user{
					{1, "John Smith", "I am John.\nI like football."},
					{2, "Ashley Madison", "I am Ashley.\nI like baseball."},
				},
			},
			want: []byte(`{"id":1,"name":"John Smith","bio":"I am John.\nI like football."}
{"id":2,"name":"Ashley Madison","bio":"I am Ashley.\nI like baseball."}
`),
			wantErr: false,
		},
		{
			name: "PointerToStruct",
			args: args{
				v: []*user{
					{1, "John Smith", "I am John.\nI like football."},
					{2, "Ashley Madison", "I am Ashley.\nI like baseball."},
				},
			},
			want: []byte(`{"id":1,"name":"John Smith","bio":"I am John.\nI like football."}
{"id":2,"name":"Ashley Madison","bio":"I am Ashley.\nI like baseball."}
`),
			wantErr: false,
		},
		{
			name: "NonSlice",
			args: args{
				// v: user{1, "John Smith", "I am John.\nI like football."},
				v: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Marshal() got = %v, want %v\ndiff: %s", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}
