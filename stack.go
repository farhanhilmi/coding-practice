package main

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}

	n := len(*s)

	top := (*s)[n-1]
	*s = (*s)[:n-1]

	return top
}

// func main() {
// 	st := Stack{}
// 	st.Push(1)
// 	st.Push(2)
// 	fmt.Println("The value popped from the stack is given as:")
// 	fmt.Println(st.Pop())
// 	fmt.Println(st.Pop())
// 	fmt.Println("Is the stack empty?")
// 	fmt.Println(st.IsEmpty())
// }
