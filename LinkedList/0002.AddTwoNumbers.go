/**
  题目:
      给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
      如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
      您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
      示例：
      输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
      输出：7 -> 0 -> 8
      原因：342 + 465 = 807
   思路:
      1. 运算问题 考虑移除与进位的处理
      2. 连表问题 如果输出的连表的头节点发生变化,即采用哨兵节点会使得编码更简单
   效果:
      执行用时：8 ms, 在所有 Go 提交中击败了93.05%的用户
      内存消耗：4.9 MB, 在所有 Go 提交中击败了22.93%的用户
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
        dummy, c := &ListNode{}, 0
        cur := dummy
        for l1 != nil || l2 != nil{
            a, b := 0, 0
            if l1 != nil{
                a = l1.Val
                l1 = l1.Next
            }
            if l2 != nil{
                b = l2.Val
                l2 = l2.Next
            }
            node := &ListNode{Val:(c + a + b)%10}
            c = (c + a + b) / 10
            cur.Next = node
            cur = cur.Next
        }
        if c != 0{
            cur.Next = &ListNode{Val:c}
        }
        return dummy.Next
}
