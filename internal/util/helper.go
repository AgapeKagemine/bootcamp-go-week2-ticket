package util

func IsEmpty[T any](list map[int]T) bool {
	return len(list) == 0
}

func IsExist[T any](list map[int]T, id int) bool {
	if _, exists := list[id]; exists {
		return true
	}
	return false
}
