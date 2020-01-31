package queue

type CircularQueue struct {
	queue  []int
	q_head int
	q_tail int
	q_cap  int
	q_len  int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) CircularQueue {
	Queue := CircularQueue{
		queue:  make([]int, 0, k),
		q_head: 0,
		q_tail: 0,
		q_cap:  k,
		q_len:  0,
	}
	for i:=0; i < k; i++{
		Queue.queue = append(Queue.queue, -1)
	}

	return Queue
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *CircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
/** length of queue is 0 
*/
	if this.q_len != 0 {
		this.q_tail = (this.q_tail + 1) % this.q_cap
	}
	this.queue[this.q_tail] = value
	this.q_len++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *CircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.q_len--
    this.queue[this.q_head] = -1
/** length of queue is 1
*/
    if this.q_len != 0 {
	   this.q_head = (this.q_head + 1) % this.q_cap
    }
	return true
}

/** Get the front item from the queue. */
func (this *CircularQueue) Front() int {
	front := this.queue[this.q_head]
	return front
}

/** Get the last item from the queue. */
func (this *CircularQueue) Rear() int {
	rear := this.queue[this.q_tail]
	return rear
}

/** Checks whether the circular queue is empty or not. */
func (this *CircularQueue) IsEmpty() bool {
	if this.q_len != 0 {
		return false
	}
	return true
}

/** Checks whether the circular queue is full or not. */
func (this *CircularQueue) IsFull() bool {
	if this.q_len != this.q_cap {
		return false
	}
	return true
}

/**
 * Your CircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

/*func main() {
	obj := Constructor(5)
	param_1 := obj.EnQueue(1)
	param_2 := obj.DeQueue()
	param_3 := obj.Front()
	param_4 := obj.Rear()
	param_5 := obj.IsEmpty()
	param_6 := obj.IsFull()
	print(param_1,param_2, param_3, param_4, param_5,param_6)
}*/
