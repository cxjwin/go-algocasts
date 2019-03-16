package leetcode49

import (
	"fmt"
	"testing"
)

func TestGetKeyBySort(t *testing.T) {
	s := "eat"
	res := getKeyBySort(s)
	if res != "aet" {
		t.Error("eat -> aet")
	}
}

func TestGetKeyByCount(t *testing.T) {
	s := "eat"
	res := getKeyByCount(s)
	if res != "10001000000000000001000000" {
		t.Error("eat -> 10001000000000000001000000")
	}
	s1 := "fee"
	s2 := "egg"
	fmt.Println(getKeyByCount(s1))
	fmt.Println(getKeyByCount(s2))
}

func TestGetKeyByCount2(t *testing.T) {
	s1 := "fee"
	s2 := "egg"
	fmt.Println(getKeyByCount2(s1))
	fmt.Println(getKeyByCount2(s2))
}

func TestGroupAnagrams(t *testing.T) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	output := groupAnagrams(input)
	if len(output) != 3 {
		t.Error("3 groups")
	}
}

func TestGroupAnagramsMyKey(t *testing.T) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	output := groupAnagramsMyKey(input)
	if len(output) != 3 {
		t.Error("3 groups")
	}
}
