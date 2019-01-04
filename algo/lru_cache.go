package algo

// LRUCache LRU cache struct
type LRUCache struct {
	capacity int
	head     *Node
	cache    map[int]*Node
}

// Node Doubly linked node
type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

// LRUCacheConstructor LRUCache Constructor
func LRUCacheConstructor(capacity int) LRUCache {
	c := LRUCache{capacity: capacity, head: nil, cache: make(map[int]*Node)}

	// make head
	c.head = &Node{-1, -1, nil, nil}
	cur := c.head
	for i := 1; i < capacity; i++ {
		temp := &Node{-1, -1, nil, nil}
		cur.next = temp
		temp.prev = cur
		cur = temp
	}
	cur.next = c.head
	c.head.prev = cur

	return c
}

func (c *LRUCache) move2Front(v *Node) {
	if v == c.head {
		c.head = c.head.prev
		return
	}

	// detach
	v.prev.next = v.next
	v.next.prev = v.prev
	// attach
	v.next = c.head.next
	v.next.prev = v
	c.head.next = v
	v.prev = c.head
}

// Get a value
func (c *LRUCache) Get(key int) int {
	if c.cache == nil {
		return -1
	}

	v, ok := c.cache[key]
	if ok {
		c.move2Front(v)
		return v.value
	}

	return -1
}

// Put a value
func (c *LRUCache) Put(key int, value int) {
	// value will always be positive
	if value < 0 {
		return
	}

	v, ok := c.cache[key]
	if ok {
		v.value = value
		c.move2Front(v)
		return
	}

	if c.head.value != -1 {
		delete(c.cache, c.head.key)
	}

	c.head.key = key
	c.head.value = value
	c.cache[key] = c.head
	c.head = c.head.prev
}
