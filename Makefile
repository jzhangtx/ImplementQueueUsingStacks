ImplementQueueUsingStacks: ImplementQueueUsingStacks.o
	g++ -g -o ImplementQueueUsingStacks.exe ImplementQueueUsingStacks.o -pthread    

ImplementQueueUsingStacks.o: ImplementQueueUsingStacks/ImplementQueueUsingStacks.cpp
	g++ -g  -c -pthread ImplementQueueUsingStacks/ImplementQueueUsingStacks.cpp
