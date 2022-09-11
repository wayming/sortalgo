package sortalgo

import "fmt"

type Iterator[T any] interface {
	next() T
	swap(iter T)
	equal(iter T) bool
	valueGreaterThan(iter T) bool
}

type Source[T2 any] interface {
	begin() T2
	end() T2
}

func BubbleSort[T3 Iterator[T3], T4 Source[T3]](s T4) {
	end := s.end()
	for x := s.begin(); !x.equal(end); x = x.next() {
		y := x
		fmt.Println("x: ", x)
		for y = y.next(); !y.equal(end); y = y.next() {
			fmt.Println("y: ", y)
			if x.valueGreaterThan(y) {
				x.swap(y)
			}
		}
	}

}

type IntArray []int

type IntArrayIter struct {
	array *IntArray
	idx   int
}

func (iter IntArrayIter) equal(otherIter IntArrayIter) bool {
	fmt.Println("equal ", iter.idx, " with ", otherIter.idx)

	return iter.idx == otherIter.idx
}

func (iter IntArrayIter) valueGreaterThan(otherIter IntArrayIter) bool {
	fmt.Println("valueGreaterThan ", (*iter.array)[iter.idx], " with ", (*otherIter.array)[otherIter.idx])
	return (*iter.array)[iter.idx] > (*otherIter.array)[otherIter.idx]
}

func (iter IntArrayIter) next() IntArrayIter {
	fmt.Println("next ", iter.idx)
	return IntArrayIter{iter.array, iter.idx + 1}
}

func (iter IntArrayIter) swap(otherIter IntArrayIter) {
	fmt.Println("swap ", iter.idx, " with ", otherIter.idx)
	(*iter.array)[iter.idx], (*otherIter.array)[otherIter.idx] = (*otherIter.array)[otherIter.idx], (*iter.array)[iter.idx]
}

func (array IntArray) begin() IntArrayIter {
	fmt.Println("begin")
	var iter IntArrayIter
	iter.array = &array
	iter.idx = 0
	return iter
}
func (array IntArray) end() IntArrayIter {
	fmt.Println("end")
	var iter IntArrayIter
	iter.array = &array
	iter.idx = len(array)
	return iter
}
