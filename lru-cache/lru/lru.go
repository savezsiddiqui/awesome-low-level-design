package lru

type Cache interface {
	Get(int) int
	Put(int, int)
	Delete(int)
}

type Node struct {
	key  int
	val  int
	next *Node
	prev *Node
}

type LRUCache struct {
	mp       map[int]*Node
	head     *Node
	tail     *Node
	capacity int
}

func (c *LRUCache) Get(key int) int {
	if node, found := c.mp[key]; found {
		val := node.val
		c.removeNode(node)
		n := &Node{
			key: key,
			val: val,
		}
		c.insertAtBeginning(n)
		c.mp[key] = n
		return val
	}

	return -1
}

func (c *LRUCache) Put(key int, val int) {
	if len(c.mp) >= c.capacity {
		node := c.tail.prev
		c.removeLastNode()
		delete(c.mp, node.key)
	}

	if node, found := c.mp[key]; found {
		c.removeNode(node)
	}

	n := &Node{
		key: key,
		val: val,
	}

	c.insertAtBeginning(n)
	c.mp[key] = n
}

func (c *LRUCache) Delete(key int) {
	if node, found := c.mp[key]; found {
		c.removeNode(node)
		delete(c.mp, key)
	}
}

func (c *LRUCache) insertAtBeginning(n *Node) {
	tmp := c.head.next
	c.head.next = n
	n.next = tmp
	n.prev = c.head
	tmp.prev = n
}

func (c *LRUCache) removeNode(n *Node) {
	prev := n.prev
	next := n.next
	prev.next = n.next
	next.prev = prev
}

func (c *LRUCache) removeLastNode() {
	c.removeNode(c.tail.prev)
}

func NewLruCache(c int) Cache {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return &LRUCache{
		mp:       make(map[int]*Node),
		head:     head,
		tail:     tail,
		capacity: c,
	}
}
