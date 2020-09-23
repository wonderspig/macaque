/**
  题目:
    给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

    注意：

    答案中不可以包含重复的四元组。

    示例：

    给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

    满足要求的四元组集合为：
    [
      [-1,  0, 0, 1],
      [-2, -1, 1, 2],
      [-2,  0, 0, 2]
    ]

思路:
  双层循环 双指针 去重复
  
效果: 
  执行用时：16 ms, 在所有 Go 提交中击败了54.60%的用户
  内存消耗：2.8 MB, 在所有 Go 提交中击败了32.15%的用户
*/
class Solution {
public:
    vector<vector<int>> fourSum(vector<int>& nums, int target) {
        int size=nums.size();
        if(size<4){
            return {};
        }
        vector<vector<int>> res;
        sort(nums.begin(),nums.end());
        for(int i=0;i<size-3;i++){
            //考虑到相同的值
            if(i>0&&nums[i]==nums[i-1]){
                continue;
            }
            for(int j=i+1;j<size-2;j++){
                //考虑到相同的值
            if(j>i+1&&nums[j]==nums[j-1]){
                continue;
            }
            int k=j+1,l=size-1;
           
            while(k<l){
                int sum=nums[i]+nums[j]+nums[k]+nums[l];
                int distance=sum-target;
                if(distance==0){
                     vector<int> record;
                    //刚好找到了
                    record.emplace_back(nums[i]);
                    record.emplace_back(nums[j]);
                    record.emplace_back(nums[k]);
                    record.emplace_back(nums[l]);
                   
                    res.push_back(record);
                }
                 if(distance>0){
                    //目的是为了快速收敛
                    for(int cur=l;l>k&&nums[l]==nums[cur];l--){}
                }else{
                     
                    //目的是为了快速收敛.去除重复的值
                    for(int cur=k;l>k&&nums[k]==nums[cur];k++){}
                
                }
            }
            }
        }
        return res;
    }
};