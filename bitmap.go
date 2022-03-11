package main

import (
	"errors"
)

const bitSize = 64

var (
	// ErrZeroCapacity 表示要创建的 bitmap 的容量小于等于 0，这样是没有意义的
	ErrZeroCapacity = errors.New("capacity should greater than 0")
	// ErrOutOfRange 表示存储的元素超出了 bitmap 的最大容量
	ErrOutOfRange = errors.New("out of range")
)

// BitMap go 目前范型还没有上线，我们定义 bitmap 只能固定一种类型，
// 这里用 uint64 就只是展示。添加了范型后会更加灵活
type BitMap interface {
	// Add 添加一个元素到 bitmap 中
	Add(uint64) error
	// Del 从 bitmap 中删除一个元素
	Del(uint64) error
	// Contain bitmap 中是否包含这个数字
	Contain(uint64) bool
}

type bitmap struct {
	store []uint64
}

// 计算 u 在 bitmap 中坐标
func (b *bitmap) coordinate(u uint64) (int, uint64) {
	return int(u / bitSize), 1 << (u % bitSize)
}

// Add 实现了 BitMap 接口的方法
func (b *bitmap) Add(u uint64) error {
	x, y := b.coordinate(u)
	if x > len(b.store) {
		return ErrOutOfRange
	}
	b.store[x] |= y
	return nil
}

// Del 实现了 BitMap 接口的方法
func (b *bitmap) Del(u uint64) error {
	x, y := b.coordinate(u)
	if x > len(b.store) {
		return ErrOutOfRange
	}
	b.store[x] &= ^y
	return nil
}

// Contain 实现了 BitMap 接口的方法
func (b *bitmap) Contain(u uint64) bool {
	x, y := b.coordinate(u)
	if x > len(b.store) {
		return false
	}
	return b.store[x]&y == y
}

// NewBitmap 创建一个 Bitmap，cap 表示 bitmap 的容量
func NewBitmap(cap int) (*bitmap, error) {
	if cap <= 0 {
		return nil, ErrZeroCapacity
	}
	return &bitmap{
		store: make([]uint64, cap/64+1), // +1 防止数字太小整除后结果为 0
	}, nil
}

// 类型检查
var _ BitMap = (*bitmap)(nil)
