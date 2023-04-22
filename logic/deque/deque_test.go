package deque

import (
	"testing"
)

func TestDeq1(t *testing.T) {
	d := Deque[int]{make([]int, 0)}
	assert_eq(d.Len(), 0, t)
	d.Push_back(3, 4)

	assert_eq(d.Len(), 2, t)

	assert_eq(d.Index(0), 3, t)
	assert_eq(d.Index(1), 4, t)
	d.Push_front(1, 2)

	assert_eq(d.Index(0), 1, t)
	assert_eq(d.Index(1), 2, t)

	assert_eq(d.Len(), 4, t)

	assert_eq(d.Pop_front(), 1, t)
	assert_eq(d.Len(), 3, t)
	assert_eq(d.Pop_back(), 4, t)

	assert_eq(d.Pop_back(), 3, t)
	assert_eq(d.Pop_front(), 2, t)

	assert_eq(d.Len(), 0, t)

}

func TestDeq2(t *testing.T) {
	d := Deque[int]{[]int{1, 2, 3, 4, 5, 6}}
	assert_eq_slice(d.Pop_back_slice(3), []int{4, 5, 6}, t)
	assert_eq_slice(d.Pop_back_slice(0), []int{}, t)

	assert_eq_slice(d.Pop_back_slice(2), []int{2, 3}, t)

	assert_eq_slice(d.Pop_back_slice(2), []int{1}, t)

	assert_eq_slice(d.Pop_back_slice(10), []int{}, t)
	assert_eq(d.Len(), 0, t)

}

func assert_eq[T comparable](a, b T, t *testing.T) {
	if a != b {
		t.Error("Error")
	}
}

func assert_eq_slice[T comparable](a, b []T, t *testing.T) {
	if len(a) != len(b) {
		t.Error("Error")
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				t.Error("Error")
				return
			}
		}
	}
}
