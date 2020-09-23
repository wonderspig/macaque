/**
题目：
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
例如，给定三角形：
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
说明：
如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
思路:
  状态转移:dp[i][j] = triangle[i][j] + min(dp[i-1][j],dp[i-1][j-1])
效果:
  执行用时：4 ms, 在所有 Go 提交中击败了94.93%的用户
  内存消耗：3.3 MB, 在所有 Go 提交中击败了22.73%的用户
*/
func minimumTotal(triangle [][]int) int {
    dp := make([][]int, len(triangle))
    dp[0] = []int{triangle[0][0]}
    for i := 1; i < len(triangle);i++{
        dp[i] = make([]int, len(triangle[i]))
        dp[i][0] = triangle[i][0] + dp[i-1][0]
        dp[i][len(dp[i])-1] = triangle[i][len(dp[i])-1] + dp[i-1][len(dp[i-1])-1]
    }
    for i := 1; i < len(dp);i++{
        for j := 1; j < len(dp[i])-1;j++{
            dp[i][j] = triangle[i][j] + min(dp[i-1][j],dp[i-1][j-1])
        }
    }
    res := int(^uint(0)>>1)
    for i := 0; i < len(dp[len(dp)-1]);i++{
        res = min(res, dp[len(dp)-1][i])
    }
    return res
}

func min(a, b int)int{
    if a > b{
        return b
    }
    return a
}
