/**
 题目:
    给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
    注意：答案中不可以包含重复的三元组。
    示例：
    给定数组 nums = [-1, 0, 1, 2, -1, -4]，
    满足要求的三元组集合为：
    [
      [-1, 0, 1],
      [-1, -1, 2]
    ]
 思路:
    1. 先排序,然后在一次遍历中使用指针对撞
    2. 利用for循环快速过滤掉重复元素
    3. 当遍历的值大于0后可直接返回结果 利用有序的特性进行剪枝
 效果
    执行用时：20 ms, 在所有 Go 提交中击败了100.00%的用户
    内存消耗：7.4 MB, 在所有 Go 提交中击败了26.75%的用户
*/

func threeSum(nums []int) [][]int {
   if len(nums) < 3{
       return [][]int{}
   }
   res := make([][]int,0)
   sort.Ints(nums)
   for i := 0; i < len(nums);i++{
       if nums[i] > 0{
           return res
       }
       j, k := i + 1, len(nums)-1
       for j < k{
           sum := nums[i] + nums[j] + nums[k]
           if sum == 0{
               res = append(res, []int{nums[i], nums[j], nums[k]})
               for ; j < k && nums[j] == nums[j+1]; j++{}
               j++
               for ; j < k && nums[k] == nums[k-1];k--{}
               k--
           }else if sum > 0{
               k--
           }else{
               j++
           }
       }
       for ; i < len(nums)-1 && nums[i] == nums[i+1];i++{}
   }
   return res
}
