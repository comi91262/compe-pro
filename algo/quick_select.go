package algo

func insertionSort(a []int, left, right int) {
	for i := left; i <= right; i++ {
		j := i
		for (j > left) && (a[j-1] > a[j]) {
			a[j-1], a[j] = a[j], a[j-1]
			j--
		}
	}
}

func median(a []int, left, right int) int {
	insertionSort(a, left, right)
	return a[left+(right-left)/2]
}

func selectPivot(a []int, left, right int) int {
	if right-left < 5 {
		return median(a, left, right)
	}

	medians := []int{}
	for i := left; i+4 <= right; i += 5 {
		medians = append(medians, median(a, i, i+4))
	}
	return innerSelect(medians, 0, len(medians)-1, (len(medians)-1)/2)

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
func innerSelect(a []int, left, right, kth int) int {
	for left != right {
		pivot := selectPivot(a, left, right)
		pivotIndex := partition(a, left, right, pivot)

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

func QuickSelect(a []int, left, right, kth int) int {
	return innerSelect(a, left, right, kth)
}
