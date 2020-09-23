/**
  题目:
    给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
    说明：每次只能向下或者向右移动一步。
     示例:
    输入:
    [
    [1,3,1]
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
 class Solution {
public:
    int minPathSum(vector<vector<int>>& grid) {
        int m=grid.size();
        int n=grid[0].size();
        //dps数组
        vector<vector<int>> dp(m,vector<int>(n));
        //表示只有一个元素的时候
          dp[0][0]=grid[0][0];
          //只有一行的时候
        for(int row=1;row<m;row++){
            dp[row][0]=dp[row-1][0]+grid[row][0];
        }
      for(int col=1;col<n;col++){
            dp[0][col]=dp[0][col-1]+grid[0][col];
        }
        //开始进行计算，实际就是更新迭代的过程
        for(int row=1;row<m;row++){
            for(int col=1;col<n;col++){
               
                dp[row][col]=min(dp[row-1][col],dp[row][col-1])+grid[row][col];
            }
        }
        return dp[m-1][n-1];
    }
};