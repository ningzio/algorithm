package bst

import (
	"reflect"
	"testing"
)

func TestBST_Delete(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *BST
		want  func() *BST
	}{
		{
			name: "only left child node",
			setup: func() *BST {
				bst := NewBST()
				bst.Insert(NewNode(2))
				bst.Insert(NewNode(1))
				bst.Delete(bst.root)
				return bst
			},
			want: func() *BST {
				bst := NewBST()
				bst.Insert(NewNode(1))
				return bst
			},
		},
		{
			name: "only right child node",
			setup: func() *BST {
				bst := NewBST()
				bst.Insert(NewNode(2))
				bst.Insert(NewNode(3))
				bst.Delete(bst.root)
				return bst
			},
			want: func() *BST {
				bst := NewBST()
				bst.Insert(NewNode(3))
				return bst
			},
		},
		{
			name: "no child node",
			setup: func() *BST {
				bst := NewBST()
				bst.Insert(NewNode(3))
				bst.Delete(bst.root)
				return bst
			},
			want: func() *BST {
				return NewBST()
			},
		},
		{
			name: "has left and right child nodes",
			setup: func() *BST {
				bst := NewBST()
				bst.Insert(NewNode(2))
				bst.Insert(NewNode(1))
				bst.Insert(NewNode(3))
				bst.Delete(bst.root)
				return bst
			},
			want: func() *BST {
				bst := NewBST()
				bst.Insert(NewNode(3))
				bst.Insert(NewNode(1))
				return bst
			},
		},
		{
			name: "has non-right successor node",
			setup: func() *BST {
				bst := NewBST()
				bst.Insert(NewNode(2))
				bst.Insert(NewNode(1))
				bst.Insert(NewNode(5))
				bst.Insert(NewNode(3))
				bst.Insert(NewNode(4))
				bst.Delete(bst.root)
				return bst
			},
			want: func() *BST {
				bst := NewBST()
				bst.Insert(NewNode(3))
				bst.Insert(NewNode(1))
				bst.Insert(NewNode(5))
				bst.Insert(NewNode(4))
				return bst
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, want := tt.setup(), tt.want(); !reflect.DeepEqual(got, want) {
				t.Errorf("Delete() = %v, want %v", got, want)
			}
		})
	}
}

func TestBST_Insert(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *BST
		want  func() *BST
	}{
		{
			name: "1",
			setup: func() *BST {
				bst := NewBST()
				bst.Insert(NewNode(2))
				bst.Insert(NewNode(1))
				bst.Insert(NewNode(3))
				return bst
			},
			want: func() *BST {
				a := NewNode(2)
				b := NewNode(1)
				c := NewNode(3)
				a.left = b
				a.right = c
				b.parent = a
				c.parent = a
				return &BST{root: a}
			},
		},
		{
			name: "2",
			setup: func() *BST {
				bst := NewBST()
				bst.Insert(NewNode(1))
				bst.Insert(NewNode(2))
				bst.Insert(NewNode(3))
				return bst
			},
			want: func() *BST {
				a := NewNode(1)
				b := NewNode(2)
				c := NewNode(3)
				a.right = b
				b.parent = a
				b.right = c
				c.parent = b
				return &BST{root: a}
			},
		},
		{
			name: "3",
			setup: func() *BST {
				bst := NewBST()
				bst.Insert(NewNode(3))
				bst.Insert(NewNode(2))
				bst.Insert(NewNode(1))
				return bst
			},
			want: func() *BST {
				a := NewNode(3)
				b := NewNode(2)
				c := NewNode(1)
				a.left = b
				b.parent = a
				b.left = c
				c.parent = b
				return &BST{root: a}
			},
		},
		{
			name: "4",
			setup: func() *BST {
				bst := NewBST()
				bst.Insert(NewNode(1))
				bst.Insert(NewNode(1))
				return bst
			},
			want: func() *BST {
				return &BST{root: NewNode(1)}
			},
		},
		{
			name: "5",
			setup: func() *BST {
				bst := NewBST()
				bst.Insert(NewNode(1))
				bst.Insert(NewNode(3))
				bst.Insert(NewNode(2))
				return bst
			},
			want: func() *BST {
				a := NewNode(1)
				b := NewNode(3)
				c := NewNode(2)
				a.right = b
				b.parent = a
				b.left = c
				c.parent = b
				return &BST{root: a}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, want := tt.setup(), tt.want(); !reflect.DeepEqual(got, want) {
				t.Errorf("Insert() = %v, want %v", got, want)
			}
		})
	}
}

func TestBST_Successor(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *Node
		want  *Node
	}{
		{
			name: "has right child node",
			setup: func() *Node {
				bst := NewBST()
				bst.Insert(NewNode(2))
				bst.Insert(NewNode(1))
				bst.Insert(NewNode(3))
				return bst.Successor(bst.root)
			},
			want: NewNode(3),
		},
		{
			name: "has multi right child nodes",
			setup: func() *Node {
				bst := NewBST()
				bst.Insert(NewNode(2))
				bst.Insert(NewNode(5))
				bst.Insert(NewNode(3))
				bst.Insert(NewNode(4))
				return bst.Successor(bst.root)
			},
			want: NewNode(3),
		},
		{
			name: "no right child nodes",
			setup: func() *Node {
				bst := NewBST()
				bst.Insert(NewNode(5))
				bst.Insert(NewNode(3))
				bst.Insert(NewNode(4))
				return bst.Successor(bst.root.left.right)
			},
			want: NewNode(5),
		},
		{
			name: "no right child nodes 1",
			setup: func() *Node {
				bst := NewBST()
				bst.Insert(NewNode(8))
				bst.Insert(NewNode(5))
				bst.Insert(NewNode(6))
				bst.Insert(NewNode(7))
				return bst.Successor(bst.root.left.right.right)
			},
			want: NewNode(8),
		},
		{
			name: "no successor",
			setup: func() *Node {
				bst := NewBST()
				bst.Insert(NewNode(1))
				return bst.Successor(bst.root)
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.setup()
			if got == nil {
				if got != tt.want {
					t.Errorf("Successor() = %v, want %v", got, tt.want)
				}
			} else {
				if got.val != tt.want.val {
					t.Errorf("Successor() = %v, want %v", got.val, tt.want.val)
				}
			}
		})
	}
}
func TestBST_Search(t *testing.T) {
	bst := NewBST()
	bst.Insert(NewNode(10))
	bst.Insert(NewNode(3))
	bst.Insert(NewNode(9))
	bst.Insert(NewNode(3))
	bst.Insert(NewNode(7))
	bst.Insert(NewNode(8))
	bst.Insert(NewNode(1))

	type fields struct {
		root *Node
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		{
			name: "",
			fields: fields{
				root: bst.root,
			},
			args: args{
				3,
			},
			want: NewNode(3),
		},
		{
			name: "",
			fields: fields{
				root: bst.root,
			},
			args: args{
				100,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BST{
				root: tt.fields.root,
			}
			got := b.Search(tt.args.val)
			if got == nil {
				if got != tt.want {
					t.Errorf("Search() = %v, want %v", got, tt.want)
				}
			} else {
				if got.val != tt.want.val {
					t.Errorf("Search() = %v, want %v", got.val, tt.want.val)
				}
			}
		})
	}
}

func TestBST_Predecessor(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *Node
		want  *Node
	}{
		{
			name: "has left child",
			setup: func() *Node {
				bst := NewBST()
				bst.Insert(NewNode(4))
				bst.Insert(NewNode(3))
				return bst.Predecessor(bst.root)
			},
			want: NewNode(3),
		},
		{
			name: "has multi left child nodes",
			setup: func() *Node {
				bst := NewBST()
				bst.Insert(NewNode(9))
				bst.Insert(NewNode(5))
				bst.Insert(NewNode(2))
				bst.Insert(NewNode(7))
				bst.Insert(NewNode(8))
				return bst.Predecessor(bst.root)
			},
			want: NewNode(8),
		},
		{
			name: "no left child node",
			setup: func() *Node {
				bst := NewBST()
				bst.Insert(NewNode(4))
				bst.Insert(NewNode(7))
				bst.Insert(NewNode(6))
				bst.Insert(NewNode(5))
				return bst.Predecessor(bst.root.right.left.left)
			},
			want: NewNode(4),
		},
		{
			name: "no predecessor",
			setup: func() *Node {
				bst := NewBST()
				bst.Insert(NewNode(1))
				return bst.Predecessor(bst.root)
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.setup()
			if got == nil {
				if got != tt.want {
					t.Errorf("Predecessor() = %v, want %v", got, tt.want)
				}
			} else {
				if got.val != tt.want.val {
					t.Errorf("Predecessor() = %v, want %v", got.val, tt.want.val)
				}
			}
		})
	}
}

func TestBST_Max(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *Node
		want  *Node
	}{
		{
			name: "",
			setup: func() *Node {
				bst := NewBST()
				bst.Insert(NewNode(2))
				bst.Insert(NewNode(5))
				bst.Insert(NewNode(3))
				bst.Insert(NewNode(4))
				return bst.root
			},
			want: NewNode(5),
		},
		{
			name: "",
			setup: func() *Node {
				bst := NewBST()
				bst.Insert(NewNode(5))
				return bst.root
			},
			want: NewNode(5),
		},
		{
			name: "",
			setup: func() *Node {
				return nil
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := tt.setup()
			b := &BST{
				root: root,
			}
			got := b.Max(root)
			if got == nil {
				if got != tt.want {
					t.Errorf("Max() = %v, want %v", got, tt.want)
				}
			} else {
				if got.val != tt.want.val {
					t.Errorf("Max() = %v, want %v", got.val, tt.want.val)
				}
			}
		})
	}
}

func TestBST_Min(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *Node
		want  *Node
	}{
		{
			name: "",
			setup: func() *Node {
				bst := NewBST()
				bst.Insert(NewNode(5))
				bst.Insert(NewNode(7))
				bst.Insert(NewNode(1))
				bst.Insert(NewNode(3))
				bst.Insert(NewNode(4))
				return bst.root
			},
			want: NewNode(1),
		},
		{
			name: "",
			setup: func() *Node {
				bst := NewBST()
				bst.Insert(NewNode(5))
				return bst.root
			},
			want: NewNode(5),
		},
		{
			name: "",
			setup: func() *Node {
				return nil
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := tt.setup()
			b := &BST{
				root: root,
			}
			got := b.Min(root)
			if got == nil {
				if got != tt.want {
					t.Errorf("Max() = %v, want %v", got, tt.want)
				}
			} else {
				if got.val != tt.want.val {
					t.Errorf("Max() = %v, want %v", got.val, tt.want.val)
				}
			}
		})
	}
}
