package xy

func getSize[T []E | string, E any](m []T) XY {
	return XY{len(m[0]), len(m)}
}

func GetSizeString(m []string) XY {
	return getSize[string, any](m)
}

func GetSize[T any](v [][]T) XY {
	return getSize[[]T, T](v)
}

func XYFromArr(arr []int) XY {
	return XY{arr[0], arr[1]}
}
