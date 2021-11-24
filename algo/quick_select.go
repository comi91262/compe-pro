package algo

// [start, end]
func insertionSort(a []int, start, end int) {
	for i := start; i <= end; i++ {
		j := i
		for (j > start) && (a[j-1] > a[j]) {
			a[j-1], a[j] = a[j], a[j-1]
			j--
		}
	}
}
