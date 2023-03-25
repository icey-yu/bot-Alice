package utils

func SliceDelete[T string | int | interface{}](list []T, index int) []T {
	if index < 0 || index >= len(list) {
		return list
	}
	// 特判，提高效率
	if index == 0 {
		return list[1:]
	} else if index == len(list)-1 {
		return list[:index]
	}
	return append(list[:index], list[index+1:]...)
}
