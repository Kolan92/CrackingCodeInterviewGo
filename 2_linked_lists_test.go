package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	t.Run("List without duplicates", func(t *testing.T) {
		list := Node[int] {
			value: 20,
			next: &Node[int] {
				value: 22,
				next: nil,
			},
		}

		removeDuplicates(&list)

		assert.NotNil(t, list.next)
	}) 

	t.Run("List with duplicates in the end", func(t *testing.T) {
		list := Node[int] {
			value: 20,
			next: &Node[int] {
				value: 20,
				next: nil,
			},
		}

		removeDuplicates(&list)

		assert.Nil(t, list.next)
	}) 

	t.Run("List with duplicates in the middle", func(t *testing.T) {
		list := Node[int] {
			value: 20,
			next: &Node[int] {
				value: 20,
				next: &Node[int] {
					value: 21,
					next: nil,
				},
			},
		}

		removeDuplicates(&list)

		assert.NotNil(t, list.next)
		assert.Equal(t, 21, list.next.value)
	})

	t.Run("Removes not consecutive duplicates", func(t *testing.T) {
		list := Node[int] {
			value: 20,
			next: &Node[int] {
				value: 21,
				next: &Node[int] {
					value: 20,
					next: nil,
				},
			},
		}

		removeDuplicates(&list)

		assert.Nil(t, list.next.next)
	})

	t.Run("Removes multiple duplicates", func(t *testing.T) {
		list := Node[int] {
			value: 20,
			next: &Node[int] {
				value: 20,
				next: &Node[int] {
					value: 20,
					next: &Node[int] {
						value: 20,
						next: nil,
					},
				},
			},
		}

		removeDuplicates(&list)

		assert.Nil(t, list.next)
	})

	t.Run("Removes multiple different duplicates", func(t *testing.T) {
		list := Node[int] {
			value: 20,
			next: &Node[int] {
				value: 21,
				next: &Node[int] {
					value: 20,
					next: &Node[int] {
						value: 21,
						next: &Node[int] {
							value: 23,
							next: nil,
						},
					},
				},
			},
		}

		removeDuplicates(&list)

		assert.Equal(t, 21, list.next.value)
		assert.Equal(t, 23, list.next.next.value)
		assert.Nil(t, list.next.next.next)
	})
}


func TestFindLast(t *testing.T) {
	t.Run("Should return first node given one node list", func(t *testing.T)  {
		list := Node[int] {
			value: 20,
			next:  nil,
		}

		node := findLast(&list, 1)

		assert.Equal(t, list.value, node.value)
	})

	t.Run("Should return 3rd last node", func(t *testing.T)  {
		list := Node[int] {
			value: 1,
			next: &Node[int] {
				value: 2,
				next: &Node[int] {
					value: 3,
					next: &Node[int] {
						value: 4,
						next: &Node[int] {
							value: 5,
							next: nil,
						},
					},
				},
			},
		}


		node := findLast(&list, 2)

		assert.Equal(t, 3, node.value)
	})

	t.Run("Should return last node given long list", func(t *testing.T)  {
		list := Node[int] {
			value: 1,
			next: &Node[int] {
				value: 2,
				next: &Node[int] {
					value: 3,
					next: &Node[int] {
						value: 4,
						next: &Node[int] {
							value: 5,
							next: nil,
						},
					},
				},
			},
		}


		node := findLast(&list, 0)

		assert.Equal(t, 5, node.value)
	})
}

func TestPartition(t *testing.T) {
	
	t.Run("One element list", func(t *testing.T)  {
		list := Node[int] {
			value: 1,
			next: nil,
		}

		partitionedList := partition(&list, 2)

		expectedList := &Node[int] {
			value: 1,
			next: nil,
		}

		assert.Equal(t, expectedList, partitionedList)
	})

	t.Run("List with multiple unordered elements", func(t *testing.T)  {
		list := Node[int] {
			value: 5,
			next: &Node[int] {
				value: 2,
				next: &Node[int] {
					value: 8,
					next: &Node[int] {
						value: 1,
						next: &Node[int] {
							value: 3,
							next: nil,
						},
					},
				},
			},
		}

		partitionedList := partition(&list, 3)

		expectedList := &Node[int] {
			value: 2,
			next: &Node[int] {
				value: 1,
				next: &Node[int] {
					value: 5,
					next: &Node[int] {
						value: 8,
						next: &Node[int] {
							value: 3,
							next: nil,
						},
					},
				},
			},
		}

		assert.Equal(t, expectedList, partitionedList)
	})

	t.Run("List all bigger elements", func(t *testing.T)  {
		list := Node[int] {
			value: 5,
			next: &Node[int] {
				value: 2,
				next: &Node[int] {
					value: 8,
					next: &Node[int] {
						value: 1,
						next: &Node[int] {
							value: 3,
							next: nil,
						},
					},
				},
			},
		}

		partitionedList := partition(&list, 0)

		expectedList := &Node[int] {
			value: 5,
			next: &Node[int] {
				value: 2,
				next: &Node[int] {
					value: 8,
					next: &Node[int] {
						value: 1,
						next: &Node[int] {
							value: 3,
							next: nil,
						},
					},
				},
			},
		}

		assert.Equal(t, expectedList, partitionedList)
	})
}

func TestSumReversed(t *testing.T) {
	t.Run("617+295=912", func(t *testing.T){
		numberA := Node[int] {
			value: 7,
			next: &Node[int] {
				value: 1,
				next: &Node[int] {
					value: 6,
					next: nil,
				},
			},
		}

		numberB := Node[int] {
			value: 5,
			next: &Node[int] {
				value: 9,
				next: &Node[int] {
					value: 2,
					next: nil,
				},
			},
		}

		actualSum := sumReversed(&numberA, &numberB)

		expectedSum := &Node[int] {
			value: 2,
			next: &Node[int] {
				value: 1,
				next: &Node[int] {
					value: 9,
					next: nil,
				},
			},
		}

		assert.Equal(t, expectedSum, actualSum)
	})

	t.Run("617+28=645", func(t *testing.T){
		numberA := Node[int] {
			value: 7,
			next: &Node[int] {
				value: 1,
				next: &Node[int] {
					value: 6,
					next: nil,
				},
			},
		}

		numberB := Node[int] {
			value: 8,
			next: &Node[int] {
				value: 2,
				next: nil,
			},
		}

		actualSum := sumReversed(&numberA, &numberB)

		expectedSum := &Node[int] {
			value: 5,
			next: &Node[int] {
				value: 4,
				next: &Node[int] {
					value: 6,
					next: nil,
				},
			},
		}

		assert.Equal(t, expectedSum, actualSum)
	})
}


func TestSum(t *testing.T) {
	t.Run("617+295=912", func(t *testing.T){
		numberA := Node[int] {
			value: 6,
			next: &Node[int] {
				value: 1,
				next: &Node[int] {
					value: 7,
					next: nil,
				},
			},
		}

		numberB := Node[int] {
			value: 2,
			next: &Node[int] {
				value: 9,
				next: &Node[int] {
					value: 5,
					next: nil,
				},
			},
		}

		actualSum := sum(&numberA, &numberB)

		expectedSum := &Node[int] {
			value: 9,
			next: &Node[int] {
				value: 1,
				next: &Node[int] {
					value: 2,
					next: nil,
				},
			},
		}

		assert.Equal(t, expectedSum, actualSum)
	})

	t.Run("617+28=645", func(t *testing.T){
		numberA := Node[int] {
			value: 6,
			next: &Node[int] {
				value: 1,
				next: &Node[int] {
					value: 7,
					next: nil,
				},
			},
		}

		numberB := Node[int] {
			value: 2,
			next: &Node[int] {
				value: 8,
				next: nil,
			},
		}

		actualSum := sum(&numberA, &numberB)

		expectedSum := &Node[int] {
			value: 6,
			next: &Node[int] {
				value: 4,
				next: &Node[int] {
					value: 5,
					next: nil,
				},
			},
		}

		assert.Equal(t, expectedSum, actualSum)
	})
}


func TestReverseLinkedList(t *testing.T) {

	list := Node[int] {
		value: 5,
		next: &Node[int] {
			value: 2,
			next: &Node[int] {
				value: 8,
				next: &Node[int] {
					value: 1,
					next: &Node[int] {
						value: 3,
						next: nil,
					},
				},
			},
		},
	}

	reversedList := reverseLinkedList(&list)

	expectedList := Node[int] {
		value: 3,
		next: &Node[int] {
			value: 1,
			next: &Node[int] {
				value: 8,
				next: &Node[int] {
					value: 2,
					next: &Node[int] {
						value: 5,
						next: nil,
					},
				},
			},
		},
	}

	assert.Equal(t, &expectedList, reversedList)
}

func TestFindLoop(t *testing.T)  {

	t.Run("Without loop", func (t *testing.T)  {	
		list := Node[int] {
			value: 1,
			next: &Node[int] {
				value: 2,
				next: &Node[int] {
					value: 3,
					next: &Node[int] {
						value: 4,
						next: &Node[int] {
							value: 5,
							next: nil,
						},
					},
				},
			},
		}
		
		foundNodeBeginingLoop := findLoop(&list)
		assert.Nil(t, foundNodeBeginingLoop)
	})

	t.Run("With loop", func (t *testing.T)  {	
		list := Node[int] {
			value: 1,
			next: &Node[int] {
				value: 2,
				next: &Node[int] {
					value: 3,
					next: &Node[int] {
						value: 4,
						next: &Node[int] {
							value: 5,
							next: nil,
						},
					},
				},
			},
		}

		nodeAtBeginingOfLoop := list.next.next

		list.next.next.next.next.next = nodeAtBeginingOfLoop
		
		foundNodeBeginingLoop := findLoop(&list)

		assert.Equal(t, nodeAtBeginingOfLoop, foundNodeBeginingLoop)
	})
}

func TestIsPalindrom(t *testing.T)  {

	t.Run("Not palindrom", func(t *testing.T)  {
		list := Node[rune] {
			value: 'K',
			next: &Node[rune] {
				value: 'A',
				next: &Node[rune] {
					value: 'Y',
					next: &Node[rune] {
						value: 'A',
						next: &Node[rune] {
							value: 'Z',
							next: nil,
						},
					},
				},
			},
		}

		isPalindrom := isPalindrom(&list)

		assert.False(t, isPalindrom)
	})

	t.Run("Palindrom", func(t *testing.T)  {
		list := Node[rune] {
			value: 'K',
			next: &Node[rune] {
				value: 'A',
				next: &Node[rune] {
					value: 'Y',
					next: &Node[rune] {
						value: 'A',
						next: &Node[rune] {
							value: 'K',
							next: nil,
						},
					},
				},
			},
		}

		isPalindrom := isPalindrom(&list)

		assert.True(t, isPalindrom)
	})
	
}