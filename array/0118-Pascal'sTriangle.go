/**
  题目:
      给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
      在杨辉三角中，每个数是它左上方和右上方的数的和。

      示例:

      输入: 5
      输出:
      [
           [1],
          [1,1],
         [1,2,1],
        [1,3,3,1],
       [1,4,6,4,1]
      ]
      通过次数105,874提交次数157,032
  思路:
    下标计算
    
  效果:
    执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
    内存消耗：2 MB, 在所有 Go 提交中击败了34.64%的用户
*/
func generate(numRows int) [][]int {
    if numRows == 0{
        return [][]int{}
    }
    if numRows == 1{
        return [][]int{[]int{1}}
    }
    res := make([][]int,0)
    res = append(res,[]int{1})
    res = append(res,[]int{1,1})
    for i := 3; i <= numRows; i++{    
        nums := make([]int, i)
        nums[0], nums[len(nums)-1] = 1, 1
        for j := 0; j < len(res[len(res)-1]) - 1;j++{
            nums[j+1] = res[len(res)-1][j] + res[len(res)-1][j+1]
        }
        res = append(res, nums)
    }
    return res
}
