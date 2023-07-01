package sliceM

func IntUnshift(nums []int, element int) []int {
	result := []int{
		element,
	}
	return append(result, nums...)
}
func StringUnshift(strings []string, element string) []string {
	result := []string{
		element,
	}
	return append(result, strings...)
}
