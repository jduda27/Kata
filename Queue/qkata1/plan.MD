# Queue

A queue is first in first out. Information is processed from the front of the object and inserted at the end of the queue.

We are implementing a virtual queue for a new attraction that is expecting to have massive numbers of visitors the idea is we want
visitors to spend less time in line by having a virtual position in queue. When they reach the end of the virtual queue we notify 
them they can join the physical queue and another process will return the customers unique ticket id to be processed by the Queue
gateway scanning device.

Plan:

implement a queue. We don't know how big this queue will become but we know that a new visitor will be added to the end. and removed from the begining.

We can use a singularly linked list as we only need to know who is coming next.

We assume all ticket ID values are unique

To start we will create our queue (a list)

Our queue contains visitors (nodes)

We have two functions EnterQueue, getNextInQueue, getVisitorsAhead and ExitQueue

EnterQueue inserts a ticket ID string into the end of the queue.

ExitQueue removes a ticket ID from a queue with a provided ID.

GetNextInQueue will return the ID of the customer at the front of the queue.

RemoveNextInQueue will remove the next customer from the queue.

GetVisitorsAhead will cound the number of visitors ahead of the provided customer's id.

Exit queue can be used for both taking the first customer in line out of it and removing customer who decided to leave the virtual queue.


