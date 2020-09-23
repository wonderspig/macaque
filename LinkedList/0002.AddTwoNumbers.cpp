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
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
         ListNode*head1=l1;
          ListNode*head2=l2;
        ListNode*h=nullptr;
        ListNode*dummy=nullptr;
         int val=0;
         while(head1||head2){
             int sum=0;
             if(head1){
                 sum= head1->val;
                 head1= head1->next;
             }
              if(head2){
                 sum+=head2->val;
                 head2= head2->next;
             }
             
                ListNode*tmp=new ListNode((val+sum)%10);
                if(!h){
                    dummy=h=tmp;
                    
                }else{
                    h->next=tmp;
                    h=tmp;
                }
                val= (sum+val)>=10?1:0;
         }
         if(val>0){
              ListNode*tmp=new ListNode((val)%10);
                if(!h){
                    dummy=h=tmp;
                    
                }else{
                    h->next=tmp;
                    h=tmp;
                }
         }
         h->next=nullptr;
         return dummy;
    }
};