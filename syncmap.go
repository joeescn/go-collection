package collection

import (
	"sync"
)

type syncMap[Key comparable, Value any] sync.Map

func newSyncMap[Key comparable, Value any]() syncMap[Key, Value] {
	return syncMap[Key, Value]{}
}

// Set the value for a key.
func (m *syncMap[Key, Value]) Set(key Key, value Value) {
	(*sync.Map)(m).Store(key, value)
}

// Get the value for a key.
func (m *syncMap[Key, Value]) Get(key Key) (Value, bool) {
	value, ok := (*sync.Map)(m).Load(key)
	if !ok {
		var empty Value
		return empty, ok
	}
	v, ok := value.(Value)
	return v, ok
}

// Has check if a key exists.
func (m *syncMap[Key, Value]) Has(key Key) bool {
	_, ok := (*sync.Map)(m).Load(key)
	return ok
}

// Delete a key.
func (m *syncMap[Key, Value]) Delete(key Key) {
	(*sync.Map)(m).Delete(key)
}

// Clear all keys.
func (m *syncMap[Key, Value]) Clear() {
	*m = newSyncMap[Key, Value]()
}

// Size return the number of keys.
func (m *syncMap[Key, Value]) Size() int {
	var size int
	(*sync.Map)(m).Range(func(key, value any) bool {
		size++
		return false
	})
	return size
}

// Clone returns a clone of the map.
func (m *syncMap[Key, Value]) Clone() Map[Key, Value] {
	var clone syncMap[Key, Value]
	m.Range(func(key Key, value Value) bool {
		clone.Set(key, value)
		return false
	})
	return &clone
}

// Values returns a slice of the values.
func (m *syncMap[Key, Value]) Values() []Value {
	var values []Value
	m.Range(func(key Key, value Value) bool {
		values = append(values, value)
		return false
	})
	return values
}

// Keys returns a slice of the keys.
func (m *syncMap[Key, Value]) Keys() []Key {
	var keys []Key
	m.Range(func(key Key, value Value) bool {
		keys = append(keys, key)
		return false
	})
	return keys
}

// Range iterates over the map.
func (m *syncMap[Key, Value]) Range(fn func(key Key, value Value) bool) {
	(*sync.Map)(m).Range(func(key, value any) bool {
		return fn(key.(Key), value.(Value))
	})
}
