// ImplementQueueUsingStacks.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>
class Stack
{
public:
	Stack(int capacity)
		: m_Capacity(capacity)
		, m_Size(0)
		, m_pArray(new int[capacity])
	{
	}

	~Stack()
	{
		delete[] m_pArray;
	}

	int size()
	{
		return m_Size;
	}

	bool isEmpty()
	{
		return m_Size == 0;
	}

	int top()
	{
		if (m_Size == 0)
			return -1;
		return m_pArray[m_Size - 1];
	}

	void push(int element)
	{
		if (m_Size == m_Capacity)
			return;
		m_pArray[m_Size++] = element;
	}

	void pop()
	{
		if (m_Size == 0)
			return;
		m_Size--;
	}

private:
	int m_Capacity;
	int* m_pArray;
	int m_Size;
};

// Implement the Queue class
class Queue {
public:
	Queue(int capacity)
		: m_Stack(capacity)
	{
	}

	bool isEmpty()
	{
		return m_Stack.isEmpty();
	}

	int size()
	{
		return m_Stack.size();
	}

	int front()
	{
		Stack tempStack(m_Stack.size());
		while (!m_Stack.isEmpty())
		{
			tempStack.push(m_Stack.top());
			m_Stack.pop();
		}
		int element = tempStack.top();
		while (!tempStack.isEmpty())
		{
			m_Stack.push(tempStack.top());
			tempStack.pop();
		}
		return element;
	}

	int back()
	{
		return m_Stack.top();
	}

	void push(int element)
	{
		m_Stack.push(element);
	}

	void pop()
	{
		Stack tempStack(m_Stack.size());
		while (!m_Stack.isEmpty())
		{
			tempStack.push(m_Stack.top());
			m_Stack.pop();
		}
		tempStack.pop();
		while (!tempStack.isEmpty())
		{
			m_Stack.push(tempStack.top());
			tempStack.pop();
		}
	}

private:
	Stack m_Stack;
};

int main()
{
	Queue q(10);
	int i = 0;
	bool finished = false;

	while (!finished)
	{
		std::cout << q.front() << " " << q.back() << " ";
		std::cout << (q.isEmpty() ? "true" : "false") << " " << q.size() << std::endl;
		switch (i++)
		{
		case 0:
			q.push(5);
			break;

		case 1:
			q.push(6);
			break;

		case 2:
			q.push(7);
			break;

		case 3:
		case 4:
		case 5:
			q.pop();
			break;

		default:
			finished = true;
		}
	}
}
