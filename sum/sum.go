package sum

func Sum(vars ...int) int {
	if len(vars) == 0 {
		return 0
	}

	res := 0
	for _, v := range vars {
		res += v
	}

	return res
}
