/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值 target  的那两个整数，
并返回它们的数组下标。你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
*/
package question

/*哈希法
使用哈希表，可以将寻找 target - x 的时间复杂度降低到从 O(N) 降低到 O(1)。
这样我们创建一个哈希表，对于每一个 x，我们首先查询哈希表中是否存在 target - x，然后将 x 插入到哈希表中，即可保证不会让 x 和自己匹配。
哈希表（Hash Table），又称散列表，是一种通过​​键（Key）直接访问值（Value）​​的高效数据结构。
其核心思想是通过哈希函数（Hash Function）将键映射到存储位置（称为哈希地址），
从而实现平均时间复杂度为 ​​O(1)​​ 的查找、插入和删除操作
*/
func TwoSum(nums []int, target int) []int {
	// 1.创建哈希表。key是数值类型，value也是数值类型
	hashTable := map[int]int{}
	// 2.遍历nums数组
	// i表示下标，x表示数值
	for i, x := range nums {
		// 3.判断某个键是否存在
		// 因为两个值之和是target，所以只需要判断target-x的值是否存在于哈希表中即可
		// 如果目标值target减去x的结果值存在，则直接返回target-x的索引和当前索引i
		// p为target-x的值, ok表示是否存在的key值 target-x
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		// 4.将nums[i]的数值作为key，下标i作为value保存到哈希表中
		hashTable[x] = i
	}
	return nil
}
