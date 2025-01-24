package utils

func Index[T comparable](slice []T, target T) int {
	for i, item := range slice {
		if item == target {
			return i
		}
	}
	return -1
}

func Swap[T any](slice []T, first int, second int) {
	slice[first], slice[second] = slice[second], slice[first]
}

func Insert[T any](slice *[]T, idx int, value T) bool {
	if idx < 0 || idx > len(*slice) {
		return false
	}
	*slice = append((*slice)[:idx], append([]T{value}, (*slice)[idx:]...)...)
	return true
}

func RemoveIndex[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}
	ret := make([]T, 0, len(slice)-1)
	ret = append(ret, slice[:index]...)
	return append(ret, slice[index+1:]...)
}

func RemoveItem[T any](slice *[]T, item T, equals func(a, b T) bool) {
	idx := -1
	for i, v := range *slice {
		if equals(v, item) {
			idx = i
			break
		}
	}
	if idx != -1 {
		*slice = RemoveIndex(*slice, idx)
	}
}

func PopAndRequeue[T any](slice *[]T) T {
	val := (*slice)[0]
	*slice = append((*slice)[1:], val)

	return val
}

func PopQueue[T any](slice *[]T) T {
	val := (*slice)[0]
	*slice = (*slice)[1:]

	return val
}

func PushQueue[T any](slice *[]T, val T) {
	*slice = append((*slice), val)
}
