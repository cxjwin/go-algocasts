package algo

// Pair key/value pair
type Pair struct {
	Key int
	Val int
}

// MyHashMap my hash map struct
type MyHashMap struct {
	mask  int
	array []Pair
}

// MapSize map size
// const MapSize int = 16384 // 2^14
const MapSize int = 16 // 2^4

// ConstructorMyHashMap Initialize your data structure here.
func ConstructorMyHashMap() MyHashMap {
	m := MyHashMap{}
	m.mask = MapSize - 1
	m.array = make([]Pair, MapSize)
	return m
}

func hash(x int) int {
	x = ((x >> 16) ^ x) * 0x45d9f3b
	x = ((x >> 16) ^ x) * 0x45d9f3b
	x = (x >> 16) ^ x
	return x
}

// Put a key/value pair
/** value will always be non-negative. */
func (m *MyHashMap) Put(key int, value int) {
	idx := hash(key) & m.mask
	v := m.array[idx]
	for v.Key > 0 && v.Key != key {
		idx = (idx + 1) & m.mask
		v = m.array[idx]
	}

	m.array[idx] = Pair{Key: key, Val: value}
}

// Get a value
/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (m *MyHashMap) Get(key int) int {
	idx := hash(key) & m.mask
	v := m.array[idx]
	for v.Key > 0 && v.Key != key {
		idx = (idx + 1) & m.mask
		v = m.array[idx]
	}
	if v.Key == key {
		return v.Val
	}
	return -1
}

// Remove a key/value pair
/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (m *MyHashMap) Remove(key int) {
	idx := hash(key) & m.mask
	v := m.array[idx]
	for v.Key > 0 && v.Key != key {
		idx = (idx + 1) & m.mask
		v = m.array[idx]
	}
	if v.Key == key {
		m.array[idx].Val = -1
	}
}
