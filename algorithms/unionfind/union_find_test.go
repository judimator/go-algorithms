package unionfind

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

var err = errors.New("expected set %v is not equal to actual set %v")

func TestFind(t *testing.T) {
	set := New(10)
	if 2 != set.Find(2) {
		t.Error("Invalid root for item \"2\"")
	}
}

func TestUnion(t *testing.T) {
	set := New(10)

	set.Union(4, 5)
	assertEqual(t, []int{0, 1, 2, 3, 5, 5, 6, 7, 8, 9}, set.Items())

	set.Union(3, 4)
	assertEqual(t, []int{0, 1, 2, 5, 5, 5, 6, 7, 8, 9}, set.Items())

	set.Union(1, 2)
	assertEqual(t, []int{0, 2, 2, 5, 5, 5, 6, 7, 8, 9}, set.Items())

	set.Union(1, 9)
	assertEqual(t, []int{0, 2, 9, 5, 5, 5, 6, 7, 8, 9}, set.Items())
}

func assertEqual(t *testing.T, expected, actual []int) {
	if !reflect.DeepEqual(expected, actual) {
		t.Error(fmt.Sprintf(err.Error(), expected, actual))
	}
}
