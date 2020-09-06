package preoblem

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testData := []int{2, 3, 4, 54, 3, 345, 43, 34, 3453, 345, 35443, 3453, 435, 34534, 5345, 435, 3452, 23234, 23, 2, 324, 342, 234, 234}
	QuickSort(testData, 0, len(testData)-1)
	for _, data := range testData {
		fmt.Printf("%d ", data)
	}
}
