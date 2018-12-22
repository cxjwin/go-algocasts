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
	root := ds.NewTree(1)

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
	fmt.Printf("recursive - tree is symmetric %v\n", isSymmetricTreeRecursive(root))
	fmt.Printf("iterative - tree is symmetric %v\n", isSymmetricTreeIterative(root))
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

func TestSameTree1(t *testing.T) {
	t1 := ds.NewTree(1)
	t1.Insert(2, true)
	t1.Insert(3, false)

	t2 := ds.NewTree(1)
	t2.Insert(2, true)
	t2.Insert(3, false)

	res := isSameTreeRecusive(t1, t2)
	if !res {
		t.Error("Recusive : t1 is same with t2")
	}
	res = isSameTreeIterative(t1, t2)
	if !res {
		t.Error("Iterative : t1 is same with t2")
	}
}
func TestSameTree2(t *testing.T) {
	t1 := ds.NewTree(1)
	t1.Insert(2, true)
	t1.Insert(1, false)

	t2 := ds.NewTree(1)
	t2.Insert(1, true)
	t2.Insert(2, false)

	res := isSameTreeRecusive(t1, t2)
	if res {
		t.Error("Recusive : t1 is not same with t2")
	}
	res = isSameTreeIterative(t1, t2)
	if res {
		t.Error("Iterative : t1 is not same with t2")
	}
}

func TestReverseLinkedListInterative(t *testing.T) {
	list := &ds.List{Head: nil, Tail: nil}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Insert(6)
	list.Desc()

	fmt.Println("========== line ==========")

	list = reverseListIterative(list)

	if list.Head.Value != 6 {
		t.Error("Reversed list's head is 6")
	}

	if list.Tail.Value != 1 {
		t.Error("Reversed list's tail is 1")
	}

	if list.Tail.Next != nil {
		t.Error("Reversed tail's next is nil")
	}

	list.Desc()
}
func TestReverseLinkedListRecursive(t *testing.T) {
	list := &ds.List{Head: nil, Tail: nil}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Insert(6)
	list.Desc()

	fmt.Println("========== line ==========")

	reverseListRecursive(list)

	if list.Head.Value != 6 {
		t.Error("Reversed list's head is 6")
	}

	if list.Tail.Value != 1 {
		t.Error("Reversed list's tail is 1")
	}

	if list.Tail.Next != nil {
		t.Error("Reversed tail's next is nil")
	}

	list.Desc()
}

func TestPow(t *testing.T) {
	n := pow(2, 4)
	if n != 16 {
		t.Error("pow(2, 4) is 16")
	}
	fmt.Println(n)

	n = pow(2, -4)
	if n != 0.0625 {
		t.Error("pow(2, -4) is 1/16")
	}
	fmt.Println(n)
}
func TestPowFast(t *testing.T) {
	n := powFast(2, 4)
	if n != 16 {
		t.Error("pow(2, 4) is 16")
	}
	fmt.Println(n)

	n = powFast(2, -4)
	if n != 0.0625 {
		t.Error("pow(2, -4) is 1/16")
	}
	fmt.Println(n)
}

func TestPermutations(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	res := permutations(nums)

	length := factorial(5)

	if res.Len() != length {
		t.Error("the number of results is 120(5!)")
	}

	e := res.Front()
	for i := 0; i < res.Len(); i++ {
		fmt.Println(e.Value)
		e = e.Next()
	}
}

func TestCountPalindromicSubstringsDynamicProgramming(t *testing.T) {
	str := "abc"
	count := countPalindromicSubstringsDynamicProgramming(str)
	if count != 3 {
		t.Error("'abc' count is 3")
	}

	str = "aaa"
	count = countPalindromicSubstringsDynamicProgramming(str)
	if count != 6 {
		t.Error("'aaa' count is 6")
	}
}
func TestCountPalindromicSubstringsExpand(t *testing.T) {
	str := "abc"
	count := countPalindromicSubstringsExpand(str)
	if count != 3 {
		t.Error("'abc' count is 3")
	}

	str = "aaa"
	count = countPalindromicSubstringsExpand(str)
	if count != 6 {
		t.Error("'aaa' count is 6")
	}
}

func TestIsPalindromeNumberUseString(t *testing.T) {
	num := 121
	res := isPalindromeNumberUseString(num)
	if !res {
		t.Error("121 is palindrome number")
	}

	num = -121
	res = isPalindromeNumberUseString(num)
	if res {
		t.Error("-121 is not palindrome number")
	}
}
func TestIsPalindromeNumber(t *testing.T) {
	num := 121
	res := isPalindromeNumber(num)
	if !res {
		t.Error("121 is palindrome number")
	}

	num = -121
	res = isPalindromeNumber(num)
	if res {
		t.Error("-121 is not palindrome number")
	}
}

func TestIsPalindromeLinkedListUsingStack(t *testing.T) {
	type testFunc func(list *ds.List) bool

	isPalindromeLinkedList := func(f testFunc, t *testing.T) {
		list := &ds.List{Head: nil, Tail: nil}
		list.Insert(1)
		list.Insert(2)
		list.Desc()

		res := f(list)
		if res {
			t.Error("'1 -> 2' is not palindrome linked list")
		}

		fmt.Println("========== line ==========")

		list = &ds.List{Head: nil, Tail: nil}
		list.Insert(1)
		list.Insert(2)
		list.Insert(2)
		list.Insert(1)
		list.Desc()

		res = f(list)

		if !res {
			t.Error("'1 -> 2 -> 2 -> 1' is not palindrome linked list")
		}

		fmt.Println("========== line ==========")

		list = &ds.List{Head: nil, Tail: nil}
		list.Insert(1)
		list.Insert(2)
		list.Insert(3)
		list.Insert(2)
		list.Insert(1)
		list.Desc()

		res = f(list)

		if !res {
			t.Error("'1 -> 2 -> 3 -> 2 -> 1' is not palindrome linked list")
		}
	}

	fmt.Println("========== stack ==========")
	isPalindromeLinkedList(isPalindromeLinkedListUsingStack, t)
	fmt.Println("========== reverse ==========")
	isPalindromeLinkedList(isPalindromeReverseLinkedList, t)
}
