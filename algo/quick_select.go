package algo

func insertionSort(a []int) {
	for i := 0; i < len(a); i++ {
		j := i
		for (j > 0) && (a[j-1] > a[j]) {
			a[j-1], a[j] = a[j], a[j-1]
			j--
		}
	}
}

func median(a []int) int {
	n := len(a)
	insertionSort(a)

	if n%2 == 0 {
		return a[n/2-1]
	} else {
		return a[n/2]
	}
}

func selectPivot(a []int, left, right int) int {
	medians := []int{}
	for i := left; i < right; i += 5 {
		medians = append(medians, median(a[i:min(i+5, len(a))]))
	}
	if len(medians) == 0 {
		return a[left]
	}
	return median(medians)
}

// Hoare’s Partition Scheme
func partition(a []int, left, right, pivot int) int {
	i := left - 1
	j := right + 1

	for {
		i += 1
		for a[i] < pivot {
			i += 1
		}

		j -= 1
		for a[j] > pivot {
			j -= 1
		}

		if i >= j {
			return j
		}
		a[i], a[j] = a[j], a[i]
	}
}

func QuickSelect(a []int, left, right, kth int) int {
	var pivotIndex int
	for left != right {
		pivotIndex = selectPivot(a, left, right)
		pivotIndex = partition(a, left, right, pivotIndex)

		switch {
		case kth == pivotIndex:
			return a[kth]
		case kth < pivotIndex:
			right = pivotIndex - 1
		case kth > pivotIndex:
			left = pivotIndex + 1
		}
	}
	return a[left]
}
