/**
 题目:
    给定一个非负整数数组，你最初位于数组的第一个位置。

    数组中的每个元素代表你在该位置可以跳跃的最大长度。

    你的目标是使用最少的跳跃次数到达数组的最后一个位置。

    示例:

    输入: [2,3,1,1,4]
    输出: 2
    解释: 跳到最后一个位置的最小跳跃数是 2。
         从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
    说明:

    假设你总是可以到达数组的最后一个位置。

  思路:
    dp 为从0 到 i 的最少跳跃次数
    
  效果:
    执行用时：744 ms, 在所有 Go 提交中击败了11.99%的用户
    内存消耗：4.5 MB, 在所有 Go 提交中击败了37.96%的用户
*/

// O(n^2)
func jump(nums []int) int {
    dp := make([]int,len(nums))
    for i := 0;i < len(dp);i++{
        dp[i] = int(^uint(0)>>1)
    }
    dp[0] = 0
    for i := 1; i < len(nums);i++{
        for j := i - 1; j >= 0;j--{
            if nums[j] + j >= i{
                dp[i] = min(dp[i],dp[j] + 1)
            }
        }
    }
    return dp[len(dp)-1]
}

func min(a, b int)int{
    if a < b{
        return a
    }
    return b
}
// O(n)

func jump(nums []int) int {
    length := len(nums)
    end := 0
    maxPosition := 0
    steps := 0
    for i := 0; i < length - 1; i++ {
        maxPosition = max(maxPosition, i + nums[i])
        if i == end {
            end = maxPosition
            steps++
        }
    }
    return steps
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
