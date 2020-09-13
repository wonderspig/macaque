/**
题目:
  给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
	请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
	你可以假设 nums1 和 nums2 不会同时为空。
 
链接：
	https://leetcode-cn.com/problems/median-of-two-sorted-arrays
 
思路:
    1. 将问题转换为 求两个有序数组 第k个位置上的值的问题
    2. 那么中位数 就是第 (m + n + 1)/2 与 (m + n + 2)/2位置的值的和的一半(兼容奇偶)
    3. 通过将逻辑上的 m+n 个元素的第k位置进行平分(k/2)得到两个在各自数组中的下标
    4. 通过两数组各自有序的特性进行二分,不断递归最终找到 第k位置的元素值
效果: 
执行用时：20 ms, 在所有 Go 提交中击败了51.35%的用户
内存消耗：5.5 MB, 在所有 Go 提交中击败了97.55%的用户
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m, n := len(nums1), len(nums2)
    left  := getKth(nums1,nums2,(m + n + 1)/2)
    right := getKth(nums1,nums2,(m + n + 2)/2)
    return (float64(left) + float64(right)) / 2.0
}

func getKth(nums1,nums2 []int, k int)int{
    if len(nums1) == 0{
        return nums2[k-1]
    }
    if len(nums2) == 0{
        return nums1[k-1]
    }
    if k == 1{
        return min(nums1[0], nums2[0])
    }
    drop1, drop2 := min(k/2, len(nums1)), min(k/2, len(nums2)) 
    if nums1[drop1 - 1] <= nums2[drop2 - 1]{
        return getKth(nums1[drop1:],nums2, k-drop1)
    }else{
        return getKth(nums1,nums2[drop2:], k-drop2)
    }
}
func min(a, b int)int{
    if  a  < b{
        return a
    }
    return b
}
