
/**
题目
    给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
    示例:
    输入: 3
    输出:
    [
     [ 1, 2, 3 ],
     [ 8, 9, 4 ],
     [ 7, 6, 5 ]
    ]
思路:
    通过 上下左右 四个指针控制二维数组的遍历
效果: 
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.1 MB, 在所有 Go 提交中击败了33.11%的用户
*/
func generateMatrix(n int) [][]int {
    res := make([][]int,n)
    for i := 0; i < len(res); i++{
        res[i] = make([]int, n)
    }
    left, right, up, down, counter := 0, len(res)-1, 0, len(res[0])-1, 1
    for {
        for i := left; i <= right; i++{
            res[up][i] = counter
            counter++
        }
        up++
        if up > down{
            break
        }
        for i := up; i <= down; i++{
            res[i][right] = counter
            counter++
        }
        right--
        if right < left{
            break
        }
        for i := right; i >= left; i--{
            res[down][i] = counter
            counter++
        }
        down--
        if down < up{
            break
        }
        for i := down; i >= up;i--{
            res[i][left] = counter
            counter++
        }
        left++
        if left > down{
            break
        }
    }
    return res
}
