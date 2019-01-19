package numbersutils

// AvgInt calculates the average of a numbers
func AvgInt(nums ...int) int {
	if len(nums) == 0 {
		panic("Please provide at least one number")
	}
	var s = 0
	for _, x := range nums {
		s += x
	}
	return s / len(nums)
}

// AvgIntToFloat calculates the average of a numbers and returns a float32
func AvgIntToFloat(nums ...int) float32 {
	if len(nums) == 0 {
		panic("Please provide at least one number")
	}
	var s = 0
	for _, x := range nums {
		s += x
	}
	return float32(s) / float32(len(nums))
}
