package main

import (
	"reflect"
	"testing"
)

func TestNewBitmap(t *testing.T) {
	type args struct {
		cap int
	}
	tests := []struct {
		name    string
		args    args
		want    *bitmap
		wantErr bool
	}{
		{
			name: "negative capacity",
			args: args{
				cap: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				cap: 1,
			},
			want: &bitmap{
				store: make([]uint64, 1),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBitmap(tt.args.cap)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBitmap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBitmap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_BitMap(t *testing.T) {
	bm, _ := NewBitmap(10)
	if err := bm.Add(10); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !bm.Contain(10) {
		t.Errorf("bitmap should contain number 10")
	}

	if err := bm.Del(10); err != nil {
		t.Errorf("unexpect error: %s", err)
	}
	if bm.Contain(10) {
		t.Errorf("bitmap should not contain number 10")
	}
}
