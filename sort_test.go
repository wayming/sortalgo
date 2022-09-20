package sortalgo_test

import (
	"log"
	"math/rand"
	"reflect"
	"sortalgo"
	"testing"
	"time"
)

func GetRandomIntegers(size int) (sortalgo.IntArray, sortalgo.IntArray) {
	src := rand.NewSource(time.Now().UnixNano())
	intRand := rand.New(src)
	source := sortalgo.IntArray{make([]int, 0), make(map[string]int, 0)}
	for i := 0; i < size; i++ {
		source.Data = append(source.Data, intRand.Intn(10*size))
	}
	sorted := source
	sortalgo.BubbleSort[sortalgo.IntArrayIter, *sortalgo.IntArray](&sorted)

	return source, sorted
}

func TestBubbleSortSanity(t *testing.T) {
	nums := sortalgo.IntArray{[]int{10, 1, 20, 50, 5}, make(map[string]int, 0)}
	expect := sortalgo.IntArray{[]int{1, 5, 10, 20, 50}, make(map[string]int, 0)}

	sortalgo.BubbleSort[sortalgo.IntArrayIter, *sortalgo.IntArray](&nums)
	if !reflect.DeepEqual(nums.Data, expect.Data) {
		t.Fatal("Unexpected values after sort, expected ", expect, ", actual ", nums)
	}
}

func TestInsertSortSanity(t *testing.T) {
	nums := sortalgo.IntArray{[]int{10, 1, 20, 50, 5}, make(map[string]int, 0)}
	expect := sortalgo.IntArray{[]int{1, 5, 10, 20, 50}, make(map[string]int, 0)}

	sortalgo.InsertSort[sortalgo.IntArrayIter, *sortalgo.IntArray](&nums)
	if !reflect.DeepEqual(nums.Data, expect.Data) {
		t.Fatal("Unexpected values after sort, expected ", expect, ", actual ", nums)
	}
}

func TestMergeSortSanity(t *testing.T) {
	nums := sortalgo.IntArray{[]int{10, 1, 20, 50, 5}, make(map[string]int, 0)}
	expect := sortalgo.IntArray{[]int{1, 5, 10, 20, 50}, make(map[string]int, 0)}

	sortalgo.MergeSort[sortalgo.IntArrayIter, *sortalgo.IntArray](&nums)
	if !reflect.DeepEqual(nums.Data, expect.Data) {
		t.Fatal("Unexpected values after sort, expected ", expect, ", actual ", nums)
	}
}

func TestQuickSortSanity(t *testing.T) {
	nums := sortalgo.IntArray{[]int{10, 1, 20, 50, 5}, make(map[string]int, 0)}
	expect := sortalgo.IntArray{[]int{1, 5, 10, 20, 50}, make(map[string]int, 0)}

	sortalgo.QuickSort[sortalgo.IntArrayIter, *sortalgo.IntArray](&nums)
	if !reflect.DeepEqual(nums.Data, expect.Data) {
		t.Fatal("Unexpected values after sort, expected ", expect, ", actual ", nums)
	}
}

func TestShellSortSanity(t *testing.T) {
	nums := sortalgo.IntArray{[]int{10, 1, 20, 50, 5}, make(map[string]int, 0)}
	expect := sortalgo.IntArray{[]int{1, 5, 10, 20, 50}, make(map[string]int, 0)}

	sortalgo.ShellSort[sortalgo.IntArrayIter, *sortalgo.IntArray](&nums)
	if !reflect.DeepEqual(nums.Data, expect.Data) {
		t.Fatal("Unexpected values after sort, expected ", expect, ", actual ", nums)
	}
}

func TestHeapSortSanity(t *testing.T) {
	nums := sortalgo.IntArray{[]int{10, 1, 20, 50, 5}, make(map[string]int, 0)}
	expect := sortalgo.IntArray{[]int{1, 5, 10, 20, 50}, make(map[string]int, 0)}

	sortalgo.HeapSort[sortalgo.IntArrayIter, *sortalgo.IntArray](&nums)
	if !reflect.DeepEqual(nums.Data, expect.Data) {
		t.Fatal("Unexpected values after sort, expected ", expect, ", actual ", nums)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	nums, _ := GetRandomIntegers(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortalgo.BubbleSort[sortalgo.IntArrayIter, *sortalgo.IntArray](&nums)
		log.Println(nums.Stats)
	}
}

func BenchmarkInsertSort(b *testing.B) {
	nums, _ := GetRandomIntegers(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortalgo.InsertSort[sortalgo.IntArrayIter, *sortalgo.IntArray](&nums)
		log.Println(nums.Stats)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	nums, _ := GetRandomIntegers(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortalgo.MergeSort[sortalgo.IntArrayIter, *sortalgo.IntArray](&nums)
		log.Println(nums.Stats)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	nums, _ := GetRandomIntegers(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortalgo.QuickSort[sortalgo.IntArrayIter, *sortalgo.IntArray](&nums)
		log.Println(nums.Stats)
	}
}

func BenchmarkShellSort(b *testing.B) {
	nums, _ := GetRandomIntegers(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortalgo.ShellSort[sortalgo.IntArrayIter, *sortalgo.IntArray](&nums)
		log.Println(nums.Stats)
	}
}

func TestInsertSortLarge(t *testing.T) {
	src, expected := GetRandomIntegers(100)
	sortalgo.QuickSort[sortalgo.IntArrayIter, *sortalgo.IntArray](&src)
	if !reflect.DeepEqual(src, expected) {
		t.Fatal("Unexpected values after sort, expected ", expected, ", actual ", src)
	}
}
