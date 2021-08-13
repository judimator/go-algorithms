package quicksort

func Sort(arr []int, low, high int) {
	if low < high {
		i := swap(arr, low, high)
		Sort(arr, low, i-1)
		Sort(arr, i+1, high)
	}
}

func swap(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]

	return i
}
