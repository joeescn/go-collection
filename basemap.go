package collection

type baseMap[Key comparable, Value any] map[Key]Value

func newBaseMap[Key comparable, Value any]() baseMap[Key, Value] {
	return make(baseMap[Key, Value])
}

// Set the value for a key.
func (m *baseMap[Key, Value]) Set(key Key, value Value) {
	(*m)[key] = value
}

// Get the value for a key.
func (m *baseMap[Key, Value]) Get(key Key) (Value, bool) {
	value, found := (*m)[key]
	return value, found
}

// Has check if a key exists.
func (m *baseMap[Key, Value]) Has(key Key) bool {
	_, found := (*m)[key]
	return found
}

// Delete a key.
func (m *baseMap[Key, Value]) Delete(key Key) {
	delete(*m, key)
}

// Clear all keys.
func (m *baseMap[Key, Value]) Clear() {
	*m = newBaseMap[Key, Value]()
}

// Size return the number of keys.
func (m *baseMap[Key, Value]) Size() int {
	return len(*m)
}

// Values returns a slice of the values.
func (m *baseMap[Key, Value]) Values() []Value {
	var values []Value
	m.Range(func(key Key, value Value) bool {
		values = append(values, value)
		return false
	})
	return values
}

// Keys returns a slice of the keys.
func (m *baseMap[Key, Value]) Keys() []Key {
	var keys []Key
	m.Range(func(key Key, value Value) bool {
		keys = append(keys, key)
		return false
	})
	return keys
}

// Range iterates over the map.
func (m *baseMap[Key, Value]) Range(fn func(Key, Value) bool) {
	for key, value := range *m {
		if fn(key, value) {
			break
		}
	}
}

func (m *baseMap[Key, Value]) Clone() Map[Key, Value] {
	clone := newBaseMap[Key, Value]()
	m.Range(func(key Key, value Value) bool {
		clone.Set(key, value)
		return false
	})
	return &clone
}
