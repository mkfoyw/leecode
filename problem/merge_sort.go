package problem

func MergeSort(a []int) {
	mergeSort(a, 0, len(a)-1)
}

func mergeSort(a []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSort(a, start, mid)
	mergeSort(a, mid+1, end)

}

func merge(a []int, start, mid, end int) {
	tmparr := make([]int, end-start+1)

	i, j := start, mid+1

	cur := 0
	for i <= mid && j <= end {
		if a[i] < a[j] {
			tmparr[cur] = a[i]
			i++
		} else {
			tmparr[cur] = a[j]
			j++
		}
		cur++
	}

	for i <= mid {
		tmparr[cur] = a[i]
		i++
		cur++
	}

	for j <= end {
		tmparr[cur] = a[j]
		j++
		cur++
	}
	copy(a[start:end+1], tmparr)
}
