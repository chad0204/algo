package main

func main() {

	decodeString("100[a]")

}

func decodeString(s string) string {
	/**
	         i
	  3[a12[cd]b]b


	  2,1

	  3[acdcdb]b
	*/

	n := len(s)
	stack := make([]byte, 0)
	for i := 0; i < n; i++ {
		if s[i] == ']' {
			//弹出']'
			back := make([]byte, 0)
			for j := len(stack) - 1; stack[j] != '['; j-- {
				back = append(back, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			//弹出'['
			stack = stack[:len(stack)-1]

			//弹出数字
			nums := make([]byte, 0)
			for j := len(stack) - 1; j >= 0 && '0' <= stack[j] && stack[j] <= '9'; j-- {
				nums = append(nums, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			num := 0
			pre := 1
			// 001
			for j := 0; j < len(nums); j++ {
				num = num + (int(nums[j])-48)*pre
				pre = pre * 10
			}
			//处理back, num
			for j := 1; j <= num; j++ {
				for b := len(back) - 1; b >= 0; b-- {
					stack = append(stack, back[b])
				}
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
