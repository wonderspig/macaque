/**
  题目: 
    实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
    如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
    必须原地修改，只允许使用额外常数空间。
    以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
    1,2,3 → 1,3,2
    3,2,1 → 1,2,3
    1,1,5 → 1,5,1
  思路:
    观察排列结果, 选择最靠右侧的最小的数与一个尽可能接近该数的更大数交换,并且保证其降序排列
  效果:
    执行用时：4 ms, 在所有 Go 提交中击败了70.88%的用户
    内存消耗：2.5 MB, 在所有 Go 提交中击败了18.35%的用户
*/
func nextPermutation(nums []int)  {
    if len(nums) == 0 || len(nums) == 1{
        return 
    }
    for i := len(nums)-1;i > 0;i--{
        if i > 0 && nums[i] > nums[i-1]{
            sort.Ints(nums[i:])
            for j := i; j < len(nums);j++{
                if nums[j] > nums[i-1]{
                    nums[j], nums[i-1] = nums[i-1], nums[j]
                    return 
                }
            }
        }
    }
    sort.Ints(nums)
}
