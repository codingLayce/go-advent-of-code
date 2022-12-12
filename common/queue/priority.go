package queue

type queueElement[T any] struct {
	value    T
	priority int
}

type PriorityQueue[T any] struct {
	values []queueElement[T]
}

func NewPriorityQueue[T any]() PriorityQueue[T] {
	return PriorityQueue[T]{}
}

func (queue *PriorityQueue[T]) IsEmpty() bool {
	return len(queue.values) == 0
}

func (queue *PriorityQueue[T]) Push(value T, priority int) {
	queue.values = append(queue.values, queueElement[T]{value: value, priority: priority})
}

func (queue *PriorityQueue[T]) Pop() T {
	popIdx := 0
	popPriority := 0
	for idx, element := range queue.values {
		if element.priority < popPriority {
			popPriority = element.priority
			popIdx = idx
		}
	}

	element := queue.values[popIdx]

	queue.values[popIdx] = queue.values[len(queue.values)-1]
	queue.values = queue.values[:len(queue.values)-1]

	return element.value
}
