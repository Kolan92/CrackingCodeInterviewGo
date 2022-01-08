package main

import "constraints"

type Node[T any] struct {
	value T
	next  *Node[T]
}

func removeDuplicates[T constraints.Ordered](head *Node[T]) {
	if head == nil {
		return
	}

	firstNode := head
	previousNode := head
	nextNode := head.next

	for firstNode != nil {
		for nextNode != nil {
			if firstNode.value == nextNode.value {
				if nextNode.next == nil {
					previousNode.next = nil
				} else {
					previousNode.next = nextNode.next
				}
				nextNode = previousNode.next
			} else {
				previousNode = previousNode.next
				nextNode = previousNode.next
			}
		}

		firstNode = firstNode.next
		previousNode = firstNode
		if previousNode == nil {
			nextNode = nil
		} else {
			nextNode = previousNode.next
		}
	}
}

func findLast[T constraints.Ordered](head *Node[T], indexFromEnd int) *Node[T] {

	count := 0

	countHead := head

	for countHead != nil {
		countHead = countHead.next
		count++
	}

	for i := 0; i < count-indexFromEnd-1; i++ {
		head = head.next
	}
	return head
}

func partition[T constraints.Ordered](head *Node[T], partitionValue T) *Node[T] {
	var smallerHead *Node[T]
	var smallerLast *Node[T]
	var biggerHead *Node[T]
	var biggerLast *Node[T]

	for head != nil {
		if head.value < partitionValue {
			if smallerHead == nil {
				smallerHead = &Node[T]{
					value: head.value,
				}
				smallerLast = smallerHead
			} else {
				smallerLast.next = &Node[T]{
					value: head.value,
				}
				smallerLast = smallerLast.next
			}
		} else {
			if biggerHead == nil {
				biggerHead = &Node[T]{
					value: head.value,
				}
				biggerLast = biggerHead
			} else {
				biggerLast.next = &Node[T]{
					value: head.value,
				}
				biggerLast = biggerLast.next
			}
		}

		head = head.next
	}

	if smallerLast != nil {
		smallerLast.next = biggerHead

		return smallerHead
	}
	return biggerHead
}

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

func sumReversed[T Numeric](numberA, numberB *Node[T]) *Node[T] {

	var sumHead *Node[T]
	var sumLast *Node[T]
	carry := false

	for numberA != nil || numberB != nil {
		var valueA T
		var valueB T

		if numberA != nil {
			valueA = numberA.value
		}
		if numberB != nil {
			valueB = numberB.value
		}

		sum, newCarry := calculateSum(valueA, valueB, carry)
		carry = newCarry

		if sumHead == nil {
			sumHead = &Node[T]{
				value: sum,
			}
			sumLast = sumHead
		} else {
			sumLast.next = &Node[T]{
				value: sum,
			}
			sumLast = sumLast.next
		}

		if numberA != nil {
			numberA = numberA.next
		}
		if numberB != nil {
			numberB = numberB.next
		}
	}

	return sumHead
}

func calculateSum[T Numeric](a, b T, previousCarry bool) (T, bool) {
	sum := a + b
	if previousCarry {
		sum++
	}

	carry := false

	if sum >= 10 {
		carry = true
		sum = sum - 10
	}

	return sum, carry
}

func sum[T Numeric](numberA, numberB *Node[T]) *Node[T] {

	reversedNumberA := reverseLinkedList(numberA)
	reversedNumberB := reverseLinkedList(numberB)
	reversedSum := sumReversed(reversedNumberA, reversedNumberB)

	return reverseLinkedList(reversedSum)
}

func reverseLinkedList[T constraints.Ordered](head *Node[T]) *Node[T] {
	var previous *Node[T]
	current := head
	following := head

	for current != nil {
		following = following.next
		current.next = previous
		previous = current
		current = following
	}

	return previous
}

func findLoop[T constraints.Ordered](head *Node[T]) *Node[T] {

	visitedNodes := make(map[*Node[T]]bool)

	current := head
	for current != nil {
		if _, hasNode := visitedNodes[current]; hasNode {
			return current
		}

		visitedNodes[current] = true
		current = current.next
	}
	return nil
}

func isPalindrom[T constraints.Ordered](head *Node[T]) bool {
	reversedList := reverseLinkedList(head)
	current := head
	for current != nil {
		if current.value != reversedList.value {
			return false
		}

		current = current.next
		reversedList = reversedList.next
	}
	return true
}
