package cha

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestNaEr(t *testing.T) {
	type args struct {
		s io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{"00.空", args{nil}, nil, true},
		{"01.全英文", args{strings.NewReader("e")}, nil, false},
		{"02.全中文", args{strings.NewReader("中")}, nil, false},
		{"03.中英中", args{strings.NewReader("中e中")}, []int{1, 4}, false},
		{"04.英中英", args{strings.NewReader("e中e")}, []int{1, 4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NaEr(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("NaEr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NaEr() = %v, want %v", got, tt.want)
			}
		})
	}
}
