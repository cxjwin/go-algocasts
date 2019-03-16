package leetcode49

import (
	"fmt"
	"sort"
	"strings"
)

// https://leetcode.com/problems/group-anagrams/

// MyKey - key type
type MyKey [26]int

func getKeyBySort(s string) string {
	sArr := []byte(s)
	sort.Slice(sArr, func(i, j int) bool {
		return sArr[i] < sArr[j]
	})
	return string(sArr)
}

func getKeyByCount(s string) string {
	c := [26]byte{}
	for _, v := range s {
		c[v-'a']++
	}
	var b strings.Builder
	for _, v := range c {
		fmt.Fprintf(&b, "%v", v)
	}
	return b.String()
}

func getKeyByCount2(s string) string {
	c := [26]int{}
	for _, v := range s {
		c[v-'a']++
	}
	var b strings.Builder
	for i, v := range c {
		if v != 0 {
			fmt.Fprintf(&b, "%s%v", string(i+'a'), v)
		}
	}
	return b.String()
}

func groupAnagrams(strs []string) [][]string {
	if strs == nil || len(strs) == 0 {
		return [][]string{}
	}

	m := make(map[string][]string)
	for _, str := range strs {
		// key := getKeyBySort(str)
		// key := getKeyByCount(str)
		key := getKeyByCount2(str)
		arr, ok := m[key]
		if ok {
			arr = append(arr, str)
			m[key] = arr
		} else {
			m[key] = []string{str}
		}
	}

	res := make([][]string, 0)
	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func getMyKeyByCount(s string) MyKey {
	c := MyKey{}
	for _, v := range s {
		c[v-'a']++
	}
	return c
}

func groupAnagramsMyKey(strs []string) [][]string {
	if strs == nil || len(strs) == 0 {
		return [][]string{}
	}

	m := make(map[MyKey][]string)
	for _, str := range strs {
		key := getMyKeyByCount(str)
		arr, ok := m[key]
		if ok {
			arr = append(arr, str)
			m[key] = arr
		} else {
			m[key] = []string{str}
		}
	}

	res := make([][]string, 0)
	for _, v := range m {
		res = append(res, v)
	}

	return res
}
