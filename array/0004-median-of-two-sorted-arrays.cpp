/**
题目:
    给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。

	请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

	你可以假设 nums1 和 nums2 不会同时为空。

 
链接：
	https://leetcode-cn.com/problems/median-of-two-sorted-arrays
 

思路:
    1.采用二分思想，无非这里是两个数组


效果: 

执行用时：O(log(n+m))，56 ms, 在所有 C++ 提交中击败了21.25%的用户
内存消耗：O(1) 89 MB, 在所有 C++ 提交中击败了36.38%的用户
*/
class Solution
{
public:
    double findMedianSortedArrays(vector<int> &nums1, vector<int> &nums2)
    {
        int m = nums1.size();
        int n = nums2.size();
        int n1 = 0, n2 = 0;
        bool isOdd = (m + n) % 2 == 0;
        int middle = (n + m) / 2;
        int a, b = 0;
        for (int i = 0; i <= middle; i++)
        {
            a = b;
            if (n1 >= m || (n2 < n && (nums1[n1] >= nums2[n2])))
            {
                b = nums2[n2++];
            }
            else if (n2 >= n || (n1 < m && (nums1[n1] < nums2[n2])))
            {
                b = nums1[n1++];
            }
        }
        if (isOdd)
        {
            return (a + b) / 2.0;
        }
        else
        {
            return b;
        }
    }

    /*暴力解法*/
    double findMedianSortedArrays1(vector<int> &nums1, vector<int> &nums2)
    {
        for (int i = 0; i < nums2.size(); i++)
        {
            nums1.push_back(nums2[i]);
        }
        sort(nums1.begin(), nums1.end());
        int size = nums1.size();
        if (size % 2 == 0)
        {
            return (nums1[size / 2 - 1] + nums1[size / 2]) / 2.0;
        }

        return (nums1[size / 2]);
    }
};