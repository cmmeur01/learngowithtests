package arr

import (
	"reflect"
	"testing"
)

func TestArrSum(t *testing.T) {
	nums := []int{1, 2, 10, 5, 2, 3}
	got := ArrSum(nums)
	want := 23
	if got != want {
		t.Errorf("Expected %d got %d, given %v", want, got, nums)
	}
}

func TestSumArrs(t *testing.T) {
	nums1 := []int{1, 2, 3}
	nums2 := []int{4, 5, 6}
	nums3 := []int{7}
	got := SumArrs(nums1, nums2, nums3)
	want := []int{6, 15, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %d, got %d, given %v and %v", want, got, nums1, nums2)
	}
}

func TestSumArrTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("handle empty slices", func(t *testing.T) {
		got := SumArrTails([]int{}, []int{4, 5, 6})
		want := []int{0, 11}

		checkSums(t, got, want)
	})

	t.Run("add the tails of some slices", func(t *testing.T) {
		got := SumArrTails([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}

		checkSums(t, got, want)
	})
}
