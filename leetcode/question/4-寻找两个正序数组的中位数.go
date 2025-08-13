/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
请你找出并返回这两个正序数组的中位数 。
算法的时间复杂度应该为 O(log (m+n)) 。
示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
*/
package question

import (
	"fmt"
	"math"
)

func FindMedianSortedArrays(a []int, b []int) float64 {
	// 暴力解法
	// //1.获取两个数组的长度
	// sumLen := len(nums1) + len(nums2)
	// //2.判断奇数还是偶数
	// isEven := sumLen%2 == 0
	// //3.合并数组
	// mergearray := append(nums1, nums2...)
	// //4.排序
	// sort.Ints(mergearray)
	// if isEven {
	// 	sum := mergearray[sumLen/2] + mergearray[sumLen/2-1]
	// 	return float64(sum) / 2
	// } else {
	// 	sum := mergearray[(sumLen-1)/2]
	// 	return float64(sum)
	// }
	if len(a) > len(b) {
		a, b = b, a
	}
	m, n := len(a), len(b)
	// 循环不变量：a[left] <= b[j+1]
	// 循环不变量：a[right] > b[j+1]
	left, right := -1, m
	for left+1 < right { // 开区间 (left, right) 不为空
		i := (left + right) / 2
		j := (m+n+1)/2 - i - 2
		if a[i] <= b[j+1] {
			left = i // 缩小二分区间为 (i, right)
		} else {
			right = i // 缩小二分区间为 (left, i)
		}
	}
	fmt.Println(left, right)
	// 此时 left 等于 right-1
	// a[left] <= b[j+1] 且 a[right] > b[j+1] = b[j]，所以答案是 i=left
	i := left
	j := (m+n+1)/2 - i - 2
	ai := math.MinInt
	if i >= 0 {
		ai = a[i]
	}
	bj := math.MinInt
	if j >= 0 {
		bj = b[j]
	}
	ai1 := math.MaxInt
	if i+1 < m {
		ai1 = a[i+1]
	}
	bj1 := math.MaxInt
	if j+1 < n {
		bj1 = b[j+1]
	}
	max1 := max(ai, bj)
	min2 := min(ai1, bj1)
	if (m+n)%2 > 0 {
		return float64(max1)
	}
	return float64(max1+min2) / 2
}
