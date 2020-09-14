/**
题目:
  给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
  示例:
  输入: [-2,1,-3,4,-1,2,1,-5,4]
  输出: 6
  解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
  进阶:
  如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解

链接：
  https://leetcode-cn.com/problems/maximum-subarray/submissions/
思路:
  局部最优解,不断更新最后得到全局最优解（实际上是个dp问题）
  
效果:
  执行用时：12 ms, 在所有 C++ 提交中击败了20.85%的用户
  内存消耗：13.2 MB, 在所有 C++ 提交中击败了10.12%的用户
*/
class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        int size=nums.size();
        if(size<=0){
          return -1;
        }
        int maxv=nums[0];
        //需要注意的是，pre不一定是结果中的最大值
        int pre=nums[0];
        for(int index=1;index<size;index++){
            pre=max(pre,0)+nums[index];
            maxv=max(pre,maxv);
        }
        return maxv;
    }
};


class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        int size=nums.size();
        vector<int> dp(size,0);
        dp[0]=nums[0];
        int maxv=nums[0];
        for(int index=1;index<size;index++){
            dp[index]=max(dp[index-1]+nums[index],nums[index]);
            maxv=max(dp[index],maxv);
        }
        return maxv;
    }
};