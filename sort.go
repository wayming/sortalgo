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

type Source[Iter any, Self any] interface {
	new() Self
	begin() Iter
	end() Iter // pass-the-end
	len() int
	firstN(n int) (Self, error)
	lastN(n int) (Self, error)
	append(iter Iter)
	copyFrom(s Self) error
}

// Sort in asending order
func BubbleSort[Iter Iterator[Iter], SRC Source[Iter, SRC]](s SRC) error {
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

// Sort in asending order
func InsertSort[Iter Iterator[Iter], SRC Source[Iter, SRC]](s SRC) error {
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

// Sort in asending order
func MergeSort[Iter Iterator[Iter], SRC Source[Iter, SRC]](s SRC) error {
	if s.len() > 2 {
		half := s.len() / 2

		var err error
		leftHalf, err := s.firstN(half)
		if err != nil {
			return err
		}

		rightHalf, err := s.lastN(s.len() - half)
		if err != nil {
			return err
		}
		// log.Println("left ", leftHalf)
		// log.Println("right ", rightHalf)
		if err = MergeSort[Iter, SRC](leftHalf); err != nil {
			return err
		}

		if err = MergeSort[Iter, SRC](rightHalf); err != nil {
			return err
		}
		// log.Println("ordered left ", leftHalf)
		// log.Println("ordered right ", rightHalf)
		left := leftHalf.begin()
		right := rightHalf.begin()
		var sorted SRC
		sorted = sorted.new()
		for !left.equal(leftHalf.end()) && !right.equal(rightHalf.end()) {
			if left.valueGreaterThan(right) {
				sorted.append(right)
				if right, err = right.next(); err != nil {
					return err
				}
			} else {
				sorted.append(left)
				if left, err = left.next(); err != nil {
					return err
				}
			}
			// log.Println("merged ", sorted)
		}

		if left.equal(leftHalf.end()) {
			for !right.equal(rightHalf.end()) {
				sorted.append(right)
				if right, err = right.next(); err != nil {
					return err
				}
			}
		}
		if right.equal(rightHalf.end()) {
			for !left.equal(leftHalf.end()) {
				sorted.append(left)
				if left, err = left.next(); err != nil {
					return err
				}
			}
		}

		s.copyFrom(sorted)
	} else if s.len() == 2 {
		left := s.begin()
		right := left
		var err error
		if right, err = left.next(); err != nil {
			return err
		}
		if left.valueGreaterThan(right) {
			left.swap(right)
		}
	}
	// log.Println("merged ", s)
	return nil
}

// Sort in asending order
func QuickSort[Iter Iterator[Iter], SRC Source[Iter, SRC]](s SRC) error {
	if s.len() > 2 {
		pivot, err := s.end().prev()
		if err != nil {
			return err
		}

		left := s.begin()
		right, _ := pivot.prev()
		leftN := 1
		rightN := 1
		for true {

			for pivot.valueGreaterThan(left) && !left.equal(right) {
				left, _ = left.next()
				leftN++
			}
			for right.valueGreaterThan(pivot) && !left.equal(right) {
				right, _ = right.prev()
				rightN++
			}

			if left.equal(right) {
				left.swap(pivot)
				break
			} else {
				left.swap(right)
			}
		}

		leftHalf, err := s.firstN(leftN)
		if err != nil {
			return err
		}
		rightHalf, err := s.lastN(rightN)
		if err != nil {
			return err
		}

		// log.Println("left ", leftHalf)
		// log.Println("right ", rightHalf)
		if err = QuickSort[Iter, SRC](leftHalf); err != nil {
			return err
		}

		if err = QuickSort[Iter, SRC](rightHalf); err != nil {
			return err
		}

	} else if s.len() == 2 {
		first := s.begin()
		second := first
		var err error
		if second, err = first.next(); err != nil {
			return err
		}
		if first.valueGreaterThan(second) {
			first.swap(second)
		}
	}
	// log.Println("merged ", s)
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
	// log.Println("comparing ", (*iter.array)[iter.idx], (*otherIter.array)[otherIter.idx])
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

func (this *IntArray) begin() IntArrayIter {
	var iter IntArrayIter
	iter.array = this
	iter.idx = 0
	return iter
}
func (this *IntArray) end() IntArrayIter {
	var iter IntArrayIter
	iter.array = this
	iter.idx = len(*this)
	return iter
}

func (this *IntArray) len() int {
	return len(*this)
}

func (this *IntArray) firstN(n int) (*IntArray, error) {
	var result IntArray
	var err error
	if n <= len(*this) {
		result = (*this)[0:n]
		err = nil
	} else {
		result = make([]int, 0)
		err = fmt.Errorf("Not enough elements, %d requested, total %d", n, len(*this))
	}
	return &result, err
}

func (this *IntArray) lastN(n int) (*IntArray, error) {
	var result IntArray
	var err error
	if n <= len(*this) {
		result = (*this)[len(*this)-n:]
		err = nil
	} else {
		result = make([]int, 0)
		err = fmt.Errorf("Not enough elements, %d requested, total %d", n, len(*this))
	}
	return &result, err
}

func (this *IntArray) append(iter IntArrayIter) {
	// log.Println("this ", this)
	// log.Println("append ", (*iter.array)[iter.idx])
	*this = append(*this, (*iter.array)[iter.idx])
	// log.Println("array ", *this)
}

func (array *IntArray) copyFrom(fromArray *IntArray) error {
	if len(*array) != len(*fromArray) {
		return fmt.Errorf("*fromArray len %d is not equal to the toArray len %d", len(*fromArray), len(*array))
	}
	copy(*array, *fromArray)
	return nil
}

func (array *IntArray) new() *IntArray {
	result := make(IntArray, 0)
	return &result
}
