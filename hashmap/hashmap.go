package hashmap

// HashMap Define hash map
// 定义哈希映射
type HashMap[K comparable, T any] struct {
	m map[K]T
}

// NewHashMap 使用New构造HashMap
// Using new to construct a hash map
func NewHashMap[K comparable, T any]() *HashMap[K, T] {
	return &HashMap[K, T]{
		m: make(map[K]T),
	}
}

// Put 添加数据，若Key存在将替换原来的数据
// Add data. If the key exists, the original data will be replaced
func (m *HashMap[K, T]) Put(key K, val T) {
	m.m[key] = val
}

// Remove 根据Key删除数据
// Delete data according to key
func (m *HashMap[K, T]) Remove(key K) {
	delete(m.m, key)
}

// ForEach 进行迭代循环遍历
// Perform iterative loop traversal
func (m *HashMap[K, T]) ForEach(fn func(k K, v T)) {
	for k, t := range m.m {
		fn(k, t)
	}
}
