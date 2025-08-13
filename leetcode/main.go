package main

import (
	"fmt"
	"go-study/leetcode/question"
	"go-study/leetcode/types"
)

func main() {
	//1.两数之和
	var nums = []int{2, 7, 11, 15}
	var target = 9
	TwoSumResult := question.TwoSum(nums, target)
	fmt.Println("----------------1.两数之和结果为-------------")
	fmt.Println(TwoSumResult)
	//2.两数相加
	var temp1 = types.ListNode{Val: 9}
	var temp2 = types.ListNode{Val: 9}
	var temp4 = types.ListNode{Val: 9}
	var temp5 = types.ListNode{Val: 9}
	temp1.Next = &temp2
	temp4.Next = &temp5
	var l1 = types.ListNode{Val: 2} //[2,1,4] 412
	l1.Next = &temp1
	var l2 = types.ListNode{Val: 7} //[7,5,6] 657
	l2.Next = &temp4
	AddTwoNumbersResult := question.AddTwoNumbers(&l1, &l2)
	for AddTwoNumbersResult != nil {
		fmt.Print(AddTwoNumbersResult.Val) //[9,6,0,1] 1069
		AddTwoNumbersResult = AddTwoNumbersResult.Next
	}
	fmt.Println("----------------2.两数之和结果为-------------")
	fmt.Println(AddTwoNumbersResult)
	//3.无重复字符的最长字串
	str := "abcabcbb"
	LengthOfLongestSubstringResult := question.LengthOfLongestSubstring(str)
	fmt.Println("----------------3.无重复字符的最长字串结果为-------------")
	fmt.Println(LengthOfLongestSubstringResult)
	//4.寻找两个正序数组的中位数
	nums1 := []int{1, 3, 5, 8, 10, 12}
	nums2 := []int{2, 4, 6, 9, 11}
	FindMedianSortedArraysResult := question.FindMedianSortedArrays(nums1, nums2)
	fmt.Println("----------------4.寻找两个正序数组的中位数结果为-------------")
	fmt.Println(FindMedianSortedArraysResult)

}
