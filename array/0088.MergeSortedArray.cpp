/**
题目:
  给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
  说明:
  初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
  你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
  示例:
  输入:
  nums1 = [1,2,3,0,0,0], m = 3
  nums2 = [2,5,6],       n = 3
  输出: [1,2,2,3,5,6]
思路:
    从后向前三指针归并,并处理长短不一的情况
效果: 
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.3 MB, 在所有 Go 提交中击败了26.12%的用户
*/class Solution {
public:
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
        //滑动窗口
        int s1= m-1,s2=n-1,k=m+n-1;
        while(s1>=0&&s2>=0){
            if(nums2[s2]>nums1[s1]){
                nums1[k--]=nums2[s2--];
            }else{
                nums1[k--]=nums1[s1--];
            }
        }
        while(s2>=0){
            nums1[k--]=nums2[s2--];
        }
    }
};