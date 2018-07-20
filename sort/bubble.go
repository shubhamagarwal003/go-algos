package sort

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func BubbleSort(arr []int) {
	l := len(arr)
	for i := 0; i < l; i++ {
		sw := false
		for j := 0; j < l-1; j++ {
			if arr[j] > arr[j+1] {
				sw = true
				swap(&arr[j], &arr[j+1])
			}
		}
		if !sw {
			break
		}
	}
}
