package main

import "fmt"

type DLink struct {
	m_Data int
	m_Next *DLink
	m_Prev *DLink
}

type Stack struct {
	m_Capacity int
	m_Size     int
	m_Stack    *DLink
}

func NewStack(capacity int) *Stack {
	s := &Stack{capacity, 0, nil}
	s.m_Stack = &DLink{-1, nil, nil}
	s.m_Stack.m_Next = s.m_Stack
	s.m_Stack.m_Prev = s.m_Stack

	return s
}

func (s *Stack) IsEmpty() bool {
	return s.m_Size == 0
}

func (s *Stack) Size() int {
	return s.m_Size
}

func (s *Stack) Top() int {
	return s.m_Stack.m_Next.m_Data
}

func (s *Stack) Push(element int) {
	if s.m_Size == s.m_Capacity {
		return
	}

	pE := &DLink{element, nil, nil}
	pE.m_Next = s.m_Stack.m_Next
	pE.m_Prev = s.m_Stack
	s.m_Stack.m_Next.m_Prev = pE
	s.m_Stack.m_Next = pE
	s.m_Size++
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}

	s.m_Stack.m_Next.m_Next.m_Prev = s.m_Stack
	s.m_Stack.m_Next = s.m_Stack.m_Next.m_Next
	s.m_Size--
}

type Queue struct {
	m_Stack *Stack
}

func NewQueue(capacity int) *Queue {
	return &Queue{NewStack(capacity)}
}

func (q *Queue) IsEmpty() bool {
	return q.m_Stack.IsEmpty()
}

func (q *Queue) Size() int {
	return q.m_Stack.Size()
}

func (q *Queue) Front() int {
	tempStack := NewStack(q.m_Stack.Size())
	for !q.m_Stack.IsEmpty() {
		tempStack.Push(q.m_Stack.Top())
		q.m_Stack.Pop()
	}

	result := tempStack.Top()

	for !tempStack.IsEmpty() {
		q.m_Stack.Push(tempStack.Top())
		tempStack.Pop()
	}

	return result
}

func (q *Queue) Back() int {
	return q.m_Stack.Top()
}

func (q *Queue) Push(element int) {
	q.m_Stack.Push(element)
}

func (q *Queue) Pop() {
	tempStack := NewStack(q.m_Stack.Size())
	for q.m_Stack.Size() > 1 {
		tempStack.Push(q.m_Stack.Top())
		q.m_Stack.Pop()
	}

	q.m_Stack.Pop()

	for !tempStack.IsEmpty() {
		q.m_Stack.Push(tempStack.Top())
		tempStack.Pop()
	}
}

func PrintTest(q *Queue) {
	// Commented is the normal implementation. The current code is as per the questions' requirement
	n := q.Front()
	fmt.Print(n, " ")
	n = q.Back()
	fmt.Print(n, " ")
	fmt.Println(q.IsEmpty(), q.Size())
}

func main() {
	q := NewQueue(10)
	PrintTest(q)

	q.Push(5)
	PrintTest(q)

	q.Push(6)
	PrintTest(q)

	q.Push(7)
	PrintTest(q)

	q.Pop()
	PrintTest(q)

	q.Pop()
	PrintTest(q)

	q.Pop()
	PrintTest(q)
}
