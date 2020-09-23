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
 /**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        ListNode*dummy=new ListNode(-1);
        ListNode*head=dummy;
        head->next=nullptr;
        while(l1&&l2){
            if(l1->val>l2->val){
                head->next=l2;
                head=l2;
                l2=l2->next;
            }else{
                head->next=l1;
                head=l1;
                l1=l1->next;
            }
        }
        if(l1){
            head->next=l1;
        }
        if(l2){
             head->next=l2;
        }
        return dummy->next;
    }
};