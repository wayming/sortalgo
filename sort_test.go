package sortalgo_test

import (
	"log"
	"math/rand"
	"reflect"
	"sortalgo"
	"testing"
	"time"
)

type SortFunc func(*sortalgo.IntArray) error

func GetRandomIntegers(size int) (sortalgo.IntArray, sortalgo.IntArray) {
	src := rand.NewSource(time.Now().UnixNano())
	intRand := rand.New(src)
	source := sortalgo.IntArray{make([]int, 0), make(map[string]int, 0)}
	for i := 0; i < size; i++ {
		source.Data = append(source.Data, intRand.Intn(10*size))
	}
	sorted := sortalgo.NewIntArrayFrom(&source)
	sortalgo.BubbleSort[sortalgo.IntArrayIter, *sortalgo.IntArray](sorted)
	sorted.Stats = make(map[string]int)
	return source, *sorted
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
	nums := sortalgo.IntArray{[]int{46, 2, 68, 0, 47, 57, 22}, make(map[string]int, 0)}
	expect := sortalgo.IntArray{[]int{0, 2, 22, 46, 47, 57, 68}, make(map[string]int, 0)}

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

func MergeStats(to map[string]int, from map[string]int) {
	for k, v := range from {
		to[k] += v
	}
}

func benchmarkSort(b *testing.B, sort SortFunc) {
	unsorted, sorted := GetRandomIntegers(10000)
	stats := make(map[string]int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		src := sortalgo.NewIntArrayFrom(&unsorted)
		sort(src)
		if !reflect.DeepEqual(src.Data, sorted.Data) {
			b.Fatal("Unexpected values after sort, expected ", sorted.Data, ", actual ", src.Data)
		}
		MergeStats(stats, src.Stats)
	}
	for k, v := range stats {
		stats[k] = v / b.N
	}
	log.Println(stats)
}

func BenchmarkBubbleSort(b *testing.B) {
	benchmarkSort(b, sortalgo.BubbleSort[sortalgo.IntArrayIter, *sortalgo.IntArray])
}

func BenchmarkInsertSort(b *testing.B) {
	benchmarkSort(b, sortalgo.InsertSort[sortalgo.IntArrayIter, *sortalgo.IntArray])
}

func BenchmarkMergeSort(b *testing.B) {
	benchmarkSort(b, sortalgo.MergeSort[sortalgo.IntArrayIter, *sortalgo.IntArray])
}

func BenchmarkQuickSort(b *testing.B) {
	benchmarkSort(b, sortalgo.QuickSort[sortalgo.IntArrayIter, *sortalgo.IntArray])
}

func BenchmarkShellSort(b *testing.B) {
	benchmarkSort(b, sortalgo.ShellSort[sortalgo.IntArrayIter, *sortalgo.IntArray])
}

func BenchmarkHeapSort(b *testing.B) {
	benchmarkSort(b, sortalgo.HeapSort[sortalgo.IntArrayIter, *sortalgo.IntArray])
}

func BenchmarkCountingSort(b *testing.B) {
	benchmarkSort(b, sortalgo.CountingSort[sortalgo.IntArrayIter, *sortalgo.IntArray])
}

func TestInsertSortLarge(t *testing.T) {
	// src, expected := GetRandomIntegers(5)
	src := sortalgo.IntArray{[]int{29, 40, 20, 41, 10}, make(map[string]int, 0)}
	expected := sortalgo.IntArray{[]int{10, 20, 29, 40, 41}, make(map[string]int, 0)}
	log.Println(src)
	log.Println(expected)

	sortalgo.CountingSort[sortalgo.IntArrayIter, *sortalgo.IntArray](&src)
	log.Println(src)

	if !reflect.DeepEqual(src.Data, expected.Data) {
		t.Fatal("Unexpected values after sort, expected ", expected.Data, ", actual ", src.Data)
	}
}
