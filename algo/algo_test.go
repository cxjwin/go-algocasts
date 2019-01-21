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

	if len(res) != length {
		t.Error("the number of results is 120(5!)")
	}

	// print result
	for _, v := range res {
		fmt.Println(v)
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
	type testFunc func(*ds.List) bool

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

func TestMissingNumber(t *testing.T) {
	type testFunc func([]int) int

	missingNumber := func(f testFunc, t *testing.T) {
		nums := []int{3, 0, 1}
		res := f(nums)
		if res != 2 {
			t.Error("missing number is 2")
		}

		nums = []int{9, 6, 4, 2, 3, 5, 7, 0, 1, 8}
		res = f(nums)
		if res != -1 {
			t.Error("no missing number")
		}

		nums = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
		res = f(nums)
		if res != 8 {
			t.Error("missing number is 8")
		}
	}

	fmt.Println("========== sort ==========")
	missingNumber(missingNumberUsingSort, t)
	fmt.Println("========== XOR ==========")
	missingNumber(missingNumberUsingXOR, t)
	fmt.Println("========== sum ==========")
	missingNumber(missingNumberUsingSum, t)
}

func TestMinmumDepthOfBindaryTree(t *testing.T) {
	type testFunc func(*ds.Tree) int

	innerFunc := func(f testFunc, t *testing.T) {
		root := ds.NewTree(1)
		root.Insert(9, true)
		node := root.Insert(20, false)
		node.Insert(15, true)
		node.Insert(7, false)
		root.Desc()

		res := f(root)
		if res != 2 {
			t.Error("min depth is 2")
		}
	}

	fmt.Println("========== recursive ==========")
	innerFunc(minDepthOfBinaryTreeRecursive, t)
	fmt.Println("========== iterative ==========")
	innerFunc(minDepthOfBinaryTreeIterative, t)
}

func TestMinStack(t *testing.T) {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	if s.GetMin() != -3 {
		t.Error("min number is -3")
	}
	s.Pop()
	if s.Top() != 0 {
		t.Error("top number is 0")
	}
	if s.GetMin() != -2 {
		t.Error("min number is -2")
	}
}
func TestMinStack2(t *testing.T) {
	s := Constructor2()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	if s.GetMin() != -3 {
		t.Error("min number is -3")
	}
	s.Pop()
	if s.Top() != 0 {
		t.Error("top number is 0")
	}
	if s.GetMin() != -2 {
		t.Error("min number is -2")
	}
}

func TestMergeTwoSortedList(t *testing.T) {
	l := &ds.ListNode{Value: 1, Next: nil}
	l1 := &ds.ListNode{Value: 3, Next: l}
	l.Next = l1
	l2 := &ds.ListNode{Value: 5, Next: l1}
	l1.Next = l2
	l2.Next = nil

	r := &ds.ListNode{Value: 2, Next: nil}
	r1 := &ds.ListNode{Value: 4, Next: r}
	r.Next = r1
	r2 := &ds.ListNode{Value: 6, Next: r1}
	r1.Next = r2
	r2.Next = nil

	res := mergeTwoSortedList(l, r)

	arr := []int{}
	for res != nil {
		arr = append(arr, res.Value.(int))
		res = res.Next
	}

	if len(arr) != 6 {
		t.Error("list len is 6")
	}

	for i, v := range arr {
		if v != i+1 {
			t.Error("merged list should be sorted")
		}
	}
}

func TestMergeSortedArray(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	mergeSortedArray(nums1, 3, nums2, 3)
	res := []int{1, 2, 2, 3, 5, 6}
	for i, v := range res {
		if v != nums1[i] {
			t.Error("wrong number")
		}
	}
}

func TestLURCache(t *testing.T) {
	cache := LRUCacheConstructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	v := cache.Get(1) // returns 1
	if v != 1 {
		t.Error("returns 1")
	}
	cache.Put(3, 3)  // evicts key 2
	v = cache.Get(2) // returns -1 (not found)
	if v != -1 {
		t.Error("returns -1 (not found)")
	}
	cache.Put(4, 4)  // evicts key 1
	v = cache.Get(1) // returns -1 (not found)
	if v != -1 {
		t.Error("returns -1 (not found)")
	}
	v = cache.Get(3) // returns 3
	if v != 3 {
		t.Error("returns 3")
	}
	v = cache.Get(4) // returns 4
	if v != 4 {
		t.Error("returns 4")
	}
}

func TestMedianOfTwoSortedArrays(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}

	res := medianOfTwoSortedArrays(nums1, nums2)
	if res != 2.0 {
		t.Error("median is 2.0")
	}

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	res = medianOfTwoSortedArrays(nums1, nums2)
	if res != 2.5 {
		t.Error("median is 2.5")
	}
}

func TestMaximunSubarray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArray(nums)
	if res != 6 {
		t.Error("max sum of subarray is 6")
	}
}

func TestMaximumDepthOfBinaryTree(t *testing.T) {
	type testFunc func(root *ds.Tree) int

	testBody := func(f testFunc, t *testing.T) {
		root := ds.NewTree(3)
		root.Insert(9, true)
		node := root.Insert(20, false)
		node.Insert(15, true)
		node.Insert(7, false)
		root.Desc()

		res := f(root)
		if res != 3 {
			t.Error("max depth is 3")
		}
	}

	fmt.Println("========== recursive ==========")
	testBody(maxDepthOfBinaryTree, t)
	fmt.Println("========== iterative ==========")
	testBody(maxDepthOfBinaryTreeIterative, t)
}

func TestMajorityElement(t *testing.T) {
	type testFunc func([]int) int

	testBody := func(f testFunc, t *testing.T) {
		nums := []int{3, 2, 3}
		res := f(nums)
		if res != 3 {
			t.Error("1 : result is 3")
		}

		nums = []int{2, 2, 1, 1, 1, 2, 2}
		res = f(nums)
		if res != 2 {
			t.Error("2 : result is 2")
		}
	}

	testBody(majorityElementWithMap, t)
	testBody(majorityElement, t)
}

func TestLengthOfLongestSubstring(t *testing.T) {
	type testFunc func(s string) int

	testBody := func(f testFunc, t *testing.T) {
		s := "abcabcbb"
		res := f(s)
		if res != 3 {
			t.Error("length is 3")
		}

		s = "bbbbb"
		res = f(s)
		if res != 1 {
			t.Error("length is 1")
		}

		s = "pwwkew"
		res = f(s)
		if res != 3 {
			t.Error("length is 3")
		}

		s = "abba"
		res = f(s)
		if res != 2 {
			t.Error("length is 2")
		}
	}

	testBody(lengthOfLongestSubstring2N, t)
	testBody(lengthOfLongestSubstring1N, t)
}

func TestLongestPalindromicSubstring(t *testing.T) {
	type testFunc func(s string) string

	testBody := func(f testFunc, t *testing.T) {
		s := "babad"
		sub := f(s)

		if (sub != "aba") && (sub != "bab") {
			t.Error("longest palindromic substring is 'aba' or 'bab'")
		}
	}

	testBody(longestPalindromicSubtringDP, t)
	testBody(longestPalindromicSubtringExpand, t)
}

func TestLinkedListCycle(t *testing.T) {
	// TODO
}

func TestDesignHashMap(t *testing.T) {
	m := ConstructorMyHashMap()
	m.Put(1, 1)
	m.Put(2, 2)
	if m.Get(1) != 1 {
		t.Error("value is 1")
	}
	if m.Get(3) != -1 {
		t.Error("3 is not in hash map")
	}
	m.Put(15578, 39252)
	if m.Get(15578) != 39252 {
		t.Error("when key is 15578, value is 39252")
	}
	m.Remove(2)
	if m.Get(2) != -1 {
		t.Error("2 is not in hash map")
	}
}

func TestIntersectionOfTwoLinkedList(t *testing.T) {

	type testFunc func(headA, headB *ds.ListNode) *ds.ListNode

	testBody := func(f testFunc, t *testing.T) {
		// 4->1->8->4->5
		// 5->0->1->8->4->5

		headA := &ds.ListNode{Value: 4, Next: nil}
		a1 := &ds.ListNode{Value: 1, Next: nil}
		headA.Next = a1

		headB := &ds.ListNode{Value: 5, Next: nil}
		b1 := &ds.ListNode{Value: 0, Next: nil}
		headB.Next = b1
		b2 := &ds.ListNode{Value: 1, Next: nil}
		b1.Next = b2

		ab1 := &ds.ListNode{Value: 8, Next: nil}
		a1.Next = ab1
		b2.Next = ab1

		ab2 := &ds.ListNode{Value: 4, Next: nil}
		ab1.Next = ab2

		ab3 := &ds.ListNode{Value: 5, Next: nil}
		ab2.Next = ab3

		node := f(headA, headB)
		if node != ab1 {
			t.Error("node is 8")
		}
	}

	testBody(intersectionOfTwoLinkedLists, t)
	testBody(intersectionOfTwoLinkedListsWithLength, t)
}
