package interview

//QuickSort 快速排序
func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	tmp := arr[left]

	i, j := left, right
	for i != j {
		for arr[j] >= tmp && i < j {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}

		for arr[i] < tmp && i < j {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = tmp
	QuickSort(arr, left, i-1)
	QuickSort(arr, i+1, right)
}
