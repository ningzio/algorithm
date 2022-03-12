package main

import (
	"reflect"
	"testing"
)

func TestNode_leftMaxHeight(t *testing.T) {
	type fields struct {
		val   int
		left  *Node
		right *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "ok",
			fields: fields{
				val:  5,
				left: NewNode(1).Add(2).Add(3),
			},
			want: 2,
		},
		{
			name:   "ok - 1",
			fields: fields{},
			want:   0,
		},
		{
			name: "ok - 2",
			fields: fields{
				left: NewNode(1),
			},
			want: 1,
		},
		{
			name: "ok - 3",
			fields: fields{
				left: NewNode(1).Add(2).Add(3).Add(4),
			},
			want: 3,
		},
		{
			name: "ok - 4",
			fields: fields{
				val: 8,
				left: &Node{
					val:  5,
					left: NewNode(3),
					right: &Node{
						val:  7,
						left: NewNode(6),
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				val:   tt.fields.val,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := n.leftMaxHeight(); got != tt.want {
				t.Errorf("leftMaxHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_maxHeight(t *testing.T) {
	type fields struct {
		val   int
		left  *Node
		right *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "case - 1",
			fields: fields{
				val:   8,
				right: NewNode(10),
				left: &Node{
					val:  5,
					left: NewNode(3),
					right: &Node{
						val:  7,
						left: NewNode(6),
					},
				},
			},
			want: 3,
		},
		{
			name: "case - 2",
			fields: fields{
				val: 5,
				left: &Node{
					val:   2,
					right: NewNode(3),
				},
				right: &Node{
					val: 8,
					left: &Node{
						val:   6,
						right: NewNode(7),
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				val:   tt.fields.val,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := n.maxHeight(); got != tt.want {
				t.Errorf("maxHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_rightMaxHeight(t *testing.T) {
	type fields struct {
		val   int
		left  *Node
		right *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "ok",
			fields: fields{right: NewNode(1).Add(2).Add(3)},
			want:   2,
		},
		{
			name:   "ok - 1",
			fields: fields{},
			want:   0,
		},
		{
			name:   "ok - 2",
			fields: fields{right: NewNode(1)},
			want:   1,
		},
		{
			name:   "",
			fields: fields{right: NewNode(1).Add(2).Add(3).Add(4)},
			want:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				val:   tt.fields.val,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := n.rightMaxHeight(); got != tt.want {
				t.Errorf("rightMaxHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_autoRotate(t *testing.T) {
	type fields struct {
		val   int
		left  *Node
		right *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		{
			name: "right rotate",
			fields: fields{
				val: 4,
				left: &Node{
					val:   3,
					left:  NewNode(1),
					right: NewNode(2),
				},
				right: nil,
			},
			want: &Node{
				val:   3,
				left:  NewNode(1),
				right: NewNode(4).Add(2),
			},
		},
		{
			name: "left rotate",
			fields: fields{
				val:  2,
				left: nil,
				right: &Node{
					val:   5,
					right: NewNode(6),
				},
			},
			want: &Node{
				val:   5,
				left:  NewNode(2),
				right: NewNode(6),
			},
		},
		{
			name: "no rotate performed",
			fields: fields{
				val:   5,
				left:  NewNode(3),
				right: NewNode(6),
			},
			want: &Node{
				val:   5,
				left:  NewNode(3),
				right: NewNode(6),
			},
		},
		{
			name: "left-right rotate",
			fields: fields{
				val: 5,
				right: &Node{
					val:  8,
					left: NewNode(7),
				},
			},
			want: &Node{
				val:   7,
				left:  NewNode(5),
				right: NewNode(8),
			},
		},
		{
			name: "right-left rotate",
			fields: fields{
				val:   8,
				right: NewNode(10),
				left: &Node{
					val:  5,
					left: NewNode(3),
					right: &Node{
						val:  7,
						left: NewNode(6),
					},
				},
			},
			want: &Node{
				val: 7,
				left: &Node{
					val:   5,
					left:  NewNode(3),
					right: NewNode(6),
				},
				right: &Node{
					val:   8,
					right: NewNode(10),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				val:   tt.fields.val,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := n.autoRotate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("autoRotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
