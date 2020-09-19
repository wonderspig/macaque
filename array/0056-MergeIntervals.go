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
    提示：
    intervals[i][0] <= intervals[i][1]
    
  思路:
    区间问题先排序,然后分情况讨论
  
   效果:
   
    执行用时：12 ms, 在所有 Go 提交中击败了82.04%的用户
    内存消耗：4.7 MB, 在所有 Go 提交中击败了19.34%的用户
*/
func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 ||  len(intervals) == 1{
        return intervals
    }
    kSort(intervals)
    res := make([][]int,0)
    res = append(res,intervals[0])
   // fmt.Printf("%+v\n",intervals)
    for i:=1;i<len(intervals);i++{
        if res[len(res)-1][1] >= intervals[i][0]{
            if intervals[i][1] > res[len(res)-1][1]{
                res[len(res)-1][1] = intervals[i][1]
            }
        }else{
            res = append(res,intervals[i])
        }
        // fmt.Printf("%+v\n",res)
    }
    return res
}
func kSort(intervals [][]int){
    if len(intervals)==0 || len(intervals) == 1{
        return 
    }
    l := sort2(intervals)
    kSort(intervals[:l])
    kSort(intervals[l:])
}

func sort2(intervals [][]int)int{
    i,j,l := 1,len(intervals)-1,intervals[0][0]
    for i<=j{
       if l > intervals[i][0]{
           i++
       }else{
           swap(i,j,intervals)
           j--
       }
    }
    swap(0,i-1,intervals)
    return i
}
func swap(a,b int,intervals [][]int){
    tmp := intervals [a]
    intervals[a] = intervals[b]
    intervals[b]=tmp
}
