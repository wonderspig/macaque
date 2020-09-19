/**
  题目:
    给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
    说明：解集不能包含重复的子集。
    示例:

  输入: nums = [1,2,3]
  输出:
  [
    [3],
    [1],
    [2],
    [1,2,3],
    [1,3],
    [2,3],
    [1,2],
    []
  思路:
    定义选择列表 nums 选择路径path 结果集 res 以及 进行去重复的idx
  效果:
    执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
    内存消耗：2.3 MB, 在所有 Go 提交中击败了33.59%的用户
*/

func subsets(nums []int) [][]int {
    if len(nums) == 0{
        return  nil
    }
    res := make([][]int, 0)
    helper(nums,&res, []int{},0)
    return res
}
func helper(nums []int,res *[][]int, path []int, idx int){
    *res = append(*res, append([]int{},path...))
    for i := idx; i < len(nums);i++{
        path = append(path, nums[i])
        helper(nums, res, path, i+1)
        path = path[:len(path)-1]
    }
    return
}
