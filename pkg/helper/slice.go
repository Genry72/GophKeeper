package helper

// SliceToAnySlice Преобразование произвольного слайса в слайс интерфейсов.
func SliceToAnySlice[T any](s []T) []any {
	result := make([]any, len(s))
	for i := range s {
		result[i] = s[i]
	}

	return result
}
