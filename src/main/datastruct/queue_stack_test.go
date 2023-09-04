package datastruct

type MyStack struct {
	in   []int
	out  []int
	size int
}

func ConstructorMyStack() MyStack {
	return MyStack{make([]int, 0), make([]int, 0), 0}
}

func (this *MyStack) Push(x int) {
	this.size++
	this.in = append(this.in, x)
}

func (this *MyStack) Pop() int {
	this.size--
	if len(this.in) > 0 {
		for i := 0; i < len(this.in); i++ {
			this.out = append(this.out, this.in[i])
		}
		this.in = []int{}
	}
	val := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return val
}

func (this *MyStack) Top() int {
	if len(this.in) > 0 {
		for i := 0; i < len(this.in); i++ {
			this.out = append(this.out, this.in[i])
		}
		this.in = []int{}
	}
	return this.out[len(this.out)-1]
}

func (this *MyStack) Empty() bool {
	return this.size == 0
}
