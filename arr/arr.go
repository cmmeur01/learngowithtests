package arr

func ArrSum(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}

	return sum
}

func SumArrs(slices ...[]int) []int {
	results := []int{}

	for _, slice := range slices {
		results = append(results, ArrSum(slice))
	}

	return results
}

func SumArrTails(slices ...[]int) []int {
	results := []int{}

	for _, slice := range slices {
		if len(slice) == 0 {
			results = append(results, 0)
		} else {
			results = append(results, ArrSum(slice[1:]))
		}
	}

	return results
}
