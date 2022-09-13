package sortalgo

import (
	"fmt"
)

type Iterator[T any] interface {
	next() (T, error)
	prev() (T, error)
	swap(iter T)
	equal(iter T) bool
	valueGreaterThan(iter T) bool
}

type Source[Iter any] interface {
	begin() Iter
	end() Iter // pass-the-end
}

func BubbleSort[Iter Iterator[Iter], SRC Source[Iter]](s SRC) error {
	end := s.end()
	x := s.begin()

	for !x.equal(end) {
		y, error := x.next()
		if error != nil {
			return error
		}

		for !y.equal(end) {
			if x.valueGreaterThan(y) {
				x.swap(y)
			}
			y, error = y.next()
			if error != nil {
				return error
			}
		}

		x, error = x.next()
		if error != nil {
			return error
		}
	}
	return nil
}

func InsertSort[Iter Iterator[Iter], SRC Source[Iter]](s SRC) error {
	end := s.end()
	x, error := s.begin().next()
	if error != nil {
		// single element
		return nil
	}

	for !x.equal(end) {
		curr := x
		prev, error := x.prev()
		if error != nil {
			continue
		}
		for ; error == nil; prev, error = prev.prev() {
			if prev.valueGreaterThan(curr) {
				curr.swap(prev)
			}
			curr = prev
		}
		x, error = x.next()
		if error != nil {
			return error
		}
	}
	return nil
}

type IntArray []int

type IntArrayIter struct {
	array *IntArray
	idx   int
}

func (iter IntArrayIter) equal(otherIter IntArrayIter) bool {
	return iter.idx == otherIter.idx
}

func (iter IntArrayIter) valueGreaterThan(otherIter IntArrayIter) bool {
	return (*iter.array)[iter.idx] > (*otherIter.array)[otherIter.idx]
}

func (iter IntArrayIter) next() (IntArrayIter, error) {
	if iter.idx >= len(*iter.array) {
		return IntArrayIter{iter.array, -1}, fmt.Errorf("pass the end")
	}
	return IntArrayIter{iter.array, iter.idx + 1}, nil
}
func (iter IntArrayIter) prev() (IntArrayIter, error) {
	if iter.idx == 0 {
		return IntArrayIter{iter.array, -1}, fmt.Errorf("pass the begin")
	}
	return IntArrayIter{iter.array, iter.idx - 1}, nil
}

func (iter IntArrayIter) swap(otherIter IntArrayIter) {
	(*iter.array)[iter.idx], (*otherIter.array)[otherIter.idx] = (*otherIter.array)[otherIter.idx], (*iter.array)[iter.idx]
}

func (array IntArray) begin() IntArrayIter {
	var iter IntArrayIter
	iter.array = &array
	iter.idx = 0
	return iter
}
func (array IntArray) end() IntArrayIter {
	var iter IntArrayIter
	iter.array = &array
	iter.idx = len(array)
	return iter
}
