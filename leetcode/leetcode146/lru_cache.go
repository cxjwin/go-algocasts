package leetcode146

// https://leetcode.com/problems/lru-cache/

type node struct {
	key  int
	val  int
	prev *node
	next *node
}

// LRUCache - cache struct
type LRUCache struct {
	capacity int
	head     *node
	mMap     map[int]*node
}

// Constructor - create cache
func Constructor(capacity int) LRUCache {
	cache := LRUCache{capacity: capacity}
	cache.mMap = make(map[int]*node)
	cache.head = &node{key: -1, val: -1}

	cur := cache.head
	for i := 0; i < capacity-1; i++ {
		node := &node{key: -1, val: -1}
		cur.next = node
		node.prev = cur
		cur = node
	}
	cache.head.prev = cur
	cur.next = cache.head

	return cache
}

func (cache *LRUCache) move2Head(cur *node) {
	if cur == cache.head {
		cache.head = cache.head.prev
		return
	}

	cur.prev.next = cur.next
	cur.next.prev = cur.prev
	
	cur.next = cache.head.next
	cur.next.prev = cur
	cache.head.next = cur
	cur.prev = cache.head
}

// Get - get a value
func (cache *LRUCache) Get(key int) int {
	node, ok := cache.mMap[key]
	if !ok {
		return -1
	}
	cache.move2Head(node)
	return node.val
}

// Put - put a value
func (cache *LRUCache) Put(key int, value int) {
	node, ok := cache.mMap[key]
	if ok {
		node.val = value
		cache.move2Head(node)
	} else {
		oldKey := cache.head.key
		cache.head.key = key
		cache.head.val = value
		delete(cache.mMap, oldKey)
		cache.mMap[key] = cache.head
		cache.head = cache.head.prev
	}
}
