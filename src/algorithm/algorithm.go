package algorithm_learing

import (
	"fmt"
	"strings"
)

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
func sumOf2Num(target int, nums []int) (int, int, bool) {
	for i, size := 0, len(nums); i < size; i++ {
		iVal := nums[i]
		for j := 0; j < i; j++ {
			jVal := nums[j]
			if iVal+jVal == target {
				return i, j, true
			}
		}
	}
	return 0, 0, false
}

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//示例：
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
func sumOfListNum(nums1, nums2 []int) []int {
	size1, size2 := len(nums1), len(nums2)
	var longerList, shorterList []int
	var longerSize, shorterSize int
	if size1 >= size2 {
		longerList, shorterList = nums1, nums2
		longerSize, shorterSize = size1, size2
	} else {
		longerList, shorterList = nums2, nums1
		longerSize, shorterSize = size2, size1
	}
	result := make([]int, shorterSize, longerSize+1)
	carry := 0
	for i := 0; i < shorterSize; i++ {
		num1, num2 := longerList[i], shorterList[i]
		sum := num1 + num2 + carry
		//余数
		remainder := sum % 10
		//进位数
		carry = sum / 10
		result[i] = remainder
	}
	if carry > 0 {
		for i := shorterSize; i < longerSize; i++ {
			if carry > 0 {
				//余数
				sum := longerList[i] + carry
				remainder := sum % 10
				//进位数
				carry = sum / 10
				result = append(result, remainder)
			} else {
				result = append(result, longerList[i])
			}
		}
		if carry > 0 {
			result = append(result, carry)
		}
	} else {
		for i := shorterSize; i < longerSize; i++ {
			result = append(result, longerList[i])
		}
	}
	return result
}

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//示例 1:
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
func longestUncomplicatableSubStr(source string) string {
	unComplicatedSubStr := ""
	var subStr string
	for j, size := 0, len(source); j < size; j++ {
		jVal := source[j]
		if indexRune := strings.IndexRune(subStr, rune(jVal)); indexRune == -1 {
			subStr += string(jVal)
			if len(subStr) > len(unComplicatedSubStr) {
				unComplicatedSubStr = subStr
			}
		} else {
			if indexRune > (len(subStr)+1)/2 {
				if len(unComplicatedSubStr) < indexRune+1 {
					unComplicatedSubStr = subStr[:indexRune+1]
				}
			}
			if indexRune > len(subStr)-1 {
				subStr = string(jVal)
			} else {
				subStr = subStr[indexRune+1:] + string(jVal)
			}
		}
	}
	return unComplicatedSubStr
}

//二分查找
func BinarySearch(arr []int, target int) (index int) {
	index = -1
	fromIdx, toIdx := 0, len(arr)
	for idx := (fromIdx + toIdx) / 2; fromIdx < toIdx; {
		idxVal := arr[idx]
		if idxVal == target {
			index = idx
			return
		} else if idxVal < target {
			fromIdx = idx + 1
		} else {
			toIdx = idx
		}
	}
	return
}

//https://github.com/MisterBooo/LeetCodeAnimation/blob/master/notes/LeetCode%E7%AC%AC131%E5%8F%B7%E9%97%AE%E9%A2%98%EF%BC%9A%E5%88%86%E5%89%B2%E5%9B%9E%E6%96%87%E4%B8%B2.md
// 给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
//返回 s 所有可能的分割方案。
//所谓回文，就是一个正读和反读都一样的字符串。
//题目解析
//首先，对于一个字符串的分割，肯定需要将所有分割情况都遍历完毕才能判断是不是回文数。不能因为 abba 是回文串，就认为它的所有子串都是回文的。
//既然需要将所有的分割方法都找出来，那么肯定需要用到DFS（深度优先搜索）或者BFS（广度优先搜索）。
//在分割的过程中对于每一个字符串而言都可以分为两部分：左边一个回文串加右边一个子串，比如 "abc" 可分为 "a" + "bc" 。 然后对"bc"分割仍然是同样的方法，分为"b"+"c"。
//在处理的时候去优先寻找更短的回文串，然后回溯找稍微长一些的回文串分割方法，不断回溯，分割，直到找到所有的分割方法。

//https://github.com/MisterBooo/LeetCodeAnimation/blob/master/notes/LeetCode第139号问题：单词拆分.md
//给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
//说明：
//拆分时可以重复使用字典中的单词。
//你可以假设字典中没有重复的单词。

//https://mp.weixin.qq.com/s/xHxjCDdFZyCW2pnY6Cz8SQ
//给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。如果数组元素个数小于 2，则返回 0。
//这里需要用到的是不经常使用的一种排序方法 —— 桶排序！

//请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。
//给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。
//保证字符串中的字符为【ASCII字符】。字符串的长度小于等于【3000】。
func isNoComplicatedCharStr(source string) bool {
	if len([]rune(source)) > 3000 {
		return false
	}
	for _, v := range source {
		if strings.Count(source, string(v)) > 1 {
			return false
		}
	}
	return true
}

//请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。
//给定一个string，请返回一个string，为翻转后的字符串。
func reverseStr(source string) string {
	runes := []rune(source)
	size := len(runes)
	for i, loopTimes := 0, size/2; i < loopTimes; i++ {
		runes[i], runes[size-i] = runes[size-i], runes[i]
	}
	return string(runes)
}

//深度优先算法
const count = 3

var a, bookD [count]int

func dfs(step int) {
	if step == count {
		fmt.Println(a)
		return
	}
	for i := 0; i < count; i++ {
		if bookD[i] == 0 {
			a[step] = i + 1 //用牌
			bookD[i] = 1    //占用标记
			dfs(step + 1)   //下探
			bookD[i] = 0    //归还
		}
	}
}

//广度优先算法
const size = 5

var head, tail int
var bookB [size]int
var q [size * size]int
var graph [size][size]int

func bfs() {
	q[tail] = 0
	bookB[tail] = 1
	tail++
	//head = tail
	for head < tail {
		cur := q[head]
		for i := 0; i < size; i++ {
			if graph[cur][i] == 1 && bookB[i] == 0 {
				fmt.Println(cur, "->", i)
				q[tail] = i
				tail++
				bookB[i] = 1
			} else {
				fmt.Println(cur, "-/>", i)
			}

			if tail > size {
				break
			}
		}
		fmt.Println(bookB)
		head++
	}
	fmt.Println(q)
}
