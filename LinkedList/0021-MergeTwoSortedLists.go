/**
题目:
  将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
  示例：
  输入：1->2->4, 1->3->4
  输出：1->1->2->3->4->4
思路:
  通过哨兵节点简化了操作
效果:
  执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
  内存消耗：2.5 MB, 在所有 Go 提交中击败了38.79%的用户
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummp := &ListNode{}
    cur := dummp
    for l1 != nil || l2 != nil{
        a, b := int(^uint(0)>>1), int(^uint(0)>>1)
        if l1 != nil{
            a = l1.Val
        }
        if l2 != nil{
            b = l2.Val
        }
        if a < b{
            cur.Next = l1
            l1 = l1.Next
        }else if a > b{
            cur.Next = l2
            l2 = l2.Next
        }else {
            cur.Next = l1
            l1 = l1.Next
            cur = cur.Next
            cur.Next = l2
            l2 = l2.Next
        }
        cur = cur.Next
    }
    return dummp.Next
}
