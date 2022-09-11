package sortalgo_test

import (
	"fmt"
	"reflect"
	"sortalgo"
	"testing"
)

func TestBubbleSOrt(t *testing.T) {
	nums := sortalgo.IntArray{10, 1, 20, 50, 5}
	expect := sortalgo.IntArray{1, 5, 10, 20, 50}
	fmt.Println(nums)

	sortalgo.BubbleSort[sortalgo.IntArrayIter, sortalgo.IntArray](nums)
	fmt.Println(nums)
	if !reflect.DeepEqual(nums, expect) {
		t.Fatal("Unexpected values after sort, expected ", expect, ", actual ", nums)
	}
}
