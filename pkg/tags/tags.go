package tags

import (
	"reflect"

	"github.com/go-auxiliaries/tagmap"
	"github.com/go-auxiliaries/tagmap/pkg/registry"
)

type TagMap[V any] struct {
	values   []V
	registry *registry.TagRegistry
	zero     V
}

func New[V any](r *registry.TagRegistry) *TagMap[V] {
	return &TagMap[V]{
		values:   make([]V, r.GetLen()),
		registry: r,
		zero:     *new(V),
	}
}

func (m *TagMap[V]) IsTagName(name tagmap.TagName) bool {
	return m.registry.GetTag(name) != tagmap.UnknownTag
}

func (m *TagMap[V]) TagByName(name tagmap.TagName) tagmap.Tag {
	return m.registry.GetTag(name)
}

// GetByName gets tag value by tag name
// !! It will fail if tag is unknown !!
// Make sure you validated tag name via IsTagName
func (m *TagMap[V]) GetByName(name tagmap.TagName) V {
	return m.GetByTag(m.TagByName(name))
}

func (m *TagMap[V]) GetByTag(tag tagmap.Tag) V {
	return m.values[tag]
}

// SetByName sets tag value by tag name
// !! It will fail if tag is unknown !!
// Make sure you validated tag name via IsTagName
func (m *TagMap[V]) SetByName(name tagmap.TagName, val V) {
	m.SetByTag(m.TagByName(name), val)
}

func (m *TagMap[V]) SetByTag(tag tagmap.Tag, val V) {
	m.values[tag] = val
}

// GetByNameOrSet sets tag value by tag name
// !! It will fail if tag is unknown !!
// Make sure you validated tag name via IsTagName
func (m *TagMap[V]) GetByNameOrSet(name tagmap.TagName, val V) (V, bool) {
	return m.GetByTagOrSet(m.TagByName(name), val)
}

func (m *TagMap[V]) GetByTagOrSet(tag tagmap.Tag, val V) (V, bool) {
	e := m.values[tag]
	if reflect.ValueOf(val).IsZero() {
		m.values[tag] = val
		return val, false
	}
	return e, true
}

// GetByNameAndDelete sets tag value by tag name
// !! It will fail if tag is unknown !!
// Make sure you validated tag name via IsTagName
func (m *TagMap[V]) GetByNameAndDelete(name tagmap.TagName) V {
	return m.GetByTagAndDelete(m.TagByName(name))
}

func (m *TagMap[V]) GetByTagAndDelete(tag tagmap.Tag) V {
	out := m.values[tag]
	m.values[tag] = *new(V)
	return out
}

// DeleteByName sets tag value by tag name
// !! It will fail if tag is unknown !!
// Make sure you validated tag name via IsTagName
func (m *TagMap[V]) DeleteByName(name tagmap.TagName) {
	m.DeleteByTag(m.TagByName(name))
}

func (m *TagMap[V]) DeleteByTag(tag tagmap.Tag) {
	m.values[tag] = *new(V)
}

func (m *TagMap[V]) ValuesByTag() map[tagmap.Tag]V {
	out := make(map[tagmap.Tag]V, len(m.values))
	for tag, value := range m.values {
		out[tagmap.Tag(tag)] = value
	}
	return out
}

func (m *TagMap[V]) ValuesByName() map[tagmap.TagName]V {
	out := make(map[tagmap.TagName]V, len(m.values))
	for tag, value := range m.values {
		out[m.registry.GetName(tagmap.Tag(tag))] = value
	}
	return out
}

func (m *TagMap[V]) GetValuesByName(names ...tagmap.TagName) tagmap.List[V] {
	out := make(tagmap.List[V], len(names))
	for idx, name := range names {
		out[idx] = m.GetByName(name)
	}
	return out
}

func (m *TagMap[V]) GetValuesByTag(tags ...tagmap.Tag) tagmap.List[V] {
	out := make(tagmap.List[V], len(tags))
	for idx, tag := range tags {
		out[idx] = m.GetByTag(tag)
	}
	return out
}
