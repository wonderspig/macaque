/**
  题目:
      给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
      在杨辉三角中，每个数是它左上方和右上方的数的和。

      示例:

      输入: 5
      输出:
      [
           [1],
          [1,1],
         [1,2,1],
        [1,3,3,1],
       [1,4,6,4,1]
      ]
      通过次数105,874提交次数157,032
  思路:
    下标计算
    
  效果:
    执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
    内存消耗：2 MB, 在所有 Go 提交中击败了34.64%的用户
*/

 class Solution {
public:
    vector<vector<int>> generate(int numRows) {
        vector<vector<int>> res;
        for(int i=1;i<=numRows;i++){
            //第一行
            vector<int> item;
            {
               if(res.empty()){
                   //表示是第一行
                   item.push_back(1);
               }else{
                   //表示第二行
                   vector<int> pre=res[res.size()-1];
                   item.push_back(1);
                   for(int i=0;i<pre.size()-1;i++){
                       item.push_back(pre[i]+pre[i+1]);
                   }
                   item.push_back(1);

               }
            }
            if(!item.empty()){
                res.push_back(item);
            }
        }
        return res;
    }
};