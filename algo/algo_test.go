package algo

import (
	"fmt"
	"testing"
)
import "github.com/cxjwin/go-algocasts/datastructure"

func TestPascalTriangle(t *testing.T) {
	numRows := 5
	res := makePascalTriangle(numRows)
	for _, v := range res {
		fmt.Println(v)
	}
}

func TestSearchA2DMatrix(t *testing.T) {
	matrix := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	res := searchA2DMatrix(34, matrix)
	fmt.Println(res)
}

func TestSingleNumber(t *testing.T) {
	nums := []int{2, 2, 1}
	res := singleNumberWithXOR(nums)
	fmt.Println(res)

	nums = []int{4, 1, 2, 1, 2}
	res = singleNumberWithMap(nums)
	fmt.Println(res)
}

func TestSumOfTwoIntegers(t *testing.T) {
	a, b := 1, 2
	fmt.Printf("recusive - res : %v\n", getSumRecusive(a, b))

	a, b = 9, 11
	fmt.Printf("interative - res : %v\n", getSumInterative(a, b))
}

func TestSymmetricTree(t *testing.T) {
	root := ds.Tree{Left: nil, Value: 1, Right: nil}

	// left
	l1 := root.Insert(2, true)
	l1.Insert(3, true)
	l1.Insert(4, false)

	// right
	l2 := root.Insert(2, false)
	l2.Insert(4, true)
	l2.Insert(3, false)

	// print
	root.Desc()

	// check
	fmt.Printf("recursive - tree is symmetric %v\n", isSymmetricTreeRecursive(&root))
	fmt.Printf("iterative - tree is symmetric %v\n", isSymmetricTreeIterative(&root))
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	fir, sec := getTwoNumSumToGivenValueHashMap(nums, 9)
	fmt.Printf("find indexes : [%v, %v]\n", fir, sec)

	fir, sec = getTwoNumSumToGivenValueHashMap(nums, 14)
	fmt.Printf("find indexes : [%v, %v]\n", fir, sec)
}

func TestTwoSum2(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	fir, sec := getTwoNumSumToGivenValue(nums, 9)
	fmt.Printf("find indexes : [%v, %v]\n", fir, sec)
	if fir != 1 || sec != 2 {
		t.Error("res is [1, 2]")
	}
}

func TestIsPalindrome(t *testing.T) {
	s1 := "A man, a plan, a canal: Panama"
	fmt.Printf("s1 is palindrome? %t\n", isPalindrome(s1))
	if !isPalindrome(s1) {
		t.Error("s1 is palindrome")
	}
	s2 := "race a car"
	fmt.Printf("s2 is palindrome? %t\n", isPalindrome(s2))
	if isPalindrome(s2) {
		t.Error("s2 is not palindrome")
	}
}
