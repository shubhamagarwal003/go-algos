package sort

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		pivot := getPivot(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func getPivot(arr []int, low int, high int) int {
	pivot := high
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= arr[pivot] {
			i++
			swap(&arr[i], &arr[j])
		}
	}
	swap(&arr[i+1], &arr[high])
	return i + 1
}
