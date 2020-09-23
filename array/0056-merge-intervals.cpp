/**
  题目:
      给出一个区间的集合，请合并所有重叠的区间。

 

  示例 1:

    输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
    输出: [[1,6],[8,10],[15,18]]
    解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
  示例 2:

    输入: intervals = [[1,4],[4,5]]
    输出: [[1,5]]
    解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
    注意：输入类型已于2019年4月15日更改。 请重置默认代码定义以获取新方法签名。
 
链接：
   https://leetcode-cn.com/problems/merge-intervals
思路:
     1.先进行排序(为了保证相对顺序不变，将一组先存放到一个pair中)
 
       
  效果:
    执行用时：40 ms, 在所有 C++ 提交中击败了86.77%的用户
    内存消耗：14.9 MB, 在所有 C++ 提交中击败了13.00%的用户
*/ 
class Solution {
public:
vector<vector<int>> res;
    vector<vector<int>> merge(vector<vector<int>>& intervals) {
        int size=intervals.size();
        if(size<2){
            return intervals;
        }
        vector<pair<int,int>> datas;
        for(int i=0;i<size;i++){
            datas.push_back({intervals[i][0],intervals[i][1]});
        }

        auto sortFunc=[](const pair<int,int>&a,const pair<int,int>&b){
            if(a.first<b.first){
                return true;
            }else if(a.first==b.first){
                return a.second<b.second;
            }else{
                return false;
            }
        };

        auto isInBound=[](const pair<int,int>&a,const pair<int,int>&b){
            if(a.second<b.first){
                return false;
            }else{
                return true;
            }
        };

        //数据是排序的
        sort(datas.begin(),datas.end(),sortFunc);
        int i=0;
        for(;i<size-1;){
            pair<int,int> a=datas[i];
            vector<int> c;
            c.push_back(a.first);
            c.push_back(a.second);
            int j=i+1;
            bool isEnd=true;
            for(;j<size;j++){
               
                pair<int,int> b=datas[j];
                if(isInBound(a,b)){
                    //那么需要进行合并了
                   c[1]=max(c[1],b.second);
                    a.second=c[1];
                    
                }else{
                    isEnd=false;
                    break;
                }
            }
           
            res.push_back(c);
            if(!isEnd) {
                i=j-1;
            }else{
                break;
            }
            ++i;
        }
        if(i==size-1){
            //表示最后一个需要加入
            pair<int,int> a=datas[size-1];
            vector<int> c;
            c.push_back(a.first);
            c.push_back(a.second);
             res.push_back(c);
        }
        return res;
    }
};
