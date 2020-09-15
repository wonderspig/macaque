/**
  题目:
    给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
    示例:
    输入:
    [
      ["1","0","1","0","0"],
      ["1","0","1","1","1"],
      ["1","1","1","1","1"],
      ["1","0","0","1","0"]
    ]
    输出: 6
  思路:
      1. 将二维转化为一维,其矩阵最大矩形面积可以转化为求直方图的最大矩形面积
      2. 而一维数组所构建的直方图的最大矩形面积可以由递增单调栈来解决
  效果:
    执行用时：72 ms, 在所有 C++ 提交中击败了25.53%的用户
    内存消耗：11.6 MB, 在所有 C++ 提交中击败了13.59%的用户
*/
class Solution {
public:
    int maximalRectangle(vector<vector<char>>& matrix) {
        int n = matrix.size(), m, ans = 0;
        if (n == 0)
            return 0;

        m = matrix[0].size();
        vector<int> heights(m + 1, 0);
        heights[m] = -1;

        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++)
                if (matrix[i][j] == '0')
                    heights[j] = 0;
                else
                    heights[j]++;

            stack<int> st;
            for (int j = 0; j <= m; j++) {
                while (!st.empty() && heights[j] < heights[st.top()]) {
                    int cur = st.top();
                    st.pop();

                    if (st.empty())
                        ans = max(ans, heights[cur] * j);
                    else
                        ans = max(ans, heights[cur] * (j - st.top() - 1));
                }

                st.push(j);
            }
        }

        return ans;
 
    }
};