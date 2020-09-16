/**
题目:
  给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
  说明:
  初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
  你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
  示例:
  输入:
  nums1 = [1,2,3,0,0,0], m = 3
  nums2 = [2,5,6],       n = 3
  输出: [1,2,2,3,5,6]
思路:
    从后向前三指针归并,并处理长短不一的情况
效果: 
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.3 MB, 在所有 Go 提交中击败了26.12%的用户
*/
func merge(nums1 []int, m int, nums2 []int, n int)  {
    i, j, k := m-1, n-1, m+n-1
    for i >= 0 && j >= 0{
        if nums1[i] < nums2[j]{
            nums1[k] = nums2[j]
            j--
        }else{
            nums1[k] = nums1[i]
            i--
        }
        k--
    }
    for i >= 0{
        nums1[k] = nums1[i]
        i--
        k--
    }
    for j >= 0{
        nums1[k] = nums2[j]
        j--
        k--
    }
    return 
}
