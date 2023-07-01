package sliceM

func IntConcat(array1 []int, array2 []int) []int {
	return append(array1, array2...)
}
