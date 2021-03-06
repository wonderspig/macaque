/**
  题目:
    反转一个单链表。
    示例:
    输入: 1->2->3->4->5->NULL
    输出: 5->4->3->2->1->NULL
    进阶:
    你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
  思路:
    连表问题 要思考 0,1,2,3 节点个数连表的特殊情况
  效果:
    执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
    内存消耗：2.5 MB, 在所有 Go 提交中击败了45.41%的用户  
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil{
        return nil
    }
    var pre *ListNode
    cur, next := head, head.Next
    for cur != nil{
        cur.Next = pre
        pre = cur
        cur = next
        if next != nil{
            next = next.Next
        }
    }
    return pre
}
