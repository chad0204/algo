package datastruct

import (
	"fmt"
	"math"
	"testing"
)

/**

先进窗口, 后移动游标

*/

func TestMinWindow(t *testing.T) {
	fmt.Println(minWindow("a", "a"))
}

// 剑指 Offer II 017. 含有所有字符的最短字符串; 76. 最小覆盖子串
func minWindow(s string, t string) string {
	needs := make(map[byte]int, 0)
	for i := range t {
		//可能有重复值, 值不一定为1
		needs[t[i]] = needs[t[i]] + 1
	}
	window := make(map[byte]int)
	left, right, checkNum := 0, 0, 0

	/**
	l记录从start开始满足条件的字符串长度
	start记录最小一次满足条件的起始位置, 因为后续的left的right还可能会满足, 会一直更新, 所以需要记录下此时的位置。只有当新的left的right比l小的时候, start才会更新成left。
	*/
	start := 0
	l := math.MaxInt32 // 不建议搞这么多变量, 定义个res := s + t就可以了

	for right < len(s) {
		add := s[right]
		if _, ok := needs[add]; ok {
			//先加后判断
			window[add] = window[add] + 1
			if window[add] == needs[add] {
				//如果一个字符够了, 计数
				checkNum++
			}
		}
		right++

		for checkNum == len(needs) {
			if l > (right - left) {
				l = right - left
				start = left
			}
			rmv := s[left]
			if _, ok := needs[rmv]; ok {
				//先判断后减
				if window[rmv] == needs[rmv] {
					//后面的操作会让一个不满足, 减掉
					checkNum--
				}
				window[rmv] = window[rmv] - 1
			}
			left++
		}
	}
	if l == math.MaxInt32 { // res == s +t
		return ""
	}
	return s[start : start+l]
}

// 567. 字符串的排列; 剑指 Offer II 014. 字符串中的变位词; 这题和438一样 就是返回值类型不一样
func checkInclusion(s1 string, s2 string) bool {
	needs := make(map[byte]int)
	for i := range s1 {
		needs[s1[i]] = needs[s1[i]] + 1
	}
	window := make(map[byte]int)
	left, right := 0, 0
	checkSum := 0

	for right < len(s2) {
		add := s2[right]
		if _, ok := needs[add]; ok {
			window[add] = window[add] + 1
			if window[add] == needs[add] {
				checkSum++
			}
		}
		right++
		for right-left == len(s1) {
			if checkSum == len(needs) {
				return true
			}
			rmv := s2[left]
			if _, ok := needs[rmv]; ok {
				if window[rmv] == needs[rmv] {
					checkSum--
				}
				window[rmv] = window[rmv] - 1
			}
			left++
		}
	}
	return false
}

func TestFindAnagrams(t *testing.T) {
	ints := make([]int, 0)
	fmt.Println(ints)
}

// 438. 找到字符串中所有字母异位词; 剑指 Offer II 015. 字符串中的所有变位词; 这题和438一样 就是返回值类型不一样
func findAnagrams(s string, p string) []int {

	window := make(map[byte]int)
	needs := make(map[byte]int)
	for i := range p {
		needs[p[i]] = needs[p[i]] + 1
	}

	left, right := 0, 0
	res := make([]int, 0)

	checkSum := 0

	for right < len(s) {
		add := s[right]
		if _, ok := needs[add]; ok {
			window[add] = window[add] + 1
			if window[add] == needs[add] {
				checkSum++
			}
		}
		right++

		for right-left == len(p) {
			if checkSum == len(needs) {
				res = append(res, left)
			}
			rmv := s[left]
			if _, ok := needs[rmv]; ok {
				if window[rmv] == needs[rmv] {
					checkSum--
				}
				window[rmv] = window[rmv] - 1
			}
			left++
		}
	}
	return res
}

func TestLOLS(t *testing.T) {
	lengthOfLongestSubstring("abcdef")
}

// 3. 无重复字符的最长子串; 剑指 Offer 48. 最长不含重复字符的子字符串 剑指 Offer II 016. 不含重复字符的最长子字符串
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	l := 0
	left, right := 0, 0
	for right < len(s) {
		add := s[right]
		window[add] = window[add] + 1
		right++

		//出现重复项
		for window[add] > 1 {
			rmv := s[left]
			window[rmv] = window[rmv] - 1
			left++
		}

		if right-left > l {
			l = right - left
		}
	}
	return l
}
