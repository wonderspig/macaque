/**
题目:
  给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
  上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Mar
  cos 贡献此图。 
  示例: 
  输入: [0,1,0,2,1,0,1,3,2,1,2,1]
  输出: 6 
  Related Topics 栈 数组 双指针
链接:
   https://leetcode-cn.com/problems/trapping-rain-water/submissions/      
思路:
  1. 指针对撞, 维护左右最大值
  2. 找坑。若R[i]>height[i]且L[i]>height[i],k位置积水,且积水量为min(R[i], L[i]) - height[i].
 
     
效果:
    执行耗时:4 ms, 在所有 C++ 提交中击败了94.89%的用户
	内存消耗:7.1 MB, 在所有 C++ 提交中击败了27.31%的用户


思路:
  1. 指针对撞, 维护左右最大值
  2. 找坑。若R[i]>height[i]且L[i]>height[i],k位置积水,且积水量为min(R[i], L[i]) - height[i].
*/
 class Solution {
public:
    
    int trap(vector<int>& height){
        int size=height.size();
        if(size<=0){
            return 0;
        }
        vector<int> left(size),right(size);
        //计算左边哇槽
        for(int i=1;i<size;i++){
            left[i]=max(left[i-1],height[i-1]);
        }
         //计算右边哇槽
        for(int i=size-2;i>=0;i--){
            right[i]=max(right[i+1],height[i+1]);
        }
        int water=0;
        for(int i=0;i<size;i++){
            int level = min(left[i], right[i]);
			water += max(0, level - height[i]);
        }
        return water;
    }
    int trap1(vector<int>& height) {
        //这个题是单调递增栈
        stack<int>stk;
        int res=0;
        for(int i=0;i<height.size();i++){
            int last=0;
            while(stk.size()&&height[stk.top()]<=height[i]){
                //这里我们找的是最大的，所以需要弹出来
                int t=stk.top();
                stk.pop();
                res+=(i-t-1)*(height[t]-last);
                last=height[t];
            }
            if(stk.size()){
                res+=(i-stk.top()-1)*(height[i]-last);
            }
            stk.push(i);
        }
        return res;
    }
};