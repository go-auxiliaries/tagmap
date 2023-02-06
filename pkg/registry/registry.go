package registry

import (
	"github.com/go-auxiliaries/tagmap"
)

type TagRegistry struct {
	tags    []tagmap.TagName
	backMap map[tagmap.TagName]int
}

func New() *TagRegistry {
	return &TagRegistry{
		tags:    make([]tagmap.TagName, 0),
		backMap: make(map[tagmap.TagName]int, 0),
	}
}

func (r *TagRegistry) RegisterTag(name tagmap.TagName) tagmap.Tag {
	_, ok := r.backMap[name]
	if ok {
		panic("tag with name " + name + " is already registered")
	}
	idx := len(r.tags)
	r.tags = append(r.tags, name)
	r.backMap[name] = idx
	return tagmap.Tag(idx)
}

func (r *TagRegistry) RegisterOrReuseTag(name tagmap.TagName) tagmap.Tag {
	idx, ok := r.backMap[name]
	if ok {
		return tagmap.Tag(idx)
	}
	idx = len(r.tags)
	r.tags = append(r.tags, name)
	r.backMap[name] = idx
	return tagmap.Tag(idx)
}

func (r *TagRegistry) GetName(tag tagmap.Tag) tagmap.TagName {
	return r.tags[tag]
}

func (r *TagRegistry) GetTag(name tagmap.TagName) tagmap.Tag {
	idx, ok := r.backMap[name]
	if ok {
		return tagmap.Tag(idx)
	}
	return tagmap.UnknownTag
}

func (r *TagRegistry) GetLen() int {
	return len(r.tags)
}
