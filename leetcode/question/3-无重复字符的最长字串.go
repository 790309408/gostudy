/*
给定一个字符串 s,请你找出其中不含有重复字符的最长子串的长度。
示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。

	请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/
package question

func LengthOfLongestSubstring(s string) int {
	// 双指针记录索引，map存储当前字符串，如果有重复,记录当前字符串长度，并将map置空，startIndex索引加1,并将值赋值给moveIndex。
	//1.自己解法
	// var charMap = make(map[string]string)
	// var maxLength = 0
	// var startIndex = 0
	// var moveIndex = startIndex
	// for startIndex < len(s) && moveIndex < len(s) {
	// 	//1.获取当前字符
	// 	key := string(s[moveIndex])

	// 	//2.判断当前字符是否在map中
	// 	_, ok := charMap[key]
	// 	if ok {
	// 		startIndex++
	// 		moveIndex = startIndex
	// 		templen := len(charMap)
	// 		if templen > maxLength {
	// 			maxLength = templen
	// 		}
	// 		charMap = make(map[string]string) // 清空 map
	// 	} else {
	// 		charMap[key] = key
	// 		templen := len(charMap)
	// 		if templen > maxLength {
	// 			maxLength = templen
	// 		}
	// 		moveIndex++
	// 	}

	// }

	// return maxLength
	//2、大牛解法
	// max函数获取多个元素中最大的值。例如：max(1,3,4,5) //输出5
	var index [128]int = [128]int{} // 记录每个字符最后一次出现的位置 +1
	var maxlen = 0
	var left = 0
	for right := 0; right < len(s); right++ {
		ch := s[right]                     //取出当前字符
		left = max(left, index[ch])        //获取left和index[ch]中最大的值
		maxlen = max(maxlen, right-left+1) //获取maxlen和right-left+1中最大的值
		index[ch] = right + 1              //更新index[ch]的值
	}
	return maxlen

}

//"abcabcbb"
//第一轮：right = 0, ch = a, left = 0, index[a] = 1, maxlen = 1
//第二轮：right = 1, ch = b, left = 0, index[b] = 2, maxlen = 2
//第三轮：right = 2, ch = c, left = 0, index[c] = 3, maxlen = 3
//第四轮：right = 3, ch = a, left = 0, index[a] = 1, maxlen = 3
//第五轮：right = 4, ch = b, left = 1, index[b] = 2, maxlen = 3
//第六轮：right = 5, ch = c, left = 2, index[c] = 3, maxlen = 3
