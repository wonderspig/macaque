/**
  题目:
    给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
    说明：每次只能向下或者向右移动一步。
     示例:
    输入:
    [
    [1,3,1],
    [1,5,1],
    [4,2,1]
    ]
    输出: 7
    解释: 因为路径 1→3→1→1→1 的总和最小。
  思路:
    一般性的dp思路
  效果:
   执行用时：8 ms, 在所有 Go 提交中击败了90.59%的用户
   内存消耗：4.3 MB, 在所有 Go 提交中击败了15.79%的用户
*/
func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([][]int,len(grid))
    for i := 0; i < len(grid);i++{
        dp[i] = make([]int, len(grid[i]))
    }
    dp[0][0] = grid[0][0]
    for i := 1; i < len(grid[0]); i++{
        dp[0][i] = grid[0][i] + dp[0][i-1]
    }
    for i := 1; i < len(grid); i++{
        dp[i][0] = grid[i][0] + dp[i-1][0]
    }
    for i := 1; i < len(dp);i++{
        for j := 1; j < len(dp[0]);j++{
            dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
        }
    }
    return dp[m-1][n-1]
}
func min(a, b int)int{
    if a < b{
        return a
    }
    return b
}
