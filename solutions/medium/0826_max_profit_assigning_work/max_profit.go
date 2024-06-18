/*
826. Most Profit Assigning Work
https://leetcode.com/problems/most-profit-assigning-work/

You have n jobs and m workers. You are given three arrays: difficulty, profit, and worker where:

difficulty[i] and profit[i] are the difficulty and the profit of the ith job, and
worker[j] is the ability of jth worker (i.e., the jth worker can only complete a job with difficulty at most worker[j]).
Every worker can be assigned at most one job, but one job can be completed multiple times.

For example, if three workers attempt the same job that pays $1, then the total profit will be $3. If a worker cannot complete any job, their profit is $0.
Return the maximum profit we can achieve after assigning the workers to the jobs.
*/

package maxprofit

// Time Complexity: O(max(N, M)) N -> no of jobs M -> no of workers
// Space Complexity: O(max(difficulty[i]))

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {

	var maxDifficulty = 0

	for _, diff := range difficulty {
		if maxDifficulty < diff {
			maxDifficulty = diff
		}
	}

	var maxProfitTillDiff = make([]int, maxDifficulty+1)
	for index := range profit {
		maxProfitTillDiff[difficulty[index]] = max(maxProfitTillDiff[difficulty[index]], profit[index])
	}

	for index := 1; index < len(maxProfitTillDiff); index++ {
		maxProfitTillDiff[index] = max(maxProfitTillDiff[index], maxProfitTillDiff[index-1])
	}

	var totalProfit = 0

	for _, workerCap := range worker {
		if workerCap > maxDifficulty {
			totalProfit += maxProfitTillDiff[maxDifficulty]
		} else {
			totalProfit += maxProfitTillDiff[workerCap]
		}
	}

	return totalProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
