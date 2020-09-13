/**
 题目:
    给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
    注意：答案中不可以包含重复的三元组。
    示例：
    给定数组 nums = [-1, 0, 1, 2, -1, -4]，
    满足要求的三元组集合为：
    [
      [-1, 0, 1],
      [-1, -1, 2]
    ]
 思路:
    1. 先排序,然后在一次遍历中使用指针对撞
    2. 利用for循环快速过滤掉重复元素
    3. 当遍历的值大于0后可直接返回结果 利用有序的特性进行剪枝
 效果
    
*/
class Solution {
public:
    class Solution {
public:
      vector<vector<int>> threeSum(vector<int>& nums) {
        int size=nums.size();
        if(size<3){
            return {};
        }
        sort(nums.begin(),nums.end());
        vector<vector<int>> pos;
        int left=0,right=size-1;
        while(left<right-1){
            
            int middle=left+1;
            int val=nums[left]+nums[right];
            int tmp=val*(-1);;
            while(middle<=right-1){
                vector<int> item;
                if(nums[middle]==tmp){
                    item.push_back(nums[left]);
                     item.push_back(nums[middle]);
                     item.push_back(nums[right]);
                      pos.push_back(item);
                }
                middle++;
            }
            if(val>0){
                right--;
            }else{
                left++;
            }
        }
       return pos;
    }
};
};
