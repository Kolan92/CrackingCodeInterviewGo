package main

import (
	"constraints"
	"errors"
)

type Stack[T constraints.Ordered] struct {
	last *Node[T]
}

func newStack[T constraints.Ordered]() *Stack[T] {
	return &Stack[T]{}
}

func (stack *Stack[T]) Push(value T) {
	if stack.last == nil {
		stack.last = &Node[T]{
			value: value,
		}
	} else {
		stack.last = &Node[T]{
			value: value,
			next:  stack.last,
		}
	}
}

func (stack *Stack[T]) Pick() *T {
	if stack.last != nil {
		return &stack.last.value
	}
	return nil
}

func (stack *Stack[T]) Pop() *T {
	if stack.last != nil {
		lastValue := stack.last.value
		stack.last = stack.last.next
		return &lastValue
	}
	return nil
}

func (stack *Stack[T]) IsEmpty() bool {
	lastElement := stack.Pick()
	return lastElement == nil
}

type Queue[T constraints.Ordered] struct {
	first *Node[T]
}

func newQueue[T constraints.Ordered]() *Queue[T] {
	return &Queue[T]{}
}

func (queue *Queue[T]) Pick() *T {
	if queue.first != nil {
		return &queue.first.value
	}
	return nil
}

func (queue *Queue[T]) Enqueue(value T) {
	if queue.first == nil {
		queue.first = &Node[T]{
			value: value,
		}
	} else {
		current := queue.first
		for current.next != nil {
			current = current.next
		}
		current.next = &Node[T]{
			value: value,
		}
	}
}

func (queue *Queue[T]) Dequeue() *T {
	if queue.first != nil {
		firstValue := queue.first.value
		queue.first = queue.first.next
		return &firstValue
	}
	return nil
}

type MinStack[T constraints.Ordered] struct {
	valueStack   *Stack[T]
	minimumStack *Stack[T]
}

func newMinStack[T constraints.Ordered]() *MinStack[T] {
	return &MinStack[T]{
		valueStack:   newStack[T](),
		minimumStack: newStack[T](),
	}
}

func (minStack *MinStack[T]) Push(value T) {
	minStack.valueStack.Push(value)
	if currentMinimum := minStack.minimumStack.Pick(); currentMinimum == nil || (*currentMinimum) > value {
		minStack.minimumStack.Push(value)
	}
}

func (minStack *MinStack[T]) Pick() *T {
	return minStack.valueStack.Pick()
}

func (minStack *MinStack[T]) Pop() *T {
	lastValue := minStack.valueStack.Pop()

	if *lastValue == *minStack.minimumStack.Pick() {
		_ = minStack.minimumStack.Pop()
	}

	return lastValue
}

func (minStack *MinStack[T]) Min() *T {
	return minStack.minimumStack.Pick()
}

type HanoiGame struct {
	height         int
	sourceRod      *Stack[int]
	spareRod       *Stack[int]
	destinationRod *Stack[int]
}

func newHanoiGame(height int) (*HanoiGame, error) {
	if height < 1 {
		return nil, errors.New("Min Height for hanoi tower is 3")
	}
	game := &HanoiGame{
		height:         height,
		sourceRod:      newStack[int](),
		spareRod:       newStack[int](),
		destinationRod: newStack[int](),
	}

	for i := height; i >= 1; i-- {
		game.sourceRod.Push(i)
	}

	return game, nil
}

func (game *HanoiGame) Solve() {
	Move(game.height, game.sourceRod, game.destinationRod, game.spareRod)
}

func Move(numberOfDisks int, source, destination, spare *Stack[int]) {
	if numberOfDisks == 1 {
		Swap(source, destination)
	} else {
		Move(numberOfDisks-1, source, spare, destination)
		Move(1, source, destination, spare)
		Move(numberOfDisks-1, spare, destination, source)
	}
}

func Swap(source, destination *Stack[int]) {
	element := source.Pop()
	destination.Push(*element)
}

type QueueOnTwoStacks[T constraints.Ordered] struct {
	newestElements *Stack[T]
	oldestElements *Stack[T]
}

func newQueueOnTwoStacks[T constraints.Ordered]() *QueueOnTwoStacks[T] {
	return &QueueOnTwoStacks[T]{
		newestElements: newStack[T](),
		oldestElements: newStack[T](),
	}
}

func (queue *QueueOnTwoStacks[T]) Pick() *T {
	shiftStacks(queue.newestElements, queue.oldestElements)
	return queue.oldestElements.Pick()
}

func (queue *QueueOnTwoStacks[T]) Enqueue(value T) {
	queue.newestElements.Push(value)
}

func (queue *QueueOnTwoStacks[T]) Dequeue() *T {
	shiftStacks(queue.newestElements, queue.oldestElements)
	return queue.oldestElements.Pop()
}

func shiftStacks[T constraints.Ordered](from, to *Stack[T]) {
	if !to.IsEmpty() {
		return
	}

	for fromElement := from.Pop(); fromElement != nil; fromElement = from.Pop() {
		to.Push(*fromElement)
	}
}

type MaxStack[T constraints.Ordered] struct {
	top *Node[T]
}

func newMaxStack[T constraints.Ordered]() *MaxStack[T] {
	return &MaxStack[T]{}
}

func (stack *MaxStack[T]) Push(value T) {
	if stack.top == nil || value >= stack.top.value {
		stack.top = &Node[T]{
			value: value,
			next:  stack.top,
		}
	} else {
		previous := stack.top
		current := previous.next
		for current != nil && current.value > value {
			previous = current
			current = previous.next
		}

		newNode := &Node[T]{
			value: value,
			next:  current,
		}
		previous.next = newNode
	}
}

func (stack *MaxStack[T]) Pick() *T {
	if stack.top != nil {
		return &stack.top.value
	}
	return nil
}

func (stack *MaxStack[T]) Pop() *T {
	if stack.top != nil {
		lastValue := stack.top.value
		stack.top = stack.top.next
		return &lastValue
	}
	return nil
}

func (stack *MaxStack[T]) IsEmpty() bool {
	lastElement := stack.Pick()
	return lastElement == nil
}

type IAnimal interface {
	GetName() string
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

type Cat struct {
	Animal
}

type Dog struct {
	Animal
}

func newCat(name string) *Cat {
	return &Cat{
		Animal{
			Name: name,
		},
	}
}

func newDog(name string) *Dog {
	return &Dog{
		Animal{
			Name: name,
		},
	}
}

type Shelter struct {
	first *Node[IAnimal]
}

func newShelter() *Shelter {
	return &Shelter{}
}

func (shelter *Shelter) Pick() IAnimal {
	if shelter.first != nil {
		return shelter.first.value
	}
	return nil
}

func (shelter *Shelter) Enqueue(value IAnimal) {
	if shelter.first == nil {
		shelter.first = &Node[IAnimal]{
			value: value,
		}
	} else {
		current := shelter.first
		for current.next != nil {
			current = current.next
		}

		current.next = &Node[IAnimal]{
			value: value,
		}
	}
}

func (shelter *Shelter) DequeueAny() IAnimal {
	if shelter.first != nil {
		firstValue := shelter.first.value
		shelter.first = shelter.first.next
		return firstValue
	}
	return nil
}

func (shelter *Shelter) DequeueDog() *Dog {
	animal := DequeueSpecific[*Dog](shelter)

	dog, _ := animal.(*Dog)
	return dog
}

func (shelter *Shelter) DequeueCat() *Cat {
	animal := DequeueSpecific[*Cat](shelter)

	cat, _ :=  animal.(*Cat)
	return cat
}

func DequeueSpecific[T IAnimal](shelter *Shelter) IAnimal {
	if animal, isSpecificAnimal := shelter.first.value.(T); isSpecificAnimal {
		shelter.first = shelter.first.next
		return animal
	}

	previousAnimal := shelter.first
	currentAnimal := previousAnimal.next

	for currentAnimal != nil {

		if animal, isSpecificAnimal := currentAnimal.value.(T); isSpecificAnimal {
			previousAnimal.next = currentAnimal.next
			return animal
		}

		previousAnimal = currentAnimal
		currentAnimal = previousAnimal.next
	}

	return nil
}