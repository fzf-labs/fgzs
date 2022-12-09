package crypt

import (
	"fmt"
	"testing"
)

func Test_hash_HashidsEncode(t *testing.T) {
	type fields struct {
		secret string
		length int
	}
	type args struct {
		params []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				secret: "12346",
				length: 6,
			},
			args: args{params: []int{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hash{
				secret: tt.fields.secret,
				length: tt.fields.length,
			}
			got, err := h.HashidsEncode(tt.args.params)
			fmt.Println(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashidsEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HashidsEncode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
