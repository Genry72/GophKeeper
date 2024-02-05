package helper

// Ptr возвращает указатель на переданное значение
func Ptr[T any](s T) *T {
	return &s
}
