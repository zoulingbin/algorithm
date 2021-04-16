package queue
/**
	head == tail 为空queue
 */
type CircularQueue struct {
	q 	[]interface{}
	head 	int
	tail 	int
	cap 	int
}

func NewCircularQueue(n int) *CircularQueue {
	if n == 0 {
		return nil
	}
	return &CircularQueue{q: make([]interface{},n),cap: n, head: 0, tail: 0}
}

func (c *CircularQueue)IsEmpty() bool {
	if c.head == c.tail {
		return true
	}
	return false
}

/**
 	head == (tail+1)%cap ,queue满
 */
func (c *CircularQueue)IsFull() bool {
	if c.head == (c.tail+1)%c.cap {
		return true
	}
	return false
}

func (c *CircularQueue)push(data interface{}) bool {
	if c.IsFull() {
		return false
	}
	c.q[c.tail] = data
	c.tail = (c.tail+1)%c.cap
	return true
}

func (c *CircularQueue)pop() interface{} {
	if c.IsEmpty() {
		return nil
	}
	v := c.q[c.head]
	c.head = (c.head + 1)%c.cap
	return v
}

func (c *CircularQueue) QueueLength() int {
	return (c.tail - c.head + c.cap)%c.cap
}