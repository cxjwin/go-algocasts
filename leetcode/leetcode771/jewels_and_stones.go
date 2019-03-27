package leetcode771

// https://leetcode.com/problems/jewels-and-stones/

func numJewelsInStones(J string, S string) int {
	if len(J) == 0 || len(S) == 0 {
		return 0
	}

	m := [256]bool{}
	for _, v := range J {
		m[v] = true
	}
	sum := 0
	for _, v := range S {
		if m[v] {
			sum++
		}
	}
	return sum
}
