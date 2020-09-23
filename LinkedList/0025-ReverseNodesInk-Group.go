/**
    题目:
          给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
          k 是一个正整数，它的值小于或等于链表的长度。
          如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
        示例：
        给你这个链表：1->2->3->4->5
        当 k = 2 时，应当返回: 2->1->4->3->5
        当 k = 3 时，应当返回: 3->2->1->4->5
        说明：
        你的算法只能使用常数的额外空间。
        你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
    思路:
        先统计长度,然后判断递归基,最后反转前k个节点,递归的对后面的节点按k反转
    效果:
        执行用时：4 ms, 在所有 Go 提交中击败了92.19%的用户
        内存消耗：3.6 MB, 在所有 Go 提交中击败了12.10%的用户    
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    cur, n := head, 0
    for cur != nil{
        cur = cur.Next
        n++
    }
    if n < k {
        return head
    }
    n = k
    var pre *ListNode
    cur, next := head, head.Next
    for n > 0{
        cur.Next = pre
        pre = cur
        cur = next
        if next != nil{
            next = next.Next
        }
        n--
    }
    head.Next = reverseKGroup(cur,k)
    return pre
}
