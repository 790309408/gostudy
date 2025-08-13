/*
给你两个非空的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字0之外，这两个数都不会以 0 开头。
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807
示例 2：
输入：l1 = [0], l2 = [0]
输出：[0]
示例 3：
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
*/
package question

import "go-study/leetcode/types"

/*每个节点值都是0-9之间的数字
1.先判断链表1,链表2是否为空，遍历链表1，链表2，并求和
2. 两个节点值相加再加上进位值得和值sum。sum%10为当前节点值，sum/10为进位值
3.判断首节点是否为空，为空则赋值给它，不为空则赋值给当前尾值节点的next
4.最后判断进位值是否有进位值，有的话则赋值给尾值
*/
func AddTwoNumbers(l1 *types.ListNode, l2 *types.ListNode) (head *types.ListNode) {
	//1.声明尾部节点tail,进位值carry
	var tail *types.ListNode
	carry := 0
	//2.遍历链表
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		//3.将当前节点值取出，并指向下一个节点
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		//4.计算当前节点值+进位值carry
		sum := n1 + n2 + carry
		//5.sum为新的节点值，carry为进位值
		sum, carry = sum%10, sum/10
		//6.如果链表首节点为nil，则将当前的和值为链表首节点，否则将当前和值为链表末尾节点的next节点
		if head == nil {
			head = &types.ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &types.ListNode{Val: sum}
			tail = tail.Next
		}
	}
	//7.进位值指向尾部节点
	if carry > 0 {
		tail.Next = &types.ListNode{Val: carry}
	}
	return

}
