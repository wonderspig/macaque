/**
  题目:
    给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
    你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
    示例:
    给定 1->2->3->4, 你应该返回 2->1->4->3.
  思路:
    头节点会变所有要 用哨兵 想清楚循环不变式
  效果:
    执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
    内存消耗：2.1 MB, 在所有 Go 提交中击败了28.77%的用户
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    pre := &ListNode{Next:head}
    cur := pre
    for cur != nil && cur.Next != nil && cur.Next.Next != nil{
        a, b := cur.Next, cur.Next.Next
        a.Next = b.Next
        b.Next = a
        cur.Next = b
        cur = cur.Next.Next
    }
    return pre.Next
}
