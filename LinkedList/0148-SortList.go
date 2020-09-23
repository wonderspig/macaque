/**
  题目:
    在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
    示例 1:

    输入: 4->2->1->3
    输出: 1->2->3->4
    示例 2:

    输入: -1->5->3->4->0
    输出: -1->0->3->4->5
    
  思路:
     归并,快排,都可以
  
  效果:
      执行用时：1228 ms, 在所有 Go 提交中击败了5.11%的用户
      内存消耗：10.1 MB, 在所有 Go 提交中击败了5.16%的用户
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    mid := getMind(head)
    l1, l2 := head, mid.Next
    mid.Next = nil
    s1,s2 := sortList(l1), sortList(l2)
    return merge(s1,s2)
}
func getMind(head *ListNode)*ListNode{
    fast,pre := head,&ListNode{Next:head}
    if fast != nil && fast.Next != nil{
        pre = pre.Next
        fast = fast.Next.Next
    }
    return pre
}
//    l1
//  1-3-5-7
//  c
//  2-4-6
//  l2 
func merge(l1,l2 *ListNode)*ListNode{
    if l1 == nil{
        return l2
    }
    if l2 == nil{
        return l2
    }
    d := &ListNode{}
    cur := d
    for l1 != nil && l2 != nil{
        if l1.Val < l2.Val{
            cur.Next = l1
            l1 = l1.Next
        }else{
            cur.Next = l2
            l2 = l2.Next
        }
        cur = cur.Next
    }
    if l1 == nil{
        cur.Next = l2
    }
    if l2 == nil{
        cur.Next = l1
    } 
    return d.Next
}
