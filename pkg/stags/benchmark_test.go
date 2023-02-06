package stags_test

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"testing"

	"github.com/go-auxiliaries/tagmap"
	"github.com/go-auxiliaries/tagmap/pkg/registry"
	"github.com/go-auxiliaries/tagmap/pkg/stags"
)

const nRoutines = 20

type testFunc func(i int) (fillExisting func(), testBody func(tag int, name string))
type funcName string

const (
	syncGet          = funcName("Sync_Get")          // -> Load
	syncSet          = funcName("Sync_Set")          // -> Store
	syncDelete       = funcName("Sync_Delete")       // -> Delete
	syncGetAndDelete = funcName("Sync_GetAndDelete") // -> LoadAndDelete
	syncGetOrSet     = funcName("Sync_GetOrSet")     // -> LoadOrStore

	syncMixed = funcName("Sync_Mixed") // -> All together

	stagsGetByName          = funcName("Stags_GetByName")          // -> GetByName
	stagsSetByName          = funcName("Stags_SetByName")          // -> Set
	stagsDeleteByName       = funcName("Stags_DeleteByName")       // -> Delete
	stagsGetByNameAndDelete = funcName("Stags_GetByNameAndDelete") // -> GetByNameAndDelete
	stagsGetByNameOrSet     = funcName("Stags_GetByNameOrSet")     // -> GetByNameOrSet

	stagsGetByTag          = funcName("Stags_GetByTag")          // -> GetByTag
	stagsSetByTag          = funcName("Stags_SetByTag")          // -> Set
	stagsDeleteByTag       = funcName("Stags_DeleteByTag")       // -> Delete
	stagsGetByTagAndDelete = funcName("Stags_GetByTagAndDelete") // -> GetByTagAndDelete
	stagsGetByTagOrSet     = funcName("Stags_GetByTagOrSet")     // -> GetByTagOrSet

	stagsMixedByName = funcName("Stags_MixedByName") // -> All together
	stagsMixedByTag  = funcName("Stags_MixedByTag")  // -> All together
)

var funcNameList = []funcName{
	syncGet, stagsGetByName, stagsGetByTag,
	syncSet, stagsSetByName, stagsSetByTag,
	syncDelete, stagsDeleteByName, stagsDeleteByTag,
	syncGetAndDelete, stagsGetByNameAndDelete, stagsGetByTagAndDelete,
	syncGetOrSet, stagsGetByNameOrSet, stagsGetByTagOrSet,
	syncMixed, stagsMixedByName, stagsMixedByTag,
}

var funcNameToTestFunc = map[funcName]testFunc{
	syncGet: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		testMap := sync.Map{}
		return func() {
				fillSyncMap(i, &testMap)
			},
			func(tag int, name string) {
				testMap.Load(name)
			}
	},
	syncSet: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		testMap := sync.Map{}
		return func() {
				fillSyncMap(i, &testMap)
			},
			func(tag int, name string) {
				testMap.Store(name, name)
			}
	},
	syncDelete: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		testMap := sync.Map{}
		return func() {
				fillSyncMap(i, &testMap)
			},
			func(tag int, name string) {
				testMap.Delete(name)
			}
	},
	syncGetAndDelete: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		testMap := sync.Map{}
		return func() {
				fillSyncMap(i, &testMap)
			},
			func(tag int, name string) {
				testMap.LoadAndDelete(name)
			}
	},
	syncGetOrSet: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		testMap := sync.Map{}
		return func() {
				fillSyncMap(i, &testMap)
			},
			func(tag int, name string) {
				testMap.LoadOrStore(name, name)
			}
	},
	syncMixed: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		testMap := sync.Map{}
		return func() {
				fillSyncMap(i, &testMap)
			},
			func(tag int, name string) {
				switch tag % 5 {
				case 0:
					testMap.LoadOrStore(name, name)
				case 1:
					testMap.Load(name)
				case 2:
					testMap.Store(name, name)
				case 3:
					testMap.Delete(name)
				case 4:
					testMap.LoadAndDelete(name)
				}
			}
	},
	stagsGetByName: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		fillRegistryTags(i, r1)
		testMap := stags.New[string](r1)
		return func() {
				fillSTagsMap(i, testMap)
			},
			func(tag int, name string) {
				testMap.GetByName(tagmap.TagName(name))
			}
	},
	stagsSetByName: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		fillRegistryTags(i, r1)
		testMap := stags.New[string](r1)
		return func() {
				fillSTagsMap(i, testMap)
			},
			func(tag int, name string) {
				testMap.SetByName(tagmap.TagName(name), name)
			}
	},
	stagsDeleteByName: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		fillRegistryTags(i, r1)
		testMap := stags.New[string](r1)
		return func() {
				fillSTagsMap(i, testMap)
			},
			func(tag int, name string) {
				testMap.DeleteByName(tagmap.TagName(name))
			}
	},
	stagsGetByNameAndDelete: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		fillRegistryTags(i, r1)
		testMap := stags.New[string](r1)
		return func() {
				fillSTagsMap(i, testMap)
			},
			func(tag int, name string) {
				testMap.GetByNameAndDelete(tagmap.TagName(name))
			}
	},
	stagsGetByNameOrSet: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		fillRegistryTags(i, r1)
		testMap := stags.New[string](r1)
		return func() {
				fillSTagsMap(i, testMap)
			},
			func(tag int, name string) {
				testMap.GetByNameOrSet(tagmap.TagName(name), name)
			}
	},
	stagsGetByTag: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		fillRegistryTags(i, r1)
		testMap := stags.New[string](r1)
		return func() {
				fillSTagsMap(i, testMap)
			},
			func(tag int, name string) {
				testMap.GetByTag(tagmap.Tag(tag))
			}
	},
	stagsSetByTag: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		fillRegistryTags(i, r1)
		testMap := stags.New[string](r1)
		return func() {
				fillSTagsMap(i, testMap)
			},
			func(tag int, name string) {
				testMap.SetByTag(tagmap.Tag(tag), name)
			}
	},
	stagsDeleteByTag: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		fillRegistryTags(i, r1)
		testMap := stags.New[string](r1)
		return func() {
				fillSTagsMap(i, testMap)
			},
			func(tag int, name string) {
				testMap.DeleteByTag(tagmap.Tag(tag))
			}
	},
	stagsGetByTagAndDelete: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		fillRegistryTags(i, r1)
		testMap := stags.New[string](r1)
		return func() {
				fillSTagsMap(i, testMap)
			},
			func(tag int, name string) {
				testMap.GetByTagAndDelete(tagmap.Tag(tag))
			}
	},
	stagsGetByTagOrSet: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		fillRegistryTags(i, r1)
		testMap := stags.New[string](r1)
		return func() {
				fillSTagsMap(i, testMap)
			},
			func(tag int, name string) {
				testMap.GetByTagOrSet(tagmap.Tag(tag), name)
			}
	},
	stagsMixedByName: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		fillRegistryTags(i, r1)
		testMap := stags.New[string](r1)
		return func() {
				fillSTagsMap(i, testMap)
			},
			func(tag int, name string) {
				switch tag % 5 {
				case 0:
					testMap.GetByNameOrSet(tagmap.TagName(name), name)
				case 1:
					testMap.GetByName(tagmap.TagName(name))
				case 2:
					testMap.SetByName(tagmap.TagName(name), name)
				case 3:
					testMap.DeleteByName(tagmap.TagName(name))
				case 4:
					testMap.GetByNameAndDelete(tagmap.TagName(name))
				}
			}
	},
	stagsMixedByTag: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		fillRegistryTags(i, r1)
		testMap := stags.New[string](r1)
		return func() {
				fillSTagsMap(i, testMap)
			},
			func(tag int, name string) {
				switch tag % 5 {
				case 0:
					testMap.GetByTagOrSet(tagmap.Tag(tag), name)
				case 1:
					testMap.GetByTag(tagmap.Tag(tag))
				case 2:
					testMap.SetByTag(tagmap.Tag(tag), name)
				case 3:
					testMap.DeleteByTag(tagmap.Tag(tag))
				case 4:
					testMap.GetByTagAndDelete(tagmap.Tag(tag))
				}
			}
	},
}

var r1 = registry.New()

type testCase struct {
	name string
	body func(*testing.B)
}

func Benchmark_TestSuite(b *testing.B) {
	testCases := make([]testCase, 0)
	for _, parallel := range []bool{true, false} {
		for _, existing := range []bool{true, false} {
			for _, nUniqueKeys := range []int{100, 10000, 1000000, 10000000} {
				for _, fName := range funcNameList {
					parallelLocal := parallel
					existingLocal := existing
					nUniqueKeysLocal := nUniqueKeys
					fNameLocal := fName
					p := "Parallel"
					if parallelLocal == false {
						p = "Linear"
					}
					e := "Existing"
					if existingLocal == false {
						e = "NonExisting"
					}
					testCases = append(testCases, testCase{
						name: fmt.Sprintf("%s_%s_%s_%d", fNameLocal, e, p, nUniqueKeysLocal),
						body: func(b *testing.B) {
							i := b.N
							createExistingFn, testBody := funcNameToTestFunc[fNameLocal](nUniqueKeysLocal)
							if existingLocal {
								createExistingFn()
							}
							if parallelLocal {
								runParallelTest(b, i, nUniqueKeysLocal, testBody)
							} else {
								runSingleThreadTest(b, i, nUniqueKeysLocal, testBody)
							}
						},
					})
				}
			}
		}
	}
	for _, t := range testCases {
		b.Run(t.name, t.body)
	}
}

var cachedTagNames = make([]string, 0)

func getNamesList(nIterates int) []string {
	if len(cachedTagNames) < nIterates {
		for n := len(cachedTagNames); n <= nIterates; n++ {
			cachedTagNames = append(cachedTagNames, strconv.Itoa(n))
		}
	}
	return cachedTagNames
}

func runParallelTest(b *testing.B, nIterates, nUniqueKeys int, body func(tag int, name string)) {
	wg := sync.WaitGroup{}
	iteratesPerRoutine := nIterates / nRoutines
	tagNames := getNamesList(nUniqueKeys)
	runtime.GC()
	b.ResetTimer()
	for n := 0; n < nRoutines; n++ {
		wg.Add(1)
		go func() {
			for k := 0; k < iteratesPerRoutine; k++ {
				tag := k % nUniqueKeys
				body(tag, tagNames[tag])
			}
			wg.Done()
		}()
	}
	wg.Wait()
	b.StopTimer()
}

func runSingleThreadTest(b *testing.B, nIterates, nUniqueKeys int, body func(tag int, name string)) {
	tagNames := getNamesList(nUniqueKeys)
	runtime.GC()
	b.ResetTimer()
	for n := 0; n < nIterates; n++ {
		tag := n % nUniqueKeys
		body(tag, tagNames[tag])
	}
	b.StopTimer()
}

func fillSTagsMap(nIterates int, testMap *stags.SafeTagMap[string]) {
	for n := 0; n <= nIterates; n++ {
		strVal := strconv.FormatInt(int64(n), 10)
		testMap.SetByTag(tagmap.Tag(n), strVal)
	}
}

func fillSyncMap(nIterates int, testMap *sync.Map) {
	for n := 0; n <= nIterates; n++ {
		strVal := strconv.FormatInt(int64(n), 10)
		testMap.Store(strVal, strVal)
	}
}

func fillRegistryTags(nIterates int, r *registry.TagRegistry) {
	for n := 0; n <= nIterates+1; n++ {
		r.RegisterOrReuseTag(tagmap.TagName(strconv.Itoa(n)))
	}
}
