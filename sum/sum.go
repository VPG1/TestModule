package sum

func Sum(vars []int) int {
	res := 0
	for _, v := range vars {
		res += v
	}

	return res
}
