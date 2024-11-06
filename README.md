Implement Queue using Stacks


Implement a queue using one or more stacks.

The Queue class should support the following methods:

int size()
boolean isEmpty()
int front()
int back()
void push(int element)
void pop()
You can assume that you've access to a Stack class which provides the following methods:

int size()
boolean isEmpty()
int top()
void push(int element)
void pop()
Main method:
Queue created of size 10
Printed: queue.front() queue.back() queue.isEmpty() queue.size()
Pushed: 5
Printed: queue.front() queue.back() queue.isEmpty() queue.size()
Pushed: 6
Printed: queue.front() queue.back() queue.isEmpty() queue.size()
Pushed: 7
Printed: queue.front() queue.back() queue.isEmpty() queue.size()
Popped
Printed: queue.front() queue.back() queue.isEmpty() queue.size()
Popped
Printed: queue.front() queue.back() queue.isEmpty() queue.size()
Popped
Printed: queue.front() queue.back() queue.isEmpty() queue.size()
Expected Output
-1 -1 true 0
5 5 false 1
5 6 false 2
5 7 false 3
6 7 false 2
7 7 false 1
-1 -1 true 0