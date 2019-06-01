package leetcode208

import "testing"

func TestTrie(t *testing.T) {
	obj := Constructor()
	obj.Insert("hello")
	if !obj.Search("hello") {
		t.Error("'hello' is 'hello'")
	}

	if obj.Search("abc") {
		t.Error("'abc' is not 'hello'")
	}

	if !obj.StartsWith("hell") {
		t.Error("'hello' is start with 'hell'")
	}
}
