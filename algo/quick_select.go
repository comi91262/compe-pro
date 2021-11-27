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

func partition(a []int, left, right, pivot int) int {
	i := left - 1
	j := right + 1

	for {
		i, j = i+1, j-1
		for a[i] < pivot {
			i += 1
		}
		for a[j] > pivot {
			j -= 1
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		} else {
			return j
		}
	}
}

func selectPivot(a []int, left, right int) int {
	if right-left < 5 {
		return median(a, left, right)
	}

	for i := left; i+4 <= right; i += 5 {
		insertionSort(a, i, i+4)
		a[i+2], a[left+(i-left)/5] = a[left+(i-left)/5], a[i+2]
	}

	n := right - left + 1
	return innerSelect(a, left, left+n/5-1, left+n/10-1)
}

func innerSelect(a []int, left, right, kth int) int {
	if left == right {
		return a[left]
	}

	pivot := selectPivot(a, left, right)
	pivotIndex := partition(a, left, right, pivot)
	if kth <= pivotIndex {
		return innerSelect(a, left, pivotIndex, kth)
	} else {
		return innerSelect(a, pivotIndex+1, right, kth)
	}
}

func QuickSelect(a []int, left, right, kth int) int {
	return innerSelect(a, left, right, kth)
}
