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
 /**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* reverseKGroup(ListNode* head, int k) {
        if(!head||k<=0){
            return head;
        }
        //
        ListNode* dummy=new ListNode(-1);
        dummy->next=head;
    
        ListNode*pre=dummy;
        ListNode*cur=head;
        ListNode*last=nullptr;
        int cnt=0;
   
        while(cur){
            
            if(++cnt%k==0){
                pre=reverseList(pre,cur);
                cur=pre->next;
            }else{
                cur=cur->next;
            }
        }
        return dummy->next;
    }
    ListNode* reverseList(ListNode* s,ListNode*e){
        ListNode*last=s->next;
        ListNode*cur=last->next;
        ListNode*end=e->next;
        while(cur!=end){
            last->next = cur->next;
            cur->next = s->next;
            s->next = cur;
            cur = last->next;
        }
        return last;
    }
};
