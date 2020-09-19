/**
  题目:
    给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

    注意：

    答案中不可以包含重复的四元组。

    示例：

    给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

    满足要求的四元组集合为：
    [
      [-1,  0, 0, 1],
      [-2, -1, 1, 2],
      [-2,  0, 0, 2]
    ]

思路:
  双层循环 双指针 去重复
  
效果: 
  执行用时：16 ms, 在所有 Go 提交中击败了54.60%的用户
  内存消耗：2.8 MB, 在所有 Go 提交中击败了32.15%的用户
*/

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			k, l := j+1, len(nums)-1
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				distance := sum - target
				if distance == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[l]})
				}
				if distance > 0 {
					for cur := l; l > k && nums[l] == nums[cur]; l-- {}
				} else {
					for cur := k; k < l && nums[k] == nums[cur]; k++ {}
				}
			}
		}
	}
	return res
}
