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
 class Solution {
public:
    int minimumTotal(vector<vector<int>>& triangle) {
        //这个题目是个dp题dp[i][j] = triangle[i][j] + min(dp[i-1][j],dp[i-1][j-1])

        int rows=triangle.size();
        vector<vector<long long>> dp(rows,vector<long long>(rows,INT_MAX));
        dp[0][0] = triangle[0][0];
        for(int i=1;i<rows;i++)
        {
            int cols=triangle[i].size();
            for(int j=0;j<cols;j++){
                if(j>0){
                    dp[i][j] = triangle[i][j] + min(dp[i-1][j],dp[i-1][j-1]);
                }
                dp[i][j]=min(dp[i][j],dp[i-1][j]+triangle[i][j]);
            }
        }
         long long res=INT_MAX;
        for(int i=0;i<rows;i++) res=min(res,dp[rows-1][i]);
        return res;
    }
};