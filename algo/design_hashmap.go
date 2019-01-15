package algo

// MyHashMap my hash map struct
type MyHashMap struct {
	array []int
}

// MapSize map size
const MapSize int = 1000000

// ConstructorMyHashMap Initialize your data structure here.
func ConstructorMyHashMap() MyHashMap {
	m := MyHashMap{}
	m.array = make([]int, MapSize)
	for i := 0; i < MapSize; i++ {
		m.array[i] = -1
	}
	return m
}

// Put a key/value pair
/** value will always be non-negative. */
func (m *MyHashMap) Put(key int, value int) {
	idx := key % MapSize
	m.array[idx] = value
}

// Get a value
/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (m *MyHashMap) Get(key int) int {
	idx := key % MapSize
	return m.array[idx]
}

// Remove a key/value pair
/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (m *MyHashMap) Remove(key int) {
	idx := key % MapSize
	m.array[idx] = -1
}
