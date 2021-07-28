package main

//爬楼梯(递归)
func climbStairs1(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < len(dp); i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}

//动态规划
func climbStairs2(n int) int {
	if n <= 3 {
		return n
	}

	a, b := 2, 3
	for i := 4; i <= n; i++ {
		b = b + a
		a = b - a
	}
	return b
}

//跳跃游戏
// 思路 动态规划法 dp[i] 表示 i下标能不能到达
// dp[i] 能不能到达 取决于 dp[0:i-1] 中能到达的地方能不能到达dp[i]
//  dp[j] 为 dp[i]  前的某个位置  那么当 dp[j]=true 并且 nums[j]>=i-j 那么 i就能到达
func canJUmp(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	dp := make([]bool, len(nums))
	// 第一个肯定能到达(起始条件)
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && nums[j]+j >= i {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(nums)-1]
}
