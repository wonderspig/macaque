/**
  题目:
    给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

    不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
    示例 1:
    给定数组 nums = [1,1,2], 
    函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。 
    你不需要考虑数组中超出新长度后面的元素。
    示例 2:
    给定 nums = [0,0,1,1,1,2,2,3,3,4],
    函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
    你不需要考虑数组中超出新长度后面的元素。
  思路:
    单向双指针遍历 确定两个指针在什么情况下更新以及发生交换
  效果:
  执行用时：16 ms, 在所有 C++ 提交中击败了90.06%的用户
  内存消耗：7.4 MB, 在所有 C++ 提交中击败了76.36%的用户
*/  
class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int size=nums.size();
        if(size<=1){
            return size;
        }
        int j=0,i=0;
        for(  i=0,j=0;i<size;i++){
            if(i+1<size&&nums[i]==nums[i+1]){
            }else{
                nums[j++]=nums[i];
            }
        }
        return j;
    }
};