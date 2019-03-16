package leetcode91

import "testing"

func TestNumDecodings(t *testing.T) {
	type testFunc func(string) int
	testBody := func(f testFunc, t *testing.T) {
		s := "12"
		res := f(s)
		if res != 2 {
			t.Error("'12' num decoding is 2")
		}
		s = "226"
		res = f(s)
		if res != 3 {
			t.Error("'226' num decoding is 3")
		}

		s = "0"
		res = f(s)
		if res != 0 {
			t.Error("'0' num decoding is 0")
		}

		s = "124"
		res = f(s)
		if res != 3 {
			t.Error("'124' num decoding is 3")
		}
	}

	testBody(numDecodings, t)
	testBody(numDecodingsDP, t)
	testBody(numDecodingsDPO1, t)
}
