package datastruct

import (
	"fmt"
	"math"
	"testing"
)

func TestMinWindow(t *testing.T) {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

// 剑指 Offer II 017. 含有所有字符的最短字符串; 76. 最小覆盖子串
func minWindow(s string, t string) string {
	needs := make(map[byte]int, len(t))
	for i := range t {
		//可能有重复值, 值不一定为1
		needs[t[i]] = needs[t[i]] + 1
	}
	window := make(map[byte]int)
	left, right, checkNum := 0, 0, 0
	start := 0
	l := math.MaxInt32

	for right < len(s) {
		add := s[right]
		right++
		if _, ok := needs[add]; ok {
			//先加后判断
			window[add] = window[add] + 1
			if window[add] == needs[add] {
				//如果一个字符够了, 计数
				checkNum++
			}
		}

		for checkNum == len(needs) {
			if l > (right - left) {
				l = right - left
				start = left
			}
			rmv := s[left]
			left++
			if _, ok := needs[rmv]; ok {
				//先判断后减
				if window[rmv] == needs[rmv] {
					//后面的操作会让一个不满足, 减掉
					checkNum--
				}
				window[rmv] = window[rmv] - 1
			}
		}
	}
	if l == math.MaxInt32 {
		return ""
	}
	return s[start : start+l]
}
