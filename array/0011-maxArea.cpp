/**
  题目:
    给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
    说明：你不能倾斜容器，且 n 的值至少为 2。
    图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
    示例：
      输入：[1,8,6,2,5,4,8,3,7]
      输出：49
链接:
   https://leetcode-cn.com/problems/container-with-most-water/    
思路:
     指针对撞,较小的一方计算其面积并更新结果变量
     
效果:
    执行耗时:16 ms, 在所有 C++ 提交中击败了83.28%的用户
	内存消耗:7.4 MB, 在所有 C++ 提交中击败了85.89%的用户
*/ 

class Solution {
public:
    int maxArea(vector<int>& height) {
        int size=height.size();
        if(size<=1){
            return 0;
        }
        int left=0,right=size-1;
        int maxArea=0;
        while(left<right){
            maxArea=max(maxArea,min(height[left],height[right])*(right-left));
            if(height[left]<height[right]){
                left++;
            }else{
                right--;
            }
        }
        return maxArea;
    }
};