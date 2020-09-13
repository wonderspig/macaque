/**
题目
    给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
    你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
    示例:
    给定 nums = [2, 7, 11, 15], target = 9
    因为 nums[0] + nums[1] = 2 + 7 = 9
    所以返回 [0, 1]

思路:
    1. 搜索两个数和等于目标值的下标, 在一次遍历中需要记忆住遍历过的值与下标的映射关系
    2. a + b = target 所以 a = target - b 因此目标值减当前元素即可得知a
    3. 如果a被遍历过并且下标不等于当前元素下标 即是搜索结果

效果:
执行用时：4 ms在所有 Go 提交中击败了96.97%的用户
内存消耗：3.8 MB在所有 Go 提交中击败了23.11%的用户
*/
func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, num := range nums{
        if v, ok := m[target-num];ok && v != i{
            return []int{i,v}
        }
        m[num]=i
    }
    return []int{}
}
