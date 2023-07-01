package sliceM

func IntShift(nums []int) []int {
	var result []int
	return append(result, nums[1:]...)
}
func StringShift(strings []string) []string {
	var result []string
	return append(result, strings[1:]...)
}

func IntPop(nums []int) []int {
	var result []int
	return append(result, nums[:len(nums)-1]...)
}
func StringPop(strings []string) []string {
	var result []string
	return append(result, strings[:len(strings)-1]...)
}
