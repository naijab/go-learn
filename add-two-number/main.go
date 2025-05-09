package main

// ListNode represents a node in a singly linked list.
type ListNode struct {
	Val  int       // Value of the node
	Next *ListNode // Pointer to the next node
}

func main() {
	// Example usage:
	// Create the first linked list: 2 -> 4 -> 3
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	// Create the second linked list: 5 -> 6 -> 4
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	// Add the two numbers represented by the linked lists
	result := addTwoNumbers(l1, l2)

	// Print the result linked list
	for result != nil {
		print(result.Val)    // Print the value of the current node
		result = result.Next // Move to the next node
	}
}

// addTwoNumbers adds two numbers represented by linked lists and returns the result as a linked list.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Pointers to traverse the input linked lists
	l1p := l1
	l2p := l2

	// Dummy node to simplify result list creation
	dummy := &ListNode{}

	// Pointer to build the result linked list
	output := dummy

	// Variable to store the carry value
	carry := 0

	// Loop until both linked lists are fully traversed
	for l1p != nil || l2p != nil {
		// Start with the carry from the previous step
		sum := carry

		// Add the value from the first linked list, if available
		if l1p != nil {
			sum += l1p.Val
			l1p = l1p.Next // Move to the next node in l1
		}

		// Add the value from the second linked list, if available
		if l2p != nil {
			sum += l2p.Val
			l2p = l2p.Next // Move to the next node in l2
		}

		// Update the carry for the next step
		carry = sum / 10

		// Create a new node with the current digit and add it to the result list
		output.Next = &ListNode{Val: sum % 10}
		output = output.Next // Move to the next node in the result list
	}

	// If there's a carry left, add it as a new node
	if carry > 0 {
		output.Next = &ListNode{Val: carry}
	}

	// Return the result list, skipping the dummy node
	return dummy.Next
}
