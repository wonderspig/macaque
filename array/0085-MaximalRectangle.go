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
    执行用时：4 ms, 在所有 Go 提交中击败了83.05%的用户
    内存消耗：3.2 MB, 在所有 Go 提交中击败了98.81%的用户
*/
func maximalRectangle(matrix [][]byte) int {
    if len(matrix) == 0{
        return 0
    }
    m, res := make([]int,len(matrix[0])), 0
    for i := 0; i < len(matrix); i++{
        for j := 0; j < len(matrix[0]);j++{
            if matrix[i][j] == '1'{
                m[j]+=1
            }else{
                m[j]=0
            }
        }
        res = max(res, maxArea(m))
    }
    return res
}
func maxArea(nums []int)int{
    s,res := []int{-1},0
    for i, num := range nums{
        for s[len(s)-1] != -1 &&nums[s[len(s)-1]] >= num{
            hIdx := s[len(s)-1]
            s = s[:len(s)-1]
            res = max(res,(i - s[len(s)-1] - 1) * nums[hIdx])
        }
        s = append(s, i)
    }
    for s[len(s)-1] != -1{
        hIdx := s[len(s)-1]
        s = s[:len(s)-1]
        res = max(res,(len(nums) - s[len(s)-1] - 1) * nums[hIdx])
    }
    return res
}
func max(a, b int)int{
    if a > b{
        return a
    }
    return b
}
