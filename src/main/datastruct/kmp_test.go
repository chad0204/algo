package datastruct

import (
	"fmt"
	"testing"
)

func TestKMP(t *testing.T) {
	fmt.Println(strStr("aabaaabaaac", "aabaaac"))
	fmt.Println(getNext("abababca"))
}

/*

参考:
https://www.zhihu.com/question/21923021/answer/281346746
这篇例子没讲清楚匹配串的后退方式, 也就是j>0时, j需要根据已有的pmt进行后退, j = pmt[j-1]

demo:
aabaaabaaac/aabaaac
ababababca/abababca
mississippi/pi

原串T  aabaaabaaac
匹配串M	 aabaaac



0. 什么是pmt?

PMT中的值是字符串的前缀集合与后缀集合的交集中最长元素的长度
比如aabaaac
i = 0, a , 无, pmt[0] = 0
i = 1, aa, {a} {a}, 无, pmt[1] = 1
i = 2, aab, {a, aa} {ab, b}, a == a, pmt[2] = 0
i = 3, aaba, {a, aa, aab} {aba, ba, a}, pmt[3] = 1
i = 4, aabaa, {a, aa, aab, aaba} {abaa, baa, aa, a}, pmt[4] = 2
i = 5, aabaaa, {a, aa, aab, aaba, aabaa} {abaaa, baaa, aaa, aa, a}, pmt[5] = 2
i = 6, aabaaac, {a, aa, aab, aaba, aabaa, aabaaa} {abaaac, baaac, aaac, aac, ac, c}, pmt[5] = 0


1. 计算匹配串的部分匹配表PMT

i = 1, j =0开始, 错开一位, 找到共同前缀和后缀, 优化后退步数

0 1 2 3 4 5 6
  i   i i i i
a a b a a a c
  j j   j j j
            a a b a a a c

0 1 0 1 2 2 0


if M[i] == M[j], j++, pmt[i] = j, i++

if M[i] != M[j], 要么j指针往后退, 要么i指针往前进。如果 j > 0, j指针后退方式是j = pmt[j-1]; 如果j == 0, j指针退无可退, i往前移动。直到if M[i] == M[j]为止

理解j = pmt[j-1]是关键, 通过前缀和后缀


2. 通过pmt[0, 1, 0, 1, 2, 2, 0]求结果, 思路和求pmt类似

0 1 0 1 2 2 0   pmt

0 1 2 3 4 5 6 7 8 9 10
  i         i       i i
a a b a a a b a a a c
        a a b a a a c
            j       j j

T[i] == M[j], i++, j++
当T[6] != M[6], j后退, j = pmt[j-1], j = 2, 继续比较T[6]和M[2]

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
				//这里就是j == 0的情况, 其实可以并到needle[i] == needle[j], 简化代码
				i++
			}
		}
	}
	return next
}
