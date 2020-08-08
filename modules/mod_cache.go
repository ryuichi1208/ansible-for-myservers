package cache

import (
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"hash/fnv"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

type unexportedInterface interface {
	Set(string, interface{}, time.Duration)
	Add(string, interface{}, time.Duration) error
	Replace(string, interface{}, time.Duration) error
	Get(string) (interface{}, bool)
	Increment(string, int64) error
	IncrementInt(string, int) (int, error)
	IncrementInt8(string, int8) (int8, error)
	IncrementInt16(string, int16) (int16, error)
	IncrementInt32(string, int32) (int32, error)
	IncrementInt64(string, int64) (int64, error)
	IncrementUint(string, uint) (uint, error)
	IncrementUintptr(string, uintptr) (uintptr, error)
	IncrementUint8(string, uint8) (uint8, error)
	IncrementUint16(string, uint16) (uint16, error)
	IncrementUint32(string, uint32) (uint32, error)
	IncrementUint64(string, uint64) (uint64, error)
	IncrementFloat(string, float64) error
	IncrementFloat32(string, float32) (float32, error)
	IncrementFloat64(string, float64) (float64, error)
	Decrement(string, int64) error
	DecrementInt(string, int) (int, error)
	DecrementInt8(string, int8) (int8, error)
	DecrementInt16(string, int16) (int16, error)
	DecrementInt32(string, int32) (int32, error)
	DecrementInt64(string, int64) (int64, error)
	DecrementUint(string, uint) (uint, error)
	DecrementUintptr(string, uintptr) (uintptr, error)
	DecrementUint8(string, uint8) (uint8, error)
	DecrementUint16(string, uint16) (uint16, error)
	DecrementUint32(string, uint32) (uint32, error)
	DecrementUint64(string, uint64) (uint64, error)
	DecrementFloat(string, float64) error
	DecrementFloat32(string, float32) (float32, error)
	DecrementFloat64(string, float64) (float64, error)
	Delete(string)
	DeleteExpired()
	Items() map[string]*Item
	ItemCount() int
	Flush()
	Save(io.Writer) error
	SaveFile(string) error
	Load(io.Reader) error
	LoadFile(io.Reader) error
}

type Item struct {
	Object     interface{}
	Expiration *time.Time
}

// Returns true if the item has expired.
func (item *Item) Expired() bool {
	if item.Expiration == nil {
		return false
	}
	return item.Expiration.Before(time.Now())
}

type Cache struct {
	*cache
	// If this is confusing, see the comment at the bottom of New()
}

type cache struct {
	sync.RWMutex
	defaultExpiration time.Duration
	items             map[string]*Item
	janitor           *janitor
}
