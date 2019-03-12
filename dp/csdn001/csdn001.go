package csdn001

// https://blog.csdn.net/u013309870/article/details/75193592

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func _fibonacciMemo(n int, memo []int) int {
	if n == 0 || n == 1 {
		memo[n] = n
		return memo[n]
	}

	if memo[n] != -1 {
		return memo[n]
	}

	memo[n] = _fibonacciMemo(n-1, memo) + _fibonacciMemo(n-2, memo)
	return memo[n]
}

func fibonacciMemo(n int) int {
	memo := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = -1
	}
	return _fibonacciMemo(n, memo)
}

func fibonacciDP(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	pre2 := 0
	pre1 := 1
	cur := 1
	for i := 2; i < n+1; i++ {
		cur = pre1 + pre2
		pre2 = pre1
		pre1 = cur
	}
	return cur
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func cut(p []int, n int) int {
	if n == 0 {
		return 0
	}

	q := -1
	for i := 1; i <= n; i++ {
		q = max(q, p[i]+cut(p, n-i))
	}
	return q
}

func _cutMemo(p []int, n int, memo []int) int {
	if memo[n] != -1 {
		return memo[n]
	}

	if n == 0 {
		memo[n] = 0
		return 0
	}

	q := -1
	for i := 1; i <= n; i++ {
		q = max(q, p[i]+_cutMemo(p, n-i, memo))
	}
	memo[n] = q
	return memo[n]
}

func cutMemo(p []int, n int) int {
	memo := make([]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = -1
	}
	return _cutMemo(p, n, memo)
}

func cutDP(p []int, n int) int {
	if n == 0 {
		return 0
	}

	d := make([]int, len(p)+1)
	d[0] = 0
	for j := 1; j <= n; j++ {
		q := -1
		for i := 1; i <= j; i++ {
			q = max(q, p[i]+d[j-i])
		}
		d[j] = q
	}
	return d[n]
}
