package sortalgo

import (
	"fmt"
	"log"
)

const STATS_KEY_COMPARE = "COMPARE"
const STATS_KEY_SWAP = "SWAP"

type Iterator[T any] interface {
	next() (T, error)
	nextN(n int) (T, error)
	prev() (T, error)
	swap(iter T)
	valueAssign(iter T)
	equal(iter T) bool
	valueGreaterThan(iter T) bool
	valueGreaterOrEqualThan(iter T) bool
	distanceFrom(iter T) int
}

type Source[Iter any, Self any] interface {
	new() Self
	begin() Iter
	end() Iter // pass-the-end
	len() int
	firstN(n int) (Self, error)
	lastN(n int) (Self, error)
	append(iter Iter)
	remove(iter Iter)
	copyFrom(s Self) error
	getStats() map[string]int
}

// Sort in asending order
func BubbleSort[Iter Iterator[Iter], SRC Source[Iter, SRC]](s SRC) error {
	end := s.end()
	begin := s.begin()

	for !end.equal(begin) {
		x := begin
		y, error := x.next()
		if error != nil {
			return error
		}

		for !y.equal(end) {
			if x.valueGreaterThan(y) {
				x.swap(y)
			}
			x, error = x.next()
			if error != nil {
				return error
			}
			y, error = y.next()
			if error != nil {
				return error
			}
		}

		end, _ = y.prev()
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
	return nil
}

// Sort in asending order
func QuickSort[Iter Iterator[Iter], SRC Source[Iter, SRC]](s SRC) error {
	if s.len() > 2 {
		left := s.begin()
		right, _ := s.end().prev()
		leftN := 0
		rightN := 0
		// Save pivot with new source
		var pivot SRC
		pivot = pivot.new()
		pivot.append(s.begin())

		for !left.equal(right) {
			for right.valueGreaterOrEqualThan(pivot.begin()) && !left.equal(right) {
				right, _ = right.prev()
				rightN++
			}
			left.valueAssign(right)

			for pivot.begin().valueGreaterOrEqualThan(left) && !left.equal(right) {
				left, _ = left.next()
				leftN++
			}
			right.valueAssign(left)
		}
		right.valueAssign(pivot.begin())
		if leftN == 0 || rightN == 0 {
			return nil
		}

		leftHalf, err := s.firstN(leftN)
		if err != nil {
			return err
		}
		rightHalf, err := s.lastN(rightN)
		if err != nil {
			return err
		}

		// log.Println("all ", s)
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
	return nil
}

// Sort in asending order
func shellSortInternal[Iter Iterator[Iter], SRC Source[Iter, SRC]](s SRC, gap int) error {

	curr := s.begin()
	currEnd, _ := curr.nextN(gap)
	for !curr.equal(currEnd) {

		first := curr
		for true {

			second, err := first.nextN(gap)
			if err != nil {
				break
			}

			if first.valueGreaterThan(second) {
				first.swap(second)
			}

			first = second

			second, err = second.nextN(gap)
			if err != nil {
				break
			}
		}

		curr, _ = curr.next()
	}

	if gap/2 >= 1 {
		return shellSortInternal[Iter, SRC](s, gap/2)
	}
	return nil
}
func ShellSort[Iter Iterator[Iter], SRC Source[Iter, SRC]](s SRC) error {
	return shellSortInternal[Iter, SRC](s, s.len()/2)
}

func heapify[Iter Iterator[Iter], SRC Source[Iter, SRC]](s SRC) {
	// Start from the last parent node
	workingNode, _ := s.begin().nextN(s.len()/2 - 1)
	var err error
	for true {
		log.Println("workingNode ", workingNode)
		pushDown(s, workingNode)
		workingNode, err = workingNode.prev()
		if err != nil {
			// All processed
			break
		}
	}
}
func pushDown[Iter Iterator[Iter], SRC Source[Iter, SRC]](s SRC, parent Iter) {
	left, errLeft := parent.nextN(parent.distanceFrom(s.begin()) + 1)
	right, errRight := parent.nextN(parent.distanceFrom(s.begin()) + 2)
	least := parent

	log.Println(parent)
	log.Println("errLeft ", errLeft, left)
	log.Println("errRight ", errRight, right)

	// Leaf node
	if errLeft != nil && errRight != nil {
		return
	}

	if errLeft == nil && least.valueGreaterThan(left) {
		least = left
	}
	if errRight == nil && least.valueGreaterThan(right) {
		least = right
	}

	if !least.equal(parent) {
		least.swap(parent)
		pushDown(s, least)
	}
}
func HeapSort[Iter Iterator[Iter], SRC Source[Iter, SRC]](s SRC) error {
	var result SRC
	result = result.new()

	heapify[Iter, SRC](s)
	for s.len() > 0 {
		result.append(s.begin())
		last, _ := s.end().prev()
		last.swap(s.begin())
		s.remove(last)
		pushDown[Iter, SRC](s, s.begin())
	}
	log.Println(result)
	for iter := result.begin(); !iter.equal(result.end()); iter, _ = iter.next() {
		s.append(iter)
	}
	log.Println(s)

	return nil
}

type IntArray struct {
	Data  []int
	Stats map[string]int
}

type IntArrayIter struct {
	array *IntArray
	idx   int
}

func (iter IntArrayIter) equal(otherIter IntArrayIter) bool {
	return iter.idx == otherIter.idx
}

func (iter IntArrayIter) valueGreaterThan(otherIter IntArrayIter) bool {
	// log.Println("comparing ", (iter.array.Data)[iter.idx], (otherIter.array.Data)[otherIter.idx])
	iter.array.Stats[STATS_KEY_COMPARE]++
	return (iter.array.Data)[iter.idx] > otherIter.array.Data[otherIter.idx]
}

func (iter IntArrayIter) valueGreaterOrEqualThan(otherIter IntArrayIter) bool {
	// log.Println("comparing ", (iter.array.Data)[iter.idx], (otherIter.array.Data)[otherIter.idx])
	iter.array.Stats[STATS_KEY_COMPARE]++
	return (iter.array.Data)[iter.idx] >= otherIter.array.Data[otherIter.idx]
}

func (iter IntArrayIter) valueAssign(otherIter IntArrayIter) {
	iter.array.Stats[STATS_KEY_SWAP]++
	(iter.array.Data)[iter.idx] = (otherIter.array.Data)[otherIter.idx]
}

func (iter IntArrayIter) next() (IntArrayIter, error) {
	if iter.idx >= len(iter.array.Data) {
		return IntArrayIter{iter.array, -1}, fmt.Errorf("pass the end")
	}
	return IntArrayIter{iter.array, iter.idx + 1}, nil
}

func (iter IntArrayIter) nextN(n int) (IntArrayIter, error) {
	if iter.idx+n >= len(iter.array.Data) {
		return IntArrayIter{iter.array, -1}, fmt.Errorf("pass the end")
	}
	return IntArrayIter{iter.array, iter.idx + n}, nil
}

func (iter IntArrayIter) distanceFrom(otherIter IntArrayIter) int {
	return iter.idx - otherIter.idx
}

func (iter IntArrayIter) prev() (IntArrayIter, error) {
	if iter.idx == 0 {
		return IntArrayIter{iter.array, -1}, fmt.Errorf("pass the begin")
	}
	return IntArrayIter{iter.array, iter.idx - 1}, nil
}

func (iter IntArrayIter) swap(otherIter IntArrayIter) {
	iter.array.Stats[STATS_KEY_SWAP]++
	(iter.array.Data)[iter.idx], (otherIter.array.Data)[otherIter.idx] = (otherIter.array.Data)[otherIter.idx], (iter.array.Data)[iter.idx]
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
	iter.idx = len((*this).Data)
	return iter
}

func (this *IntArray) len() int {
	return len((*this).Data)
}

func (this *IntArray) firstN(n int) (*IntArray, error) {
	var result IntArray
	var err error
	if n <= len((*this).Data) {
		result = IntArray{((*this).Data)[0:n], make(map[string]int, 0)}
		err = nil
	} else {
		result = IntArray{make([]int, 0), make(map[string]int, 0)}
		err = fmt.Errorf("Not enough elements, %d requested, total %d", n, len((*this).Data))
	}
	return &result, err
}

func (this *IntArray) lastN(n int) (*IntArray, error) {
	var result IntArray
	var err error
	if n <= len((*this).Data) {
		result = IntArray{((*this).Data)[len((*this).Data)-n:], make(map[string]int, 0)}
		err = nil
	} else {
		result = IntArray{make([]int, 0), make(map[string]int, 0)}
		err = fmt.Errorf("Not enough elements, %d requested, total %d", n, len((*this).Data))
	}
	return &result, err
}

func (this *IntArray) append(iter IntArrayIter) {
	// log.Println("this ", this)
	// log.Println("append ", (iter.array.Data)[iter.idx])
	(*this).Data = append((*this).Data, (iter.array.Data)[iter.idx])
	// log.Println("array ", (*this).Data)
}
func (this *IntArray) remove(iter IntArrayIter) {
	// log.Println("this ", this)
	// log.Println("append ", (iter.array.Data)[iter.idx])
	(*this).Data = append((*this).Data[0:iter.idx], (*this).Data[iter.idx+1:]...)
	// log.Println("array ", (*this).Data)
}

func (this *IntArray) copyFrom(from *IntArray) error {
	if len((*this).Data) != len((*from).Data) {
		return fmt.Errorf("from len %d is not equal to the toArray len %d", len((*from).Data), len((*this).Data))
	}
	copy((*this).Data, (*from).Data)
	return nil
}

func (this *IntArray) new() *IntArray {
	result := IntArray{make([]int, 0), make(map[string]int, 0)}
	return &result
}

func (this *IntArray) getStats() map[string]int {
	return (*this).Stats
}
