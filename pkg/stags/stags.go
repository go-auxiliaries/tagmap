package stags

import (
	"sync/atomic"
	"unsafe"

	"github.com/go-auxiliaries/tagmap"
	"github.com/go-auxiliaries/tagmap/pkg/registry"
)

type SafeTagMap[V any] struct {
	values   []*V
	registry *registry.TagRegistry
}

func New[V any](r *registry.TagRegistry) *SafeTagMap[V] {
	return &SafeTagMap[V]{
		registry: r,
		values:   make([]*V, r.GetLen()),
	}
}

func (m *SafeTagMap[V]) IsTagName(name tagmap.TagName) bool {
	return m.registry.GetTag(name) != tagmap.UnknownTag
}

func (m *SafeTagMap[V]) TagByName(name tagmap.TagName) tagmap.Tag {
	return m.registry.GetTag(name)
}

func (m *SafeTagMap[V]) getTag(name tagmap.TagName) tagmap.Tag {
	tag := m.TagByName(name)
	if tag == tagmap.UnknownTag {
		panic("there is no such tag with name " + name)
	}
	return tag
}

// GetByName gets tag value by tag name
// !! It will fail if tag is unknown !!
// Make sure you validated tag name via IsTagName
func (m *SafeTagMap[V]) GetByName(name tagmap.TagName) V {
	return m.GetByTag(m.getTag(name))
}

// SetByName sets tag value by tag name
// !! It will fail if tag is unknown !!
// Make sure you validated tag name via IsTagName
func (m *SafeTagMap[V]) SetByName(name tagmap.TagName, val V) {
	m.SetByTag(m.getTag(name), val)
}

func (m *SafeTagMap[V]) SetByName2(name tagmap.TagName, val *V) {
	m.SetByTag2(m.getTag(name), val)
}

func (m *SafeTagMap[V]) GetByTag(tag tagmap.Tag) V {
	val := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&m.values[tag])))
	if val == unsafe.Pointer(nil) {
		return *new(V)
	}
	return *(*V)(val)
}

func (m *SafeTagMap[V]) GetByNameOrSet(name tagmap.TagName, val V) (V, bool) {
	return m.GetByTagOrSet(m.getTag(name), val)
}

func (m *SafeTagMap[V]) GetByNameOrSet2(name tagmap.TagName, val *V) (*V, bool) {
	return m.GetByTagOrSet2(m.getTag(name), val)
}

func (m *SafeTagMap[V]) GetByTagOrSet(tag tagmap.Tag, val V) (V, bool) {
	ok := atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&m.values[tag])), unsafe.Pointer(nil), unsafe.Pointer(&val))
	if ok {
		return val, false
	}
	return *(*V)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&m.values[tag])))), true
}

func (m *SafeTagMap[V]) GetByTagOrSet2(tag tagmap.Tag, val *V) (*V, bool) {
	ok := atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&m.values[tag])), unsafe.Pointer(nil), unsafe.Pointer(val))
	if ok {
		return val, false
	}
	return (*V)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&m.values[tag])))), true
}

func (m *SafeTagMap[V]) GetByNameAndDelete(name tagmap.TagName) V {
	return m.GetByTagAndDelete(m.getTag(name))
}

func (m *SafeTagMap[V]) GetByTagAndDelete(tag tagmap.Tag) V {
	val := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&m.values[tag])))
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&m.values[tag])), unsafe.Pointer(nil))
	if val == unsafe.Pointer(nil) {
		return *new(V)
	}
	return *(*V)(val)
}

func (m *SafeTagMap[V]) SetByTag(tag tagmap.Tag, val V) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&m.values[tag])), unsafe.Pointer(&val))
}

func (m *SafeTagMap[V]) SetByTag2(tag tagmap.Tag, val *V) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&m.values[tag])), unsafe.Pointer(val))
}

func (m *SafeTagMap[V]) DeleteByName(name tagmap.TagName) {
	m.DeleteByTag(m.getTag(name))
}

func (m *SafeTagMap[V]) DeleteByTag(tag tagmap.Tag) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&m.values[tag])), unsafe.Pointer(nil))
}

func (m *SafeTagMap[V]) ValuesByTag() map[tagmap.Tag]V {
	out := make(map[tagmap.Tag]V, len(m.values))
	for tag := range m.values {
		value := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&m.values[tag])))
		if value != nil {
			out[tagmap.Tag(tag)] = *(*V)(value)
		}
	}
	return out
}

func (m *SafeTagMap[V]) ValuesByName() map[tagmap.TagName]V {
	out := make(map[tagmap.TagName]V, len(m.values))
	for tag := range m.values {
		value := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&m.values[tag])))
		if value != nil {
			out[m.registry.GetName(tagmap.Tag(tag))] = *(*V)(value)
		}
	}
	return out
}

func (m *SafeTagMap[V]) GetValuesByName(names ...tagmap.TagName) tagmap.List[V] {
	out := make(tagmap.List[V], len(names))
	for idx, name := range names {
		out[idx] = m.GetByName(name)
	}
	return out
}

func (m *SafeTagMap[V]) GetValuesByTag(tags ...tagmap.Tag) tagmap.List[V] {
	out := make(tagmap.List[V], len(tags))
	for idx, tag := range tags {
		out[idx] = m.GetByTag(tag)
	}
	return out
}
