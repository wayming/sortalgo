package sortalgo_test

import (
	"math/rand"
	"reflect"
	"sortalgo"
	"testing"
	"time"
)

func GetRandomIntegers(size int) ([]int, []int) {
	src := rand.NewSource(time.Now().UnixNano())
	intRand := rand.New(src)
	source := make([]int, 0)
	for i := 0; i < size; i++ {
		source = append(source, intRand.Intn(10*size))
	}
	sorted := source
	sortalgo.BubbleSort[sortalgo.IntArrayIter, sortalgo.IntArray](sorted)

	return source, sorted
}

func TestBubbleSortSanity(t *testing.T) {
	nums := sortalgo.IntArray{10, 1, 20, 50, 5}
	expect := sortalgo.IntArray{1, 5, 10, 20, 50}

	sortalgo.BubbleSort[sortalgo.IntArrayIter, sortalgo.IntArray](nums)
	if !reflect.DeepEqual(nums, expect) {
		t.Fatal("Unexpected values after sort, expected ", expect, ", actual ", nums)
	}
}

func TestInsertSortSanity(t *testing.T) {
	nums := sortalgo.IntArray{10, 1, 20, 50, 5}
	expect := sortalgo.IntArray{1, 5, 10, 20, 50}

	sortalgo.InsertSort[sortalgo.IntArrayIter, sortalgo.IntArray](nums)
	if !reflect.DeepEqual(nums, expect) {
		t.Fatal("Unexpected values after sort, expected ", expect, ", actual ", nums)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	nums, _ := GetRandomIntegers(10000)

	for i := 0; i < b.N; i++ {
		sortalgo.BubbleSort[sortalgo.IntArrayIter, sortalgo.IntArray](nums)
	}
}

func BenchmarkInsertSort(b *testing.B) {
	nums, _ := GetRandomIntegers(10000)

	for i := 0; i < b.N; i++ {
		sortalgo.InsertSort[sortalgo.IntArrayIter, sortalgo.IntArray](nums)
	}
}

func TestInsertSortLarge(t *testing.T) {
	src, expected := GetRandomIntegers(10000)
	sortalgo.InsertSort[sortalgo.IntArrayIter, sortalgo.IntArray](src)
	if !reflect.DeepEqual(src, expected) {
		t.Fatal("Unexpected values after sort, expected ", expected, ", actual ", src)
	}
}
