/**
题目:
给定一个未排序的整数数组，找出最长连续序列的长度。
要求算法的时间复杂度为 O(n)。
示例:
输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。

思路:
    一次遍历,哈希表减少查找时间复杂度
效果:
  执行用时：8 ms, 在所有 Go 提交中击败了67.90%的用户
  内存消耗：3.5 MB, 在所有 Go 提交中击败了77.91%的用户
*/
func longestConsecutive(nums []int) int {
    m, res := make(map[int]bool), 0
    for _, num := range nums{
        m[num] = true
    }
    for i := 0; i < len(nums);i++{
        curNum, matchLen := 0, 1
        if !m[nums[i] - 1]{
            curNum = nums[i]
            for m[curNum + 1]{
                matchLen++
                curNum++
            }
            res = max(res, matchLen)
        }
        m[nums[i]]=true
    }
    return res
}

func max(a, b int)int{
    if a > b {
        return a
    }
    return b
}
