package main

type Node struct {
	Data int
	Next *Node
}

func findLastNthNode(head *Node, n int) *Node {
	if head == nil || n <= 0 {
		return nil
	}

	fast := head
	slow := head

	for i := 0; i < n; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next

	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}

// func main() {
// 	// Example usage
// 	// Creating a linked list: 1 -> 2 -> 3 -> 4 -> 5
// 	head := &Node{Data: 1, Next: &Node{Data: 2, Next: &Node{Data: 3, Next: &Node{Data: 4, Next: &Node{Data: 5}}}}}

// 	n := 2
// 	result := findLastNthNode(head, n)

// 	if result != nil {
// 		fmt.Printf("The last %dth node is: %d\n", n, result.Data)
// 	} else {
// 		fmt.Printf("Invalid input or the list is shorter than %d\n", n)
// 	}
// }
