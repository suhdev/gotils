package numbersutils

func MaxInt(nums ...int) (int, int) {
	if len(nums) == 0 {
		panic("Not enough parameters provided, please specify at least one number")
	}
	max := nums[0]
	i := 0
	for idx, v := range nums {
		if v > max {
			max = v
			i = idx
		}
	}

	return max, i
}

func MaxBool(nums ...byte) (byte, int) {
	if len(nums) == 0 {
		panic("Not enough parameters provided, please specify at least one number")
	}
	max := nums[0]
	i := 0
	for idx, v := range nums {
		if v > max {
			max = v
			i = idx
		}
	}

	return max, i
}

func MaxInt64(nums ...int64) (int64, int) {
	if len(nums) == 0 {
		panic("Not enough parameters provided, please specify at least one number")
	}
	max := nums[0]
	i := 0
	for idx, v := range nums {
		if v > max {
			max = v
			i = idx
		}
	}

	return max, i
}

func MaxFloat32(nums ...float32) (float32, int) {
	if len(nums) == 0 {
		panic("Not enough parameters provided, please specify at least one number")
	}
	max := nums[0]
	i := 0
	for idx, v := range nums {
		if v > max {
			max = v
			i = idx
		}
	}

	return max, i
}

func MaxFloat64(nums ...float64) (float64, int) {
	if len(nums) == 0 {
		panic("Not enough parameters provided, please specify at least one number")
	}
	max := nums[0]
	i := 0
	for idx, v := range nums {
		if v > max {
			max = v
			i = idx
		}
	}

	return max, i
}

func MinInt(nums ...int) (int, int) {
	if len(nums) == 0 {
		panic("Not enough parameters provided, please specify at least one number")
	}
	min := nums[0]
	i := 0
	for idx, v := range nums {
		if v < min {
			min = v
			i = idx
		}
	}

	return min, i
}

func MinInt64(nums ...int64) (int64, int) {
	if len(nums) == 0 {
		panic("Not enough parameters provided, please specify at least one number")
	}
	min := nums[0]
	i := 0
	for idx, v := range nums {
		if v < min {
			min = v
			i = idx
		}
	}

	return min, i
}

func MinFloat64(nums ...float64) (float64, int) {
	if len(nums) == 0 {
		panic("Not enough parameters provided, please specify at least one number")
	}
	min := nums[0]
	i := 0
	for idx, v := range nums {
		if v < min {
			min = v
			i = idx
		}
	}

	return min, i
}

func MinFloat32(nums ...float32) (float32, int) {
	if len(nums) == 0 {
		panic("Not enough parameters provided, please specify at least one number")
	}
	min := nums[0]
	i := 0
	for idx, v := range nums {
		if v < min {
			min = v
			i = idx
		}
	}

	return min, i
}

func MinByte(nums ...byte) (byte, int) {
	if len(nums) == 0 {
		panic("Not enough parameters provided, please specify at least one number")
	}
	min := nums[0]
	i := 0
	for idx, v := range nums {
		if v < min {
			min = v
			i = idx
		}
	}

	return min, i
}
