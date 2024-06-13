/*
70. Climbing Stairs
https://leetcode.com/problems/climbing-stairs/

You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

*/

package climbingstairs

// Dynamic Programming Tabulation with space optimisation
// Time Complexity: O(N)
// Space Complexity: O(1)

func climbStairs(n int) int {
	var (
		prev = 1
		cur  = 1
		next = 0
	)

	for i := 2; i <= n; i++ {
		next = cur + prev
		prev = cur
		cur = next
	}

	return cur
}

// Dynamic Programming Tabulation
// Time Complexity: O(N)
// Space Complexity: O(N)

func climbStairs1(n int) int {
	var dp = make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// Dynamic Programming Memoization
// Time Complexity: O(N)
// Space Complexity: O(N)

func climbStairs2(n int) int {
	var dp = make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}
	return climb(n, dp)
}

func climb(n int, dp []int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if dp[n] != -1 {
		return dp[n]
	}

	steps := climb(n-1, dp) + climb(n-2, dp)
	dp[n] = steps
	return steps
}

// Recursion
// Time Complexity: O(N!)
// Space Complexity: O(N) call stack can have upto N calls at max

func climbStairs3(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return climbStairs3(n-1) + climbStairs3(n-2)
}
