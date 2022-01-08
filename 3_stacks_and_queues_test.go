package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Pick returns nil given empty stack", func(t *testing.T) {
		stack := newStack[int]()
		assert.Nil(t, stack.Pick())
	})

	t.Run("Pick returns last value given stack with value", func(t *testing.T) {
		stack := newStack[int]()
		stack.Push(44)
		assert.Equal(t, 44, *stack.Pick())
	})

	t.Run("Pop returns last value given stack with value", func(t *testing.T) {
		stack := newStack[int]()
		stack.Push(44)
		assert.Equal(t, 44, *stack.Pop())
	})

	t.Run("Pop returns last value given stack with multiple values", func(t *testing.T) {
		stack := newStack[int]()
		stack.Push(44)
		stack.Push(12)
		assert.Equal(t, 12, *stack.Pop())
		assert.Equal(t, 44, *stack.Pop())
		assert.Nil(t, stack.Pop())
	})
}

func TestQueue(t *testing.T) {
	t.Run("Pick returns nil given empty queue", func(t *testing.T) {
		queue := newQueue[int]()
		assert.Nil(t, queue.Pick())
	})

	t.Run("Pick returns first value given queue with value", func(t *testing.T) {
		queue := newQueue[int]()
		queue.Enqueue(44)
		assert.Equal(t, 44, *queue.Pick())
	})

	t.Run("Dequeue returns first value given queue with value", func(t *testing.T) {
		queue := newQueue[int]()
		queue.Enqueue(44)
		assert.Equal(t, 44, *queue.Dequeue())
	})

	t.Run("Dequeue returns first value given queue with multiple values", func(t *testing.T) {
		queue := newQueue[int]()
		queue.Enqueue(44)
		queue.Enqueue(12)
		queue.Enqueue(14)
		assert.Equal(t, 44, *queue.Dequeue())
		assert.Equal(t, 12, *queue.Dequeue())
		assert.Equal(t, 14, *queue.Dequeue())
		assert.Nil(t, queue.Dequeue())
	})
}

func TestMinStack(t *testing.T) {
	t.Run("Return nil for min given empty stack", func(t *testing.T) {
		stack := newMinStack[int]()
		assert.Nil(t, stack.Min())
	})

	t.Run("Min returns minimum value given stack with multiple values", func(t *testing.T) {
		stack := newMinStack[int]()
		stack.Push(12)
		stack.Push(44)
		assert.Equal(t, 12, *stack.Min())
	})

	t.Run("Min returns minimum value given after poping up minimum value", func(t *testing.T) {
		stack := newMinStack[int]()
		stack.Push(15)
		stack.Push(44)
		stack.Push(12)
		_ = stack.Pop()
		assert.Equal(t, 15, *stack.Min())
	})
}

func TestHanoiGame(t *testing.T) {

	t.Run("Error when creating 0 elements tower", func(t *testing.T) {
		game, err := newHanoiGame(0)
		assert.Nil(t, game)
		assert.NotNil(t, err)
	})

	t.Run("Check initial conditions", func(t *testing.T) {
		game, err := newHanoiGame(5)
		assert.Nil(t, err)
		assert.Equal(t, 1, *game.sourceRod.Pop())
		assert.Equal(t, 2, *game.sourceRod.Pop())
		assert.Equal(t, 3, *game.sourceRod.Pop())
		assert.Equal(t, 4, *game.sourceRod.Pop())
		assert.Equal(t, 5, *game.sourceRod.Pop())
		assert.True(t, game.spareRod.IsEmpty())
		assert.True(t, game.destinationRod.IsEmpty())
	})

	t.Run("Solves game with 1 disk", func(t *testing.T) {
		game, _ := newHanoiGame(1)
		game.Solve()

		assert.True(t, game.sourceRod.IsEmpty())
		assert.True(t, game.spareRod.IsEmpty())
		assert.Equal(t, 1, *game.destinationRod.Pop())
	})

	t.Run("Solves game with 2 disk", func(t *testing.T) {
		game, _ := newHanoiGame(2)
		game.Solve()

		assert.True(t, game.sourceRod.IsEmpty())
		assert.True(t, game.spareRod.IsEmpty())
		assert.Equal(t, 1, *game.destinationRod.Pop())
		assert.Equal(t, 2, *game.destinationRod.Pop())
	})

	t.Run("Solves game with 3 disk", func(t *testing.T) {
		game, _ := newHanoiGame(3)
		game.Solve()

		assert.True(t, game.sourceRod.IsEmpty())
		assert.True(t, game.spareRod.IsEmpty())
		assert.Equal(t, 1, *game.destinationRod.Pop())
		assert.Equal(t, 2, *game.destinationRod.Pop())
		assert.Equal(t, 3, *game.destinationRod.Pop())
	})

	t.Run("Solves game with 3 disk", func(t *testing.T) {
		game, _ := newHanoiGame(6)
		game.Solve()

		assert.True(t, game.sourceRod.IsEmpty())
		assert.True(t, game.spareRod.IsEmpty())
		assert.Equal(t, 1, *game.destinationRod.Pop())
		assert.Equal(t, 2, *game.destinationRod.Pop())
		assert.Equal(t, 3, *game.destinationRod.Pop())
		assert.Equal(t, 4, *game.destinationRod.Pop())
		assert.Equal(t, 5, *game.destinationRod.Pop())
		assert.Equal(t, 6, *game.destinationRod.Pop())
	})
}

func TestQueueOnTwoStacks(t *testing.T) {
	t.Run("Pick returns nil given empty queue", func(t *testing.T) {
		queue := newQueueOnTwoStacks[int]()
		assert.Nil(t, queue.Pick())
	})

	t.Run("Pick returns first value given queue with value", func(t *testing.T) {
		queue := newQueueOnTwoStacks[int]()
		queue.Enqueue(44)
		assert.Equal(t, 44, *queue.Pick())
	})

	t.Run("Dequeue returns first value given queue with value", func(t *testing.T) {
		queue := newQueueOnTwoStacks[int]()
		queue.Enqueue(44)
		assert.Equal(t, 44, *queue.Dequeue())
	})

	t.Run("Dequeue returns first value given queue with multiple values", func(t *testing.T) {
		queue := newQueueOnTwoStacks[int]()
		queue.Enqueue(44)
		queue.Enqueue(12)
		assert.Equal(t, 44, *queue.Dequeue())
		assert.Equal(t, 12, *queue.Dequeue())
		assert.Nil(t, queue.Dequeue())
	})
}

func TestMaxStack(t *testing.T) {
	t.Run("Pick returns nil given empty stack", func(t *testing.T) {
		stack := newMaxStack[int]()
		assert.Nil(t, stack.Pick())
	})

	t.Run("Pick returns last value given stack with value", func(t *testing.T) {
		stack := newMaxStack[int]()
		stack.Push(44)
		assert.Equal(t, 44, *stack.Pick())
	})

	t.Run("Pop returns last value given stack with value", func(t *testing.T) {
		stack := newMaxStack[int]()
		stack.Push(44)
		assert.Equal(t, 44, *stack.Pop())
	})

	t.Run("Pop returns max value given stack with two values", func(t *testing.T) {
		stack := newMaxStack[int]()
		stack.Push(108)
		stack.Push(44)
		assert.Equal(t, 108, *stack.Pick())
		assert.Equal(t, 108, *stack.Pop())

		assert.Equal(t, 44, *stack.Pick())
		assert.Equal(t, 44, *stack.Pop())

		assert.Nil(t, stack.Pick())
		assert.Nil(t, stack.Pop())
	})

	t.Run("Pop returns max value given stack with multiple values", func(t *testing.T) {
		stack := newMaxStack[int]()
		stack.Push(44)
		stack.Push(15)
		stack.Push(108)
		stack.Push(55)
		stack.Push(12)
		assert.Equal(t, 108, *stack.Pick())
		assert.Equal(t, 108, *stack.Pop())

		assert.Equal(t, 55, *stack.Pick())
		assert.Equal(t, 55, *stack.Pop())

		assert.Equal(t, 44, *stack.Pick())
		assert.Equal(t, 44, *stack.Pop())

		assert.Equal(t, 15, *stack.Pick())
		assert.Equal(t, 15, *stack.Pop())

		assert.Equal(t, 12, *stack.Pick())
		assert.Equal(t, 12, *stack.Pop())

		assert.Nil(t, stack.Pick())
		assert.Nil(t, stack.Pop())
	})
}

func TestShelter(t *testing.T) {
	t.Run("Pick returns nil given empty shelter", func(t *testing.T) {
		shelter := newShelter()
		assert.Nil(t, shelter.Pick())
	})

	t.Run("Pick returns first animal given shelter with one cat", func(t *testing.T) {
		shelter := newShelter()
		cat := newCat("Cat1")
		shelter.Enqueue(cat)
		assert.Equal(t, "Cat1", shelter.Pick().GetName())
	})

	t.Run("DequeueAny returns first animal given shelter with one cat", func(t *testing.T) {
		shelter := newShelter()
		cat := newCat("Cat1")
		shelter.Enqueue(cat)
		assert.Equal(t, "Cat1", shelter.DequeueAny().GetName())
	})

	t.Run("Pick returns first animal given shelter with one dog", func(t *testing.T) {
		shelter := newShelter()
		dog := newDog("Dog1")
		shelter.Enqueue(dog)
		assert.Equal(t, "Dog1", shelter.Pick().GetName())
	})

	t.Run("DequeueAny returns first animal given shelter with one dog", func(t *testing.T) {
		shelter := newShelter()
		dog := newDog("Dog1")
		shelter.Enqueue(dog)
		assert.Equal(t, "Dog1", shelter.DequeueAny().GetName())
	})

	t.Run("DequeueDog returns nil given shelter with only cats", func(t *testing.T) {
		shelter := newShelter()
		cat := newCat("Cat1")
		shelter.Enqueue(cat)
		assert.Nil(t, shelter.DequeueDog())
	})

	t.Run("DequeueDog returns first dog given shelter with multiple animals", func(t *testing.T) {
		shelter := newShelter()
		cat1 := newCat("Cat1")
		shelter.Enqueue(cat1)
		dog := newDog("Dog1")
		shelter.Enqueue(dog)
		cat2 := newCat("Cat2")
		shelter.Enqueue(cat2)
		assert.Equal(t, "Dog1", shelter.DequeueDog().GetName())
		assert.Equal(t, "Cat1", shelter.DequeueAny().GetName())
		assert.Equal(t, "Cat2", shelter.DequeueAny().GetName())
	})

	t.Run("DequeueCat returns nil given shelter with only dogs", func(t *testing.T) {
		shelter := newShelter()
		dog := newDog("Dog1")
		shelter.Enqueue(dog)
		assert.Nil(t, shelter.DequeueCat())
	})

	t.Run("DequeueCat returns first cat given shelter with multiple animals", func(t *testing.T) {
		shelter := newShelter()
		dog1 := newDog("Dog1")
		shelter.Enqueue(dog1)
		cat := newCat("Cat1")
		shelter.Enqueue(cat)
		dog2 := newDog("Dog2")
		shelter.Enqueue(dog2)
		assert.Equal(t, "Cat1", shelter.DequeueCat().GetName())
		assert.Equal(t, "Dog1", shelter.DequeueAny().GetName())
		assert.Equal(t, "Dog2", shelter.DequeueAny().GetName())
	})
}
