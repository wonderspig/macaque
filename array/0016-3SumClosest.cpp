/**
  题目:
      给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
      示例：
      输入：nums = [-1,2,1,-4], target = 1
      输出：2
      解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。 
      提示：
      3 <= nums.length <= 10^3
      -10^3 <= nums[i] <= 10^3
      -10^4 <= target <= 10^4
 链接：
    https://leetcode-cn.com/problems/3sum-closest/submissions/     
  思路:
       排序,遍历,内部指针对撞,计算目标值与三数和的最小绝对值
       
  效果:
    执行用时：6 ms, 在所有 Java 提交中击败了85.91%的用户
    内存消耗：39.3 MB, 在所有 Java 提交中击败了82.69%的用户
*/ 



class Solution {
public:
    int threeSumClosest(vector<int>& nums, int target) {
        int size=nums.size();
        if(size<2){
            return 0;
        }
        sort(nums.begin(),nums.end());
        int closestNum=nums[0]+nums[1]+nums[2];
        
        for(int l=0;l<size-2;l++){
            int middle=l+1,r=size-1;
            while(middle<r){
                int val=nums[l]+nums[middle]+nums[r];
                if(abs(val-target)<abs(closestNum-target)){
                    closestNum=val;
                    
                }
                if(val>target){
                    r--;
                }else if(val<target){
                    middle++;
                }else{
                    return target;
                }
            }
        }
        return closestNum;
    }
};