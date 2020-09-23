/**
  题目:
      给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

    说明：解集不能包含重复的子集。

    示例:

        输入: nums = [1,2,3]
        输出:
          [
            [3],
            [1],
            [2],
            [1,2,3],
            [1,3],
            [ 2,3],
            [1,2],
            []
         ]

 
 链接：
   https://leetcode-cn.com/problems/subsets 
思路:
     1.典型的dfs+递归
    注意：
    可以看看排列和组合
 
       
  效果:
    执行用时：8 ms, 在所有 C++ 提交中击败了9.97%的用户
    内存消耗：7 MB, 在所有 C++ 提交中击败了55.45%的用户
*/ 


class Solution {
public:
vector<vector<int>> res;
    vector<vector<int>> subsets(vector<int>& nums) {
        vector<int> path;
        subsets(nums,0,nums.size(),path);
        return res;
    }
    void subsets(vector<int>&nums,int index,int size,vector<int>&path){
        res.push_back(path);
        for(int i=index;i<size;i++){
            path.push_back(nums[i]);
            subsets(nums,i+1,size,path);
            path.pop_back();
        }
    }
};