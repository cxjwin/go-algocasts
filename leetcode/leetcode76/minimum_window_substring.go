package leetcode76

// https://leetcode.com/problems/minimum-window-substring/

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	required := make([]int, 256)

	n := len(s)
	requiredLen := len(t)

	for i := 0; i < requiredLen; i++ {
		required[t[i]]++
	}

	start := 0
	length := n + 1

	left, right := 0, 0

	for ; right < n; right++ {
		r := s[right]
		if required[r] > 0 {
			requiredLen--
		}
		required[r]--

		for requiredLen == 0 {
			if right-left+1 < length {
				start = left
				length = right - left + 1
			}

			l := s[left]
			required[l]++
			if required[l] > 0 {
				requiredLen++
			}
			left++
		}
	}

	if length == n+1 {
		return ""
	}

	return s[start : start+length]
}
