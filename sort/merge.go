package sort

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, left int, right int) {
	if left < right {
		mid := (left + right) / 2
		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
}

func merge(arr []int, left int, mid int, right int) {
	arr1 := []int{}
	arr2 := []int{}

	for l := left; l <= mid; l++ {
		arr1 = append(arr1, arr[l])
	}
	for l := mid + 1; l <= right; l++ {
		arr2 = append(arr2, arr[l])
	}
	i := 0
	j := 0
	k := left
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			arr[k] = arr1[i]
			i++
			k++
		} else {
			arr[k] = arr2[j]
			j++
			k++
		}
	}
	for i < len(arr1) {
		arr[k] = arr1[i]
		i++
		k++
	}

	for j < len(arr2) {
		arr[k] = arr2[j]
		j++
		k++
	}
}
