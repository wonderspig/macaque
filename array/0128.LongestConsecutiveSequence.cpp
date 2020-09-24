/**
题目:
给定一个未排序的整数数组，找出最长连续序列的长度。
要求算法的时间复杂度为 O(n)。
示例:
输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。

思路:
    一次遍历,哈希表减少查找时间复杂度
效果:
  执行用时：8 ms, 在所有 Go 提交中击败了67.90%的用户
  内存消耗：3.5 MB, 在所有 Go 提交中击败了77.91%的用户
*/ 

class Solution {
public:
    int longestConsecutive(vector<int>& nums) {
        if(nums.empty()){
            return 0;
        }
        unordered_map<int,int> records;
        for(int i=0;i<nums.size();i++){
            records[nums[i]]=1;
        }
        int cnt=1;
        int res=1;
        for(auto iter=records.begin();iter!=records.end();iter++){
            int key=iter->first;
            cnt=1;
            if(records.count(key+1)==0){
                while(records.count(key-1)>0){
                ++cnt;
                key-=1;
            }
            res=max(res,cnt);
            }
        }
        return res;
    }
};