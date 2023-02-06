package registry_test

import (
	"testing"

	"github.com/go-auxiliaries/tagmap"

	"github.com/go-auxiliaries/tagmap/pkg/registry"
	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	r := registry.New()
	var tag1 = r.RegisterTag("tag1")
	var tag2 = r.RegisterTag("tag2")
	var tag3 = r.RegisterTag("tag3")
	assert.Equal(t, tag1, r.GetTag("tag1"))
	assert.Equal(t, tag2, r.GetTag("tag2"))
	assert.Equal(t, tag3, r.GetTag("tag3"))
	assert.Equal(t, tagmap.TagName("tag1"), r.GetName(tag1))
	assert.Equal(t, tagmap.TagName("tag2"), r.GetName(tag2))
	assert.Equal(t, tagmap.TagName("tag3"), r.GetName(tag3))
}
