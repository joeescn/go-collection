package collection

type Map[Key comparable, Value any] interface {
	// Set the value for a key.
	Set(Key, Value)
	// Get the value for a key.
	Get(Key) (Value, bool)
	// Has check if a key exists.
	Has(Key) bool
	// Delete a key.
	Delete(Key)
	// Clear all keys.
	Clear()
	// Size return the number of keys.
	Size() int
	// Clone returns a clone of the map.
	Clone() Map[Key, Value]
	// Values returns a slice of the values.
	Values() []Value
	// Keys returns a slice of the keys.
	Keys() []Key
	// Range iterates over the map.
	Range(func(Key, Value) bool)
}

// NewMap returns a new empty map.
// 	map is thread unsafe
func NewMap[Key comparable, Value any]() Map[Key, Value] {
	bm := newBaseMap[Key, Value]()
	return &bm
}

// NewSyncMap returns a new empty sync map.
// 	map is thread safe
func NewSyncMap[Key comparable, Value any]() Map[Key, Value] {
	bm := newSyncMap[Key, Value]()
	return &bm
}
