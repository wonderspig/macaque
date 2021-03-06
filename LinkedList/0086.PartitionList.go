/**
  题目:
    给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
    你应当保留两个分区中每个节点的初始相对位置。
  示例:
  输入: head = 1->4->3->2->5->2, x = 3
  输出: 1->2->2->4->3->5  
  
  思路:
    正常的连表 遍历 注意 bad case
  
  效果:
    执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
    内存消耗：2.4 MB, 在所有 Go 提交中击败了28.07%的用户
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    a, b := &ListNode{}, &ListNode{}
    aa,bb := a,b
    for head != nil{
        if head.Val >= x{
            bb.Next = head
            bb = bb.Next
        }else if head.Val < x{
            aa.Next = head
            aa = aa.Next
        }
        head = head.Next
    }
    bb.Next = nil
    aa.Next = b.Next
    return a.Next
}
