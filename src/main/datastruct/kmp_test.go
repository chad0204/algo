package datastruct

import (
	"fmt"
	"testing"
)

func TestKMP(t *testing.T) {
	fmt.Println(strStr("aabaaabaaac", "aabaaac"))
	fmt.Println(getNext("abababca"))
	fmt.Println(getNextV2("abababca"))

}

/*
demo:
aabaaabaaac/aabaaac
ababababca/abababca
mississippi/pi

原串T  ababababca
匹配串M	 abababca

计算匹配串的部分匹配表PMT

i = 1, j =0开始, 错开一位

0 1 2 3 4 5 6 7
a b a b a b c a

	a b a b a b c a

if M[i] != M[j], pmt[i] = 0, i++, j = 0

if M[i] == M[j], pmt[i] = pmt[i-1]+1, i++, j++

pmt = [?, 0, 1, 2, 3, 4, 0, 1]

求结果
i i i ----- i i   i i
a b a b a b a b c a

	a b a b a b c a

j j ------- j j

	j     j

?,0,1,2,3,4,0,1

			10
		      i
	      i

aabaaabaaac

	    aabaaac

		j
*/
func strStr(haystack string, needle string) int {
	m := len(haystack)
	n := len(needle)

	next := getNext(needle)
	i := 0
	j := 0
	for i < m && j < n {
		if haystack[i] == needle[j] {
			j++
			i++
		} else {
			if j > 0 {
				j = next[j-1]
			} else {
				i++
			}
		}
	}
	if j == n {
		return i - j
	}
	return -1
}

func getNext(needle string) []int {
	n := len(needle)
	i := 1
	j := 0
	next := make([]int, n)
	for i < n {
		if needle[i] == needle[j] {
			j++
			next[i] = j
			i++
		} else {
			if j > 0 {
				//如果j>0, 且needle[i] != needle[j], j回到next[j-1], 直到needle[i] == needle[j]或者遍历完字符串
				//此时i不用动
				j = next[j-1]
			} else {
				i++
			}
		}
	}
	return next
}

func getNextV2(needle string) []int {
	m := len(needle)
	next := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = next[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}
	return next
}
