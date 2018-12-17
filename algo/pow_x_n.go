package algo

func pow(x float64, n int) float64 {
	absN := n
	if n < 0 {
		absN = -n
	}
	res := float64(1)
	for i := 0; i < absN; i++ {
		res *= x
	}
	if n < 0 {
		return 1 / res
	}
	return res
}

func powFast(x float64, n int) float64 {
	absN := n
	if n < 0 {
		absN = -n
	}

	res := float64(1)
	for absN != 0 {
		if (absN & 1) == 1 {
			res *= x
		}
		x *= x
		absN >>= 1
	}

	if n < 0 {
		return 1 / res
	}

	return res
}
