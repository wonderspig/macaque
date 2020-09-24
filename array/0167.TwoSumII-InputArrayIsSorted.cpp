/**
  题目:
    反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
    说明:
    1 ≤ m ≤ n ≤ 链表长度。
    示例:
    输入: 1->2->3->4->5->NULL, m = 2, n = 4
    输出: 1->4->3->2->5->NULL
    
  思路:
   双指针
   
  效果:
    执行用时：4 ms, 在所有 Go 提交中击败了96.96%的用户
    内存消耗：2.9 MB, 在所有 Go 提交中击败了100.00%的用户
*/
class Solution {
public:
    vector<int> twoSum(vector<int>& numbers, int target) {
        if(numbers.empty()){
            return {-1,-1};
        }
        int left=0,right=numbers.size()-1; 
        while(left<right){
            if(numbers[left]+numbers[right]==target){
                return {left+1,right+1};
            }else if(numbers[left]+numbers[right]>target){
                right--;
            }else{
                left++;
            }
        }
        return {-1,-1};
    }
};