/**
  题目:
      给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
      示例：
      输入：nums = [-1,2,1,-4], target = 1
      输出：2
      解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。 
      提示：
      3 <= nums.length <= 10^3
      -10^3 <= nums[i] <= 10^3
      -10^4 <= target <= 10^4
      
  思路:
       排序,遍历,内部指针对撞,计算目标值与三数和的最小绝对值
       
  效果:
       执行用时：4 ms, 在所有 Go 提交中击败了94.83%的用户
       内存消耗：2.8 MB, 在所有 Go 提交中击败了20.00%的用户
*/
func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    res, minGap := 0, int(^uint(0)>>1)
    for i := 0; i < len(nums);i++{
        left, right := i + 1, len(nums)-1
        for left < right{
            sum := nums[i] + nums[left] + nums[right]
            if sum == target{
                return sum
            }
            if sum < target{
                left++
            }else{
                right--
            }
            gap := 0
            if target > sum{
                gap = target - sum
            }else{
                gap = sum - target
            }
            if minGap > gap{
                minGap = gap
                res = sum
            }
        }
    }
    return res
}
