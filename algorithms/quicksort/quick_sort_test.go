package quicksort

import (
	"reflect"
	"testing"
)

var tests = [][][]int{
	{
		[]int{1, 2, 5, 3, 1, 6, 4, 9, 6, 5},
		[]int{1, 1, 2, 3, 4, 5, 5, 6, 6, 9},
	},
}

func TestSort(t *testing.T) {
	for _, test := range tests {
		Sort(test[0], 0, len(test[0])-1)

		if !reflect.DeepEqual(test[1], test[0]) {
			t.Error("Invalid array order!")
		}
	}
}
