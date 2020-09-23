/**
  题目:
      给定一个非负整数数组，你最初位于数组的第一个位置。

    数组中的每个元素代表你在该位置可以跳跃的最大长度。

    你的目标是使用最少的跳跃次数到达数组的最后一个位置。

  示例:

    输入: [2,3,1,1,4]
    输出: 2
  解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
  说明:

    假设你总是可以到达数组的最后一个位置。
 
 
 链接：
   https://leetcode-cn.com/problems/jump-game-ii
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