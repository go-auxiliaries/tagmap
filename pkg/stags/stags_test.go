package stags_test

import (
	"sync"
	"testing"

	"github.com/go-auxiliaries/tagmap"

	"github.com/go-auxiliaries/tagmap/pkg/registry"
	"github.com/go-auxiliaries/tagmap/pkg/stags"
	"github.com/stretchr/testify/assert"
)

var r = registry.New()

var tag1Name = tagmap.TagName("tag1")
var tag2Name = tagmap.TagName("tag2")
var tag3Name = tagmap.TagName("tag3")
var tag4Name = tagmap.TagName("tag4")

var tag1 = r.RegisterTag(tag1Name)
var tag2 = r.RegisterTag(tag2Name)
var tag3 = r.RegisterTag(tag3Name)

var tagMap = stags.New[string](r)

func Test(t *testing.T) {
	assert.True(t, tagMap.IsTagName(tag1Name))
	assert.True(t, tagMap.IsTagName(tag2Name))
	assert.True(t, tagMap.IsTagName(tag3Name))
	assert.False(t, tagMap.IsTagName(tag4Name))
	assert.Equal(t, tag1, tagMap.TagByName(tag1Name))
	assert.Equal(t, tag2, tagMap.TagByName(tag2Name))
	assert.Equal(t, tag3, tagMap.TagByName(tag3Name))
	assert.Equal(t, tagmap.UnknownTag, tagMap.TagByName(tag4Name))
	assert.Equal(t, "", tagMap.GetByTag(tag1))
	assert.Equal(t, "", tagMap.GetByTag(tag2))
	assert.Equal(t, "", tagMap.GetByTag(tag3))
	// Unknown tag will cause panic, you should make sure it is real tag via `IsTagName`
	//assert.Equal(t, "", tagMap.GetByTag(tag4))
	assert.Equal(t, "", tagMap.GetByName(tag1Name))
	assert.Equal(t, "", tagMap.GetByName(tag2Name))
	assert.Equal(t, "", tagMap.GetByName(tag3Name))
	// Unknown tag will cause panic, you should make sure it is real tag via `IsTagName`
	//assert.Equal(t, "", tagMap.GetByName("tag4"))

	tagMap.SetByTag(tag1, "SetByTag1")
	tagMap.SetByTag(tag2, "SetByTag2")
	tagMap.SetByTag(tag3, "SetByTag3")

	assert.Equal(t, "SetByTag1", tagMap.GetByTag(tag1))
	assert.Equal(t, "SetByTag2", tagMap.GetByTag(tag2))
	assert.Equal(t, "SetByTag3", tagMap.GetByTag(tag3))
	assert.Equal(t, "SetByTag1", tagMap.GetByName(tag1Name))
	assert.Equal(t, "SetByTag2", tagMap.GetByName(tag2Name))
	assert.Equal(t, "SetByTag3", tagMap.GetByName(tag3Name))

	tagMap.SetByName(tag1Name, "SetByTag1")
	tagMap.SetByName(tag2Name, "SetByTag2")
	tagMap.SetByName(tag3Name, "SetByTag3")

	assert.Equal(t, "SetByTag1", tagMap.GetByTag(tag1))
	assert.Equal(t, "SetByTag2", tagMap.GetByTag(tag2))
	assert.Equal(t, "SetByTag3", tagMap.GetByTag(tag3))
	assert.Equal(t, "SetByTag1", tagMap.GetByName(tag1Name))
	assert.Equal(t, "SetByTag2", tagMap.GetByName(tag2Name))
	assert.Equal(t, "SetByTag3", tagMap.GetByName(tag3Name))
}

func TestParallel(t *testing.T) {
	wg := sync.WaitGroup{}
	for n := 0; n < 10; n++ {
		wg.Add(1)
		go func() {
			for k := 0; k < 20; k++ {
				tagMap.SetByTag(tag1, "SetByTag1")
				tagMap.GetByName(tag2Name)

				tagMap.SetByTag(tag2, "SetByTag2")
				tagMap.GetByName(tag3Name)

				tagMap.SetByTag(tag3, "SetByTag3")
				tagMap.GetByName(tag1Name)

				tagMap.SetByName(tag1Name, "SetByTag1")
				tagMap.GetByTag(tag2)

				tagMap.SetByName(tag2Name, "SetByTag2")
				tagMap.GetByTag(tag3)

				tagMap.SetByName(tag3Name, "SetByTag3")
				tagMap.GetByTag(tag1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
