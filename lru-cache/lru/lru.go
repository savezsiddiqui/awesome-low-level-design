package lru

type ListNode struct {
	Key   int
	Value int
	Next  *ListNode
	Prev  *ListNode
}

type LRUCache struct {
	head *ListNode
	tail *ListNode
	size int
	mp   map[int]*ListNode
}

func NewLruCache(size int) *LRUCache {
	head := &ListNode{}
	tail := &ListNode{}

	head.Next = tail
	tail.Prev = head

	return &LRUCache{
		head: head,
		tail: tail,
		size: size,
		mp:   make(map[int]*ListNode),
	}
}

func (lru *LRUCache) insertNodeAtBeginning(ln *ListNode) {
	tmp := lru.head.Next
	lru.head.Next = ln
	ln.Next = tmp
	ln.Prev = lru.head
	tmp.Prev = ln
}

func (lru *LRUCache) removeNode(ln *ListNode) {
	next := ln.Next
	prev := ln.Prev

	prev.Next = next
	next.Prev = prev
}

func (lru *LRUCache) Put(key int, val int) {
	node, ok := lru.mp[key]

	if ok {
		lru.removeNode(node)
	}

	if len(lru.mp) >= lru.size {
		key := lru.tail.Prev.Key
		lru.removeNode(lru.tail.Prev)
		delete(lru.mp, key)
	}

	newNode := &ListNode{
		Value: val,
		Key:   key,
	}

	lru.insertNodeAtBeginning(newNode)
	lru.mp[key] = newNode
}

func (lru *LRUCache) Get(key int) int {
	node, ok := lru.mp[key]

	if ok {
		val := node.Value
		lru.removeNode(node)
		newNode := &ListNode{
			Key:   key,
			Value: val,
		}
		lru.insertNodeAtBeginning(newNode)
		lru.mp[key] = newNode
		return val
	} else {
		return -1
	}
}

func (lru *LRUCache) Delete(key int) {
	node, ok := lru.mp[key]

	if ok {
		lru.removeNode(node)
		delete(lru.mp, key)
	}
}
