package sort

func InsertionSort(arr []int) {
	l := len(arr)
	for i := 1; i < l; i++ {
		j := i - 1
		key := arr[i]
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}
