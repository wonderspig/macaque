/**
  题目:
    给你一个链表数组，每个链表都已经按升序排列。
    请你将所有链表合并到一个升序链表中，返回合并后的链表。
    示例 1：

    输入：lists = [[1,4,5],[1,3,4],[2,6]]
    输出：[1,1,2,3,4,4,5,6]
    解释：链表数组如下：
    [
      1->4->5,
      1->3->4,
      2->6
    ]
    将它们合并到一个有序链表中得到。
    1->1->2->3->4->4->5->6
    示例 2：

    输入：lists = []
    输出：[]
    示例 3：

    输入：lists = [[]]
    输出：[]


    提示：

    k == lists.length
    0 <= k <= 10^4
    0 <= lists[i].length <= 500
    -10^4 <= lists[i][j] <= 10^4
    lists[i] 按 升序 排列
    lists[i].length 的总和不超过 10^4
    
 思路: 递归 + 化简问题为合并两个有序列表
 
 效果:
    执行用时：8 ms, 在所有 Go 提交中击败了94.91%的用户
    内存消耗：5.3 MB, 在所有 Go 提交中击败了66.95%的用户
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    n := len(lists)
    if n == 0{
        return nil
    }
    if n ==1{
        return lists[0]
    }
    if n == 2{
        return merge2List(lists[0],lists[1])
    }
    key := n/2
    return merge2List(mergeKLists(lists[:key]),mergeKLists(lists[key:]))
}

func merge2List(node1,node2 *ListNode)*ListNode{
    if node1 == nil{
        return node2
    }
    if node2 == nil{
        return node1
    }
    dummp := &ListNode{}
    curr := dummp
    for node1 != nil && node2 != nil{
         if node1.Val > node2.Val{
            curr.Next = node2
            curr = curr.Next
            node2 = node2.Next
         }else{
             curr.Next = node1
             curr = curr.Next
             node1 = node1.Next
         }
    }
    if node1 != nil{
        curr.Next = node1
    }else{
        curr.Next = node2
    }
    return dummp.Next
}
