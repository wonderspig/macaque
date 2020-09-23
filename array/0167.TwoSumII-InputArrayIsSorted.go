/**
  题目:
    反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
    说明:
    1 ≤ m ≤ n ≤ 链表长度。
    示例:
    输入: 1->2->3->4->5->NULL, m = 2, n = 4
    输出: 1->4->3->2->5->NULL
    
  思路:
   双指针
   
  效果:
    执行用时：4 ms, 在所有 Go 提交中击败了96.96%的用户
    内存消耗：2.9 MB, 在所有 Go 提交中击败了100.00%的用户
*/
func twoSum(numbers []int, target int) []int {
    i, j := 0,len(numbers)-1
    for i < j{
        sum := numbers[i] + numbers[j]
        if sum == target{
            return []int{i+1,j+1}
        }else if sum > target{
            j --
        }else{
            i ++
        }
    }
    return []int{0,0}
}
