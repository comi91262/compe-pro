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
