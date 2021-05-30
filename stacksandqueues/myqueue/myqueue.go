package myqueue

import "github.com/eternity-wings/dev_talks/stacksandqueues"

type Queue struct {
	mainStack      stacksandqueues.Stack
	temporaryStack stacksandqueues.Stack
}

func (q *Queue) Enqueue(value int) {
	q.mainStack.Push(value)
}

func (q *Queue) Display() {
	q.mainStack.Display()
	q.temporaryStack.Display()
}

func (q *Queue) Dequeue() (value *int) {
	for value := q.mainStack.Pop(); value != nil; value = q.mainStack.Pop() {
		q.temporaryStack.Push(*value)
	}
	dequeuedValue := q.temporaryStack.Pop()
	for value := q.temporaryStack.Pop(); value != nil; value = q.temporaryStack.Pop() {
		q.mainStack.Push(*value)
	}
	return dequeuedValue
}

func (q *Queue) Dequeue1() (value *int) {
	return q.mainStack.Pop()
}

func (q *Queue) Enqueue1(queuedValue int) {
	for value := q.mainStack.Pop(); value != nil; value = q.mainStack.Pop() {
		q.temporaryStack.Push(*value)
	}
	q.temporaryStack.Push(queuedValue)
	for value := q.temporaryStack.Pop(); value != nil; value = q.temporaryStack.Pop() {
		q.mainStack.Push(*value)
	}
}
