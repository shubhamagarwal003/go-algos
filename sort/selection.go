package sort

func SelectionSort(arr []int) {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		minIndex := i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		swap(&arr[minIndex], &arr[i])
	}
}
