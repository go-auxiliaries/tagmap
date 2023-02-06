## Description ##
Non-locking, blazingly fast thread-safe map implementation, in some cases it is `113422%` faster than sync.Map.
It is targeting use case when you have limited number of keys.

## Implementation ##
All keys are indexed, while values are stored in the list according to the indexes.

## Tradeoff ##
The tradeoff is that memory is getting reserved for keys that are not occupied.
Therefore, if you are want to store structs, consider using pointer on structs. 

## How to use ##

1. Create tag registry, an instance where tags are registered: `var r = registry.New()`
2. As of know all tags should be registered before you instantiate any tagmap: `var tag1 = r.RegisterTag("tag1")`
3. Once you registered you can instantiate tagmap: `testMap := tags.New[string](r)`
4. Fastest way to access tags is `tagmap.tag` (int value): `testMap.SetByTag(tag1, "SetByTag1")`
5. Alternatively, you can access them by `tagmap.tagName` (string value): `testMap.SetByName("tag1", "SetByTag2")`

## Example ##

```go
package main

import (
	"fmt"
	"github.com/go-auxiliaries/tagmap"
	"github.com/go-auxiliaries/tagmap/pkg/registry"
	"github.com/go-auxiliaries/tagmap/pkg/tags"
)

var r = registry.New()

var tag1 = r.RegisterTag("tag1")
var tag2 = r.RegisterTag("tag2")
var tag3 = r.RegisterTag("tag3")

func main() {
	testMap := tags.New[string](r)
	testMap.SetByTag(tag1, "SetByTag1")
	testMap.SetByName("tag2", "SetByTag2")
	// map[someKey1:someVal1 someKey2:someVal2]
	fmt.Printf("%v\n", testMap.ValuesByName())
	testMap.DeleteByTag(tag1)
	// map[someKey2:someVal2]
	fmt.Printf("%v\n", testMap.ValuesByName())
	// someVal3, true
	val, ok := testMap.GetByNameOrSet("tag1", "someVal3")
	fmt.Printf("%v, %b\n", val, ok)
	// someVal3, false
	val, ok = testMap.GetByTagOrSet(tag2, "someVal4")
	fmt.Printf("%v, %b\n", val, ok)
}
```

## Benchmarks ##
These are the benchmarks against `sync.Map`

```shell
goos: linux
goarch: amd64
pkg: github.com/go-auxiliaries/tagmap/pkg/stags
cpu: 12th Gen Intel(R) Core(TM) i9-12900HK
Benchmark_TestSuite/Sync_Get_Existing_Parallel_100-20         	515149191	         1.943 ns/op
Benchmark_TestSuite/Stags_GetByName_Existing_Parallel_100-20  	1000000000	         1.210 ns/op
Benchmark_TestSuite/Stags_GetByTag_Existing_Parallel_100-20   	1000000000	         0.3865 ns/op
Benchmark_TestSuite/Sync_Set_Existing_Parallel_100-20         	 3612550	       455.1 ns/op
Benchmark_TestSuite/Stags_SetByName_Existing_Parallel_100-20  	148331192	         8.838 ns/op
Benchmark_TestSuite/Stags_SetByTag_Existing_Parallel_100-20   	164656738	         7.153 ns/op
Benchmark_TestSuite/Sync_Delete_Existing_Parallel_100-20      	576646665	         2.058 ns/op
Benchmark_TestSuite/Stags_DeleteByName_Existing_Parallel_100-20         	291168806	         4.107 ns/op
Benchmark_TestSuite/Stags_DeleteByTag_Existing_Parallel_100-20          	376814229	         3.159 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_Existing_Parallel_100-20          	557630275	         2.289 ns/op
Benchmark_TestSuite/Stags_GetByNameAndDelete_Existing_Parallel_100-20   	180761204	         6.686 ns/op
Benchmark_TestSuite/Stags_GetByTagAndDelete_Existing_Parallel_100-20    	166247649	         7.046 ns/op
Benchmark_TestSuite/Sync_GetOrSet_Existing_Parallel_100-20              	84145364	        13.05 ns/op
Benchmark_TestSuite/Stags_GetByNameOrSet_Existing_Parallel_100-20       	123586713	         9.640 ns/op
Benchmark_TestSuite/Stags_GetByTagOrSet_Existing_Parallel_100-20        	132756217	         8.712 ns/op
Benchmark_TestSuite/Sync_Mixed_Existing_Parallel_100-20                 	137217655	         9.042 ns/op
Benchmark_TestSuite/Stags_MixedByName_Existing_Parallel_100-20          	177474196	         7.310 ns/op
Benchmark_TestSuite/Stags_MixedByTag_Existing_Parallel_100-20           	200294023	         5.983 ns/op
Benchmark_TestSuite/Sync_Get_Existing_Parallel_10000-20                 	323890902	         3.863 ns/op
Benchmark_TestSuite/Stags_GetByName_Existing_Parallel_10000-20          	336633704	         2.971 ns/op
Benchmark_TestSuite/Stags_GetByTag_Existing_Parallel_10000-20           	1000000000	         0.3968 ns/op
Benchmark_TestSuite/Sync_Set_Existing_Parallel_10000-20                 	 2795330	       376.5 ns/op
Benchmark_TestSuite/Stags_SetByName_Existing_Parallel_10000-20          	121793810	        10.46 ns/op
Benchmark_TestSuite/Stags_SetByTag_Existing_Parallel_10000-20           	166172490	         7.460 ns/op
Benchmark_TestSuite/Sync_Delete_Existing_Parallel_10000-20              	299370060	         4.023 ns/op
Benchmark_TestSuite/Stags_DeleteByName_Existing_Parallel_10000-20       	365563194	         3.324 ns/op
Benchmark_TestSuite/Stags_DeleteByTag_Existing_Parallel_10000-20        	1000000000	         0.8529 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_Existing_Parallel_10000-20        	297906320	         3.694 ns/op
Benchmark_TestSuite/Stags_GetByNameAndDelete_Existing_Parallel_10000-20 	317368238	         3.647 ns/op
Benchmark_TestSuite/Stags_GetByTagAndDelete_Existing_Parallel_10000-20  	1000000000	         1.064 ns/op
Benchmark_TestSuite/Sync_GetOrSet_Existing_Parallel_10000-20            	57248026	        21.47 ns/op
Benchmark_TestSuite/Stags_GetByNameOrSet_Existing_Parallel_10000-20     	116730278	         9.920 ns/op
Benchmark_TestSuite/Stags_GetByTagOrSet_Existing_Parallel_10000-20      	152563394	         7.274 ns/op
Benchmark_TestSuite/Sync_Mixed_Existing_Parallel_10000-20               	85138987	        13.63 ns/op
Benchmark_TestSuite/Stags_MixedByName_Existing_Parallel_10000-20        	191351470	         6.305 ns/op
Benchmark_TestSuite/Stags_MixedByTag_Existing_Parallel_10000-20         	314240161	         3.521 ns/op
Benchmark_TestSuite/Sync_Get_Existing_Parallel_1000000-20               	110049453	         9.904 ns/op
Benchmark_TestSuite/Stags_GetByName_Existing_Parallel_1000000-20        	221542700	         4.835 ns/op
Benchmark_TestSuite/Stags_GetByTag_Existing_Parallel_1000000-20         	1000000000	         0.4105 ns/op
Benchmark_TestSuite/Sync_Set_Existing_Parallel_1000000-20               	 2271349	       501.8 ns/op
Benchmark_TestSuite/Stags_SetByName_Existing_Parallel_1000000-20        	99583053	        10.54 ns/op
Benchmark_TestSuite/Stags_SetByTag_Existing_Parallel_1000000-20         	276534032	         4.619 ns/op
Benchmark_TestSuite/Sync_Delete_Existing_Parallel_1000000-20            	73604054	        15.21 ns/op
Benchmark_TestSuite/Stags_DeleteByName_Existing_Parallel_1000000-20     	200250944	         6.011 ns/op
Benchmark_TestSuite/Stags_DeleteByTag_Existing_Parallel_1000000-20      	1000000000	         0.7163 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_Existing_Parallel_1000000-20      	82778274	        13.40 ns/op
Benchmark_TestSuite/Stags_GetByNameAndDelete_Existing_Parallel_1000000-20         	179542653	         6.817 ns/op
Benchmark_TestSuite/Stags_GetByTagAndDelete_Existing_Parallel_1000000-20          	1000000000	         0.7161 ns/op
Benchmark_TestSuite/Sync_GetOrSet_Existing_Parallel_1000000-20                    	34212993	        32.92 ns/op
Benchmark_TestSuite/Stags_GetByNameOrSet_Existing_Parallel_1000000-20             	100000000	        10.06 ns/op
Benchmark_TestSuite/Stags_GetByTagOrSet_Existing_Parallel_1000000-20              	274029691	         4.617 ns/op
Benchmark_TestSuite/Sync_Mixed_Existing_Parallel_1000000-20                       	16559948	        61.90 ns/op
Benchmark_TestSuite/Stags_MixedByName_Existing_Parallel_1000000-20                	139525976	         8.961 ns/op
Benchmark_TestSuite/Stags_MixedByTag_Existing_Parallel_1000000-20                 	583090605	         2.375 ns/op
Benchmark_TestSuite/Sync_Get_Existing_Parallel_10000000-20                        	 3378445	       388.7 ns/op
Benchmark_TestSuite/Stags_GetByName_Existing_Parallel_10000000-20                 	230943303	         6.487 ns/op
Benchmark_TestSuite/Stags_GetByTag_Existing_Parallel_10000000-20                  	1000000000	         0.3424 ns/op
Benchmark_TestSuite/Sync_Set_Existing_Parallel_10000000-20                        	 2419239	       534.0 ns/op
Benchmark_TestSuite/Stags_SetByName_Existing_Parallel_10000000-20                 	100000000	        12.21 ns/op
Benchmark_TestSuite/Stags_SetByTag_Existing_Parallel_10000000-20                  	284541192	         4.341 ns/op
Benchmark_TestSuite/Sync_Delete_Existing_Parallel_10000000-20                     	 2621204	       525.3 ns/op
Benchmark_TestSuite/Stags_DeleteByName_Existing_Parallel_10000000-20              	150548632	         8.455 ns/op
Benchmark_TestSuite/Stags_DeleteByTag_Existing_Parallel_10000000-20               	1000000000	         0.6595 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_Existing_Parallel_10000000-20               	 2394014	       514.6 ns/op
Benchmark_TestSuite/Stags_GetByNameAndDelete_Existing_Parallel_10000000-20        	155029335	         7.914 ns/op
Benchmark_TestSuite/Stags_GetByTagAndDelete_Existing_Parallel_10000000-20         	1000000000	         0.7116 ns/op
Benchmark_TestSuite/Sync_GetOrSet_Existing_Parallel_10000000-20                   	 2862931	       602.4 ns/op
Benchmark_TestSuite/Stags_GetByNameOrSet_Existing_Parallel_10000000-20            	100000000	        11.98 ns/op
Benchmark_TestSuite/Stags_GetByTagOrSet_Existing_Parallel_10000000-20             	287029824	         4.654 ns/op
Benchmark_TestSuite/Sync_Mixed_Existing_Parallel_10000000-20                      	 2950747	       513.7 ns/op
Benchmark_TestSuite/Stags_MixedByName_Existing_Parallel_10000000-20               	121410548	        10.32 ns/op
Benchmark_TestSuite/Stags_MixedByTag_Existing_Parallel_10000000-20                	808862713	         2.090 ns/op
Benchmark_TestSuite/Sync_Get_NonExisting_Parallel_100-20                          	1000000000	         1.003 ns/op
Benchmark_TestSuite/Stags_GetByName_NonExisting_Parallel_100-20                   	995925337	         1.182 ns/op
Benchmark_TestSuite/Stags_GetByTag_NonExisting_Parallel_100-20                    	1000000000	         0.3737 ns/op
Benchmark_TestSuite/Sync_Set_NonExisting_Parallel_100-20                          	 2973510	       445.1 ns/op
Benchmark_TestSuite/Stags_SetByName_NonExisting_Parallel_100-20                   	192515263	         7.168 ns/op
Benchmark_TestSuite/Stags_SetByTag_NonExisting_Parallel_100-20                    	204248528	         6.182 ns/op
Benchmark_TestSuite/Sync_Delete_NonExisting_Parallel_100-20                       	955576756	         1.118 ns/op
Benchmark_TestSuite/Stags_DeleteByName_NonExisting_Parallel_100-20                	353896022	         3.481 ns/op
Benchmark_TestSuite/Stags_DeleteByTag_NonExisting_Parallel_100-20                 	387873159	         3.086 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_NonExisting_Parallel_100-20                 	965617872	         1.159 ns/op
Benchmark_TestSuite/Stags_GetByNameAndDelete_NonExisting_Parallel_100-20          	180745224	         6.668 ns/op
Benchmark_TestSuite/Stags_GetByTagAndDelete_NonExisting_Parallel_100-20           	181188679	         6.449 ns/op
Benchmark_TestSuite/Sync_GetOrSet_NonExisting_Parallel_100-20                     	94278711	        10.95 ns/op
Benchmark_TestSuite/Stags_GetByNameOrSet_NonExisting_Parallel_100-20              	141910474	         8.901 ns/op
Benchmark_TestSuite/Stags_GetByTagOrSet_NonExisting_Parallel_100-20               	154397931	         8.207 ns/op
Benchmark_TestSuite/Sync_Mixed_NonExisting_Parallel_100-20                        	166835875	         7.361 ns/op
Benchmark_TestSuite/Stags_MixedByName_NonExisting_Parallel_100-20                 	190970215	         6.831 ns/op
Benchmark_TestSuite/Stags_MixedByTag_NonExisting_Parallel_100-20                  	227518478	         5.976 ns/op
Benchmark_TestSuite/Sync_Get_NonExisting_Parallel_10000-20                        	988548008	         1.212 ns/op
Benchmark_TestSuite/Stags_GetByName_NonExisting_Parallel_10000-20                 	559923709	         2.039 ns/op
Benchmark_TestSuite/Stags_GetByTag_NonExisting_Parallel_10000-20                  	1000000000	         0.3781 ns/op
Benchmark_TestSuite/Sync_Set_NonExisting_Parallel_10000-20                        	 2378086	       451.7 ns/op
Benchmark_TestSuite/Stags_SetByName_NonExisting_Parallel_10000-20                 	179754109	         6.607 ns/op
Benchmark_TestSuite/Stags_SetByTag_NonExisting_Parallel_10000-20                  	275460876	         4.550 ns/op
Benchmark_TestSuite/Sync_Delete_NonExisting_Parallel_10000-20                     	1000000000	         1.221 ns/op
Benchmark_TestSuite/Stags_DeleteByName_NonExisting_Parallel_10000-20              	445604222	         2.587 ns/op
Benchmark_TestSuite/Stags_DeleteByTag_NonExisting_Parallel_10000-20               	1000000000	         0.8841 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_NonExisting_Parallel_10000-20               	973996096	         1.212 ns/op
Benchmark_TestSuite/Stags_GetByNameAndDelete_NonExisting_Parallel_10000-20        	348487789	         3.074 ns/op
Benchmark_TestSuite/Stags_GetByTagAndDelete_NonExisting_Parallel_10000-20         	1000000000	         1.248 ns/op
Benchmark_TestSuite/Sync_GetOrSet_NonExisting_Parallel_10000-20                   	57463047	        22.20 ns/op
Benchmark_TestSuite/Stags_GetByNameOrSet_NonExisting_Parallel_10000-20            	158191060	         7.280 ns/op
Benchmark_TestSuite/Stags_GetByTagOrSet_NonExisting_Parallel_10000-20             	244515508	         5.422 ns/op
Benchmark_TestSuite/Sync_Mixed_NonExisting_Parallel_10000-20                      	87337557	        15.00 ns/op
Benchmark_TestSuite/Stags_MixedByName_NonExisting_Parallel_10000-20               	278372388	         4.528 ns/op
Benchmark_TestSuite/Stags_MixedByTag_NonExisting_Parallel_10000-20                	575078730	         2.954 ns/op
Benchmark_TestSuite/Sync_Get_NonExisting_Parallel_1000000-20                      	1000000000	         1.303 ns/op
Benchmark_TestSuite/Stags_GetByName_NonExisting_Parallel_1000000-20               	275809702	         5.414 ns/op
Benchmark_TestSuite/Stags_GetByTag_NonExisting_Parallel_1000000-20                	1000000000	         0.3682 ns/op
Benchmark_TestSuite/Sync_Set_NonExisting_Parallel_1000000-20                      	 2546941	       653.3 ns/op
Benchmark_TestSuite/Stags_SetByName_NonExisting_Parallel_1000000-20               	100000000	        12.35 ns/op
Benchmark_TestSuite/Stags_SetByTag_NonExisting_Parallel_1000000-20                	291297054	         4.357 ns/op
Benchmark_TestSuite/Sync_Delete_NonExisting_Parallel_1000000-20                   	971296260	         1.241 ns/op
Benchmark_TestSuite/Stags_DeleteByName_NonExisting_Parallel_1000000-20            	203412058	         6.755 ns/op
Benchmark_TestSuite/Stags_DeleteByTag_NonExisting_Parallel_1000000-20             	1000000000	         0.7636 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_NonExisting_Parallel_1000000-20             	988498303	         1.150 ns/op
Benchmark_TestSuite/Stags_GetByNameAndDelete_NonExisting_Parallel_1000000-20      	172544164	         7.725 ns/op
Benchmark_TestSuite/Stags_GetByTagAndDelete_NonExisting_Parallel_1000000-20       	1000000000	         0.7711 ns/op
Benchmark_TestSuite/Sync_GetOrSet_NonExisting_Parallel_1000000-20                 	 4204712	       322.1 ns/op
Benchmark_TestSuite/Stags_GetByNameOrSet_NonExisting_Parallel_1000000-20          	100000000	        11.27 ns/op
Benchmark_TestSuite/Stags_GetByTagOrSet_NonExisting_Parallel_1000000-20           	273848727	         4.603 ns/op
Benchmark_TestSuite/Sync_Mixed_NonExisting_Parallel_1000000-20                    	 3996606	       448.9 ns/op
Benchmark_TestSuite/Stags_MixedByName_NonExisting_Parallel_1000000-20             	141074736	         8.758 ns/op
Benchmark_TestSuite/Stags_MixedByTag_NonExisting_Parallel_1000000-20              	762834824	         2.315 ns/op
Benchmark_TestSuite/Sync_Get_NonExisting_Parallel_10000000-20                     	1000000000	         1.089 ns/op
Benchmark_TestSuite/Stags_GetByName_NonExisting_Parallel_10000000-20              	250666585	         5.915 ns/op
Benchmark_TestSuite/Stags_GetByTag_NonExisting_Parallel_10000000-20               	1000000000	         0.3216 ns/op
Benchmark_TestSuite/Sync_Set_NonExisting_Parallel_10000000-20                     	 2077191	       601.1 ns/op
Benchmark_TestSuite/Stags_SetByName_NonExisting_Parallel_10000000-20              	100000000	        11.44 ns/op
Benchmark_TestSuite/Stags_SetByTag_NonExisting_Parallel_10000000-20               	321173930	         3.912 ns/op
Benchmark_TestSuite/Sync_Delete_NonExisting_Parallel_10000000-20                  	1000000000	         1.047 ns/op
Benchmark_TestSuite/Stags_DeleteByName_NonExisting_Parallel_10000000-20           	171868215	         8.917 ns/op
Benchmark_TestSuite/Stags_DeleteByTag_NonExisting_Parallel_10000000-20            	1000000000	         0.6819 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_NonExisting_Parallel_10000000-20            	1000000000	         1.011 ns/op
Benchmark_TestSuite/Stags_GetByNameAndDelete_NonExisting_Parallel_10000000-20     	154876414	         9.076 ns/op
Benchmark_TestSuite/Stags_GetByTagAndDelete_NonExisting_Parallel_10000000-20      	1000000000	         0.6872 ns/op
Benchmark_TestSuite/Sync_GetOrSet_NonExisting_Parallel_10000000-20                	 3775057	       366.3 ns/op
Benchmark_TestSuite/Stags_GetByNameOrSet_NonExisting_Parallel_10000000-20         	100000000	        11.25 ns/op
Benchmark_TestSuite/Stags_GetByTagOrSet_NonExisting_Parallel_10000000-20          	315483378	         3.993 ns/op
Benchmark_TestSuite/Sync_Mixed_NonExisting_Parallel_10000000-20                   	 3011869	       478.9 ns/op
Benchmark_TestSuite/Stags_MixedByName_NonExisting_Parallel_10000000-20            	100000000	        10.99 ns/op
Benchmark_TestSuite/Stags_MixedByTag_NonExisting_Parallel_10000000-20             	841262736	         2.175 ns/op
Benchmark_TestSuite/Sync_Get_Existing_Linear_100-20                               	64135998	        22.49 ns/op
Benchmark_TestSuite/Stags_GetByName_Existing_Linear_100-20                        	138167636	         8.610 ns/op
Benchmark_TestSuite/Stags_GetByTag_Existing_Linear_100-20                         	412108671	         3.185 ns/op
Benchmark_TestSuite/Sync_Set_Existing_Linear_100-20                               	12026851	       104.6 ns/op
Benchmark_TestSuite/Stags_SetByName_Existing_Linear_100-20                        	43753597	        30.86 ns/op
Benchmark_TestSuite/Stags_SetByTag_Existing_Linear_100-20                         	57533101	        23.95 ns/op
Benchmark_TestSuite/Sync_Delete_Existing_Linear_100-20                            	72704298	        16.14 ns/op
Benchmark_TestSuite/Stags_DeleteByName_Existing_Linear_100-20                     	100000000	        12.35 ns/op
Benchmark_TestSuite/Stags_DeleteByTag_Existing_Linear_100-20                      	150138570	         8.130 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_Existing_Linear_100-20                      	68751396	        15.95 ns/op
Benchmark_TestSuite/Stags_GetByNameAndDelete_Existing_Linear_100-20               	89137496	        13.24 ns/op
Benchmark_TestSuite/Stags_GetByTagAndDelete_Existing_Linear_100-20                	142515030	         8.823 ns/op
Benchmark_TestSuite/Sync_GetOrSet_Existing_Linear_100-20                          	15731851	        85.76 ns/op
Benchmark_TestSuite/Stags_GetByNameOrSet_Existing_Linear_100-20                   	37504826	        34.89 ns/op
Benchmark_TestSuite/Stags_GetByTagOrSet_Existing_Linear_100-20                    	34961576	        28.77 ns/op
Benchmark_TestSuite/Sync_Mixed_Existing_Linear_100-20                             	29658508	        43.78 ns/op
Benchmark_TestSuite/Stags_MixedByName_Existing_Linear_100-20                      	61113160	        21.49 ns/op
Benchmark_TestSuite/Stags_MixedByTag_Existing_Linear_100-20                       	85890772	        15.03 ns/op
Benchmark_TestSuite/Sync_Get_Existing_Linear_10000-20                             	32876323	        37.57 ns/op
Benchmark_TestSuite/Stags_GetByName_Existing_Linear_10000-20                      	79390954	        14.49 ns/op
Benchmark_TestSuite/Stags_GetByTag_Existing_Linear_10000-20                       	410397988	         2.946 ns/op
Benchmark_TestSuite/Sync_Set_Existing_Linear_10000-20                             	 9842378	       115.6 ns/op
Benchmark_TestSuite/Stags_SetByName_Existing_Linear_10000-20                      	31134234	        40.28 ns/op
Benchmark_TestSuite/Stags_SetByTag_Existing_Linear_10000-20                       	49671159	        23.91 ns/op
Benchmark_TestSuite/Sync_Delete_Existing_Linear_10000-20                          	32735252	        38.66 ns/op
Benchmark_TestSuite/Stags_DeleteByName_Existing_Linear_10000-20                   	68336078	        18.10 ns/op
Benchmark_TestSuite/Stags_DeleteByTag_Existing_Linear_10000-20                    	148549945	         8.082 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_Existing_Linear_10000-20                    	31855213	        36.93 ns/op
Benchmark_TestSuite/Stags_GetByNameAndDelete_Existing_Linear_10000-20             	62047755	        18.50 ns/op
Benchmark_TestSuite/Stags_GetByTagAndDelete_Existing_Linear_10000-20              	122873876	        10.32 ns/op
Benchmark_TestSuite/Sync_GetOrSet_Existing_Linear_10000-20                        	15540217	        76.90 ns/op
Benchmark_TestSuite/Stags_GetByNameOrSet_Existing_Linear_10000-20                 	28233720	        42.06 ns/op
Benchmark_TestSuite/Stags_GetByTagOrSet_Existing_Linear_10000-20                  	45680296	        29.29 ns/op
Benchmark_TestSuite/Sync_Mixed_Existing_Linear_10000-20                           	20646872	        59.10 ns/op
Benchmark_TestSuite/Stags_MixedByName_Existing_Linear_10000-20                    	43911988	        28.39 ns/op
Benchmark_TestSuite/Stags_MixedByTag_Existing_Linear_10000-20                     	80575627	        14.35 ns/op
Benchmark_TestSuite/Sync_Get_Existing_Linear_1000000-20                           	12899958	        89.34 ns/op
Benchmark_TestSuite/Stags_GetByName_Existing_Linear_1000000-20                    	22023495	        52.68 ns/op
Benchmark_TestSuite/Stags_GetByTag_Existing_Linear_1000000-20                     	363218978	         3.112 ns/op
Benchmark_TestSuite/Sync_Set_Existing_Linear_1000000-20                           	 5820958	       220.5 ns/op
Benchmark_TestSuite/Stags_SetByName_Existing_Linear_1000000-20                    	 9396736	       134.0 ns/op
Benchmark_TestSuite/Stags_SetByTag_Existing_Linear_1000000-20                     	44078179	        23.93 ns/op
Benchmark_TestSuite/Sync_Delete_Existing_Linear_1000000-20                        	14949362	        93.34 ns/op
Benchmark_TestSuite/Stags_DeleteByName_Existing_Linear_1000000-20                 	20151462	        57.68 ns/op
Benchmark_TestSuite/Stags_DeleteByTag_Existing_Linear_1000000-20                  	146450450	         8.212 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_Existing_Linear_1000000-20                  	14337141	        76.33 ns/op
Benchmark_TestSuite/Stags_GetByNameAndDelete_Existing_Linear_1000000-20           	19095189	        55.92 ns/op
Benchmark_TestSuite/Stags_GetByTagAndDelete_Existing_Linear_1000000-20            	143765811	         8.334 ns/op
Benchmark_TestSuite/Sync_GetOrSet_Existing_Linear_1000000-20                      	 6629923	       166.9 ns/op
Benchmark_TestSuite/Stags_GetByNameOrSet_Existing_Linear_1000000-20               	 8736514	       141.2 ns/op
Benchmark_TestSuite/Stags_GetByTagOrSet_Existing_Linear_1000000-20                	44569380	        29.33 ns/op
Benchmark_TestSuite/Sync_Mixed_Existing_Linear_1000000-20                         	10018512	       115.7 ns/op
Benchmark_TestSuite/Stags_MixedByName_Existing_Linear_1000000-20                  	14373344	        97.75 ns/op
Benchmark_TestSuite/Stags_MixedByTag_Existing_Linear_1000000-20                   	80074353	        16.03 ns/op
Benchmark_TestSuite/Sync_Get_Existing_Linear_10000000-20                          	 7703142	       174.1 ns/op
Benchmark_TestSuite/Stags_GetByName_Existing_Linear_10000000-20                   	22040661	        71.62 ns/op
Benchmark_TestSuite/Stags_GetByTag_Existing_Linear_10000000-20                    	375924350	         3.187 ns/op
Benchmark_TestSuite/Sync_Set_Existing_Linear_10000000-20                          	 5182489	       216.9 ns/op
Benchmark_TestSuite/Stags_SetByName_Existing_Linear_10000000-20                   	10122578	       159.3 ns/op
Benchmark_TestSuite/Stags_SetByTag_Existing_Linear_10000000-20                    	56861983	        23.70 ns/op
Benchmark_TestSuite/Sync_Delete_Existing_Linear_10000000-20                       	 5888670	       193.9 ns/op
Benchmark_TestSuite/Stags_DeleteByName_Existing_Linear_10000000-20                	20703673	        81.60 ns/op
Benchmark_TestSuite/Stags_DeleteByTag_Existing_Linear_10000000-20                 	148995268	         8.061 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_Existing_Linear_10000000-20                 	 6451372	       162.6 ns/op
Benchmark_TestSuite/Stags_GetByNameAndDelete_Existing_Linear_10000000-20          	16990797	        96.97 ns/op
Benchmark_TestSuite/Stags_GetByTagAndDelete_Existing_Linear_10000000-20           	140878298	         9.532 ns/op
Benchmark_TestSuite/Sync_GetOrSet_Existing_Linear_10000000-20                     	 5392754	       227.9 ns/op
Benchmark_TestSuite/Stags_GetByNameOrSet_Existing_Linear_10000000-20              	 9002740	       163.9 ns/op
Benchmark_TestSuite/Stags_GetByTagOrSet_Existing_Linear_10000000-20               	47349920	        29.14 ns/op
Benchmark_TestSuite/Sync_Mixed_Existing_Linear_10000000-20                        	 5921178	       207.8 ns/op
Benchmark_TestSuite/Stags_MixedByName_Existing_Linear_10000000-20                 	13039176	       109.6 ns/op
Benchmark_TestSuite/Stags_MixedByTag_Existing_Linear_10000000-20                  	79292043	        14.82 ns/op
Benchmark_TestSuite/Sync_Get_NonExisting_Linear_100-20                            	132590212	         8.766 ns/op
Benchmark_TestSuite/Stags_GetByName_NonExisting_Linear_100-20                     	138643650	         8.417 ns/op
Benchmark_TestSuite/Stags_GetByTag_NonExisting_Linear_100-20                      	412713123	         2.908 ns/op
Benchmark_TestSuite/Sync_Set_NonExisting_Linear_100-20                            	10556001	       104.2 ns/op
Benchmark_TestSuite/Stags_SetByName_NonExisting_Linear_100-20                     	44909979	        32.86 ns/op
Benchmark_TestSuite/Stags_SetByTag_NonExisting_Linear_100-20                      	51275067	        23.09 ns/op
Benchmark_TestSuite/Sync_Delete_NonExisting_Linear_100-20                         	141884536	         8.565 ns/op
Benchmark_TestSuite/Stags_DeleteByName_NonExisting_Linear_100-20                  	103735947	        12.37 ns/op
Benchmark_TestSuite/Stags_DeleteByTag_NonExisting_Linear_100-20                   	146405194	         7.977 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_NonExisting_Linear_100-20                   	138289502	         8.865 ns/op
Benchmark_TestSuite/Stags_GetByNameAndDelete_NonExisting_Linear_100-20            	91323559	        13.02 ns/op
Benchmark_TestSuite/Stags_GetByTagAndDelete_NonExisting_Linear_100-20             	145475472	         8.308 ns/op
Benchmark_TestSuite/Sync_GetOrSet_NonExisting_Linear_100-20                       	16034251	        68.62 ns/op
Benchmark_TestSuite/Stags_GetByNameOrSet_NonExisting_Linear_100-20                	36379666	        34.78 ns/op
Benchmark_TestSuite/Stags_GetByTagOrSet_NonExisting_Linear_100-20                 	44250291	        28.83 ns/op
Benchmark_TestSuite/Sync_Mixed_NonExisting_Linear_100-20                          	31156238	        39.44 ns/op
Benchmark_TestSuite/Stags_MixedByName_NonExisting_Linear_100-20                   	61584320	        21.17 ns/op
Benchmark_TestSuite/Stags_MixedByTag_NonExisting_Linear_100-20                    	86706040	        14.95 ns/op
Benchmark_TestSuite/Sync_Get_NonExisting_Linear_10000-20                          	136633240	         8.938 ns/op
Benchmark_TestSuite/Stags_GetByName_NonExisting_Linear_10000-20                   	85078352	        14.40 ns/op
Benchmark_TestSuite/Stags_GetByTag_NonExisting_Linear_10000-20                    	411826192	         2.908 ns/op
Benchmark_TestSuite/Sync_Set_NonExisting_Linear_10000-20                          	11281323	       110.5 ns/op
Benchmark_TestSuite/Stags_SetByName_NonExisting_Linear_10000-20                   	31088023	        40.52 ns/op
Benchmark_TestSuite/Stags_SetByTag_NonExisting_Linear_10000-20                    	53308947	        24.14 ns/op
Benchmark_TestSuite/Sync_Delete_NonExisting_Linear_10000-20                       	136189429	         9.643 ns/op
Benchmark_TestSuite/Stags_DeleteByName_NonExisting_Linear_10000-20                	69146665	        16.94 ns/op
Benchmark_TestSuite/Stags_DeleteByTag_NonExisting_Linear_10000-20                 	147982254	         8.092 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_NonExisting_Linear_10000-20                 	138116193	         8.703 ns/op
Benchmark_TestSuite/Stags_GetByNameAndDelete_NonExisting_Linear_10000-20          	60107905	        18.53 ns/op
Benchmark_TestSuite/Stags_GetByTagAndDelete_NonExisting_Linear_10000-20           	142993947	         8.270 ns/op
Benchmark_TestSuite/Sync_GetOrSet_NonExisting_Linear_10000-20                     	16033028	        75.30 ns/op
Benchmark_TestSuite/Stags_GetByNameOrSet_NonExisting_Linear_10000-20              	28846483	        42.56 ns/op
Benchmark_TestSuite/Stags_GetByTagOrSet_NonExisting_Linear_10000-20               	44887286	        29.20 ns/op
Benchmark_TestSuite/Sync_Mixed_NonExisting_Linear_10000-20                        	23007126	        56.71 ns/op
Benchmark_TestSuite/Stags_MixedByName_NonExisting_Linear_10000-20                 	36245797	        28.87 ns/op
Benchmark_TestSuite/Stags_MixedByTag_NonExisting_Linear_10000-20                  	86776353	        14.93 ns/op
Benchmark_TestSuite/Sync_Get_NonExisting_Linear_1000000-20                        	137268302	         8.638 ns/op
Benchmark_TestSuite/Stags_GetByName_NonExisting_Linear_1000000-20                 	23275532	        52.03 ns/op
Benchmark_TestSuite/Stags_GetByTag_NonExisting_Linear_1000000-20                  	404432061	         3.304 ns/op
Benchmark_TestSuite/Sync_Set_NonExisting_Linear_1000000-20                        	 4072378	       245.6 ns/op
Benchmark_TestSuite/Stags_SetByName_NonExisting_Linear_1000000-20                 	 9089946	       139.3 ns/op
Benchmark_TestSuite/Stags_SetByTag_NonExisting_Linear_1000000-20                  	50094222	        24.18 ns/op
Benchmark_TestSuite/Sync_Delete_NonExisting_Linear_1000000-20                     	140567605	         8.491 ns/op
Benchmark_TestSuite/Stags_DeleteByName_NonExisting_Linear_1000000-20              	21012192	        58.13 ns/op
Benchmark_TestSuite/Stags_DeleteByTag_NonExisting_Linear_1000000-20               	149185338	         8.057 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_NonExisting_Linear_1000000-20               	137379998	         8.856 ns/op
Benchmark_TestSuite/Stags_GetByNameAndDelete_NonExisting_Linear_1000000-20        	19715952	        61.87 ns/op
Benchmark_TestSuite/Stags_GetByTagAndDelete_NonExisting_Linear_1000000-20         	143998003	         8.315 ns/op
Benchmark_TestSuite/Sync_GetOrSet_NonExisting_Linear_1000000-20                   	 4793258	       221.0 ns/op
Benchmark_TestSuite/Stags_GetByNameOrSet_NonExisting_Linear_1000000-20            	 8797029	       142.3 ns/op
Benchmark_TestSuite/Stags_GetByTagOrSet_NonExisting_Linear_1000000-20             	42555210	        31.61 ns/op
Benchmark_TestSuite/Sync_Mixed_NonExisting_Linear_1000000-20                      	10818464	       104.3 ns/op
Benchmark_TestSuite/Stags_MixedByName_NonExisting_Linear_1000000-20               	13636026	        86.74 ns/op
Benchmark_TestSuite/Stags_MixedByTag_NonExisting_Linear_1000000-20                	95300569	        14.98 ns/op
Benchmark_TestSuite/Sync_Get_NonExisting_Linear_10000000-20                       	128846654	         8.492 ns/op
Benchmark_TestSuite/Stags_GetByName_NonExisting_Linear_10000000-20                	22992602	        72.73 ns/op
Benchmark_TestSuite/Stags_GetByTag_NonExisting_Linear_10000000-20                 	376039119	         3.092 ns/op
Benchmark_TestSuite/Sync_Set_NonExisting_Linear_10000000-20                       	 3132130	       400.5 ns/op
Benchmark_TestSuite/Stags_SetByName_NonExisting_Linear_10000000-20                	 9282535	       166.9 ns/op
Benchmark_TestSuite/Stags_SetByTag_NonExisting_Linear_10000000-20                 	53623555	        23.73 ns/op
Benchmark_TestSuite/Sync_Delete_NonExisting_Linear_10000000-20                    	141092006	         8.457 ns/op
Benchmark_TestSuite/Stags_DeleteByName_NonExisting_Linear_10000000-20             	21747552	        81.83 ns/op
Benchmark_TestSuite/Stags_DeleteByTag_NonExisting_Linear_10000000-20              	137782957	         8.208 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_NonExisting_Linear_10000000-20              	139413709	         8.606 ns/op
Benchmark_TestSuite/Stags_GetByNameAndDelete_NonExisting_Linear_10000000-20       	13905829	        81.59 ns/op
Benchmark_TestSuite/Stags_GetByTagAndDelete_NonExisting_Linear_10000000-20        	125322384	         9.227 ns/op
Benchmark_TestSuite/Sync_GetOrSet_NonExisting_Linear_10000000-20                  	 3259032	       359.7 ns/op
Benchmark_TestSuite/Stags_GetByNameOrSet_NonExisting_Linear_10000000-20           	 8880156	       162.6 ns/op
Benchmark_TestSuite/Stags_GetByTagOrSet_NonExisting_Linear_10000000-20            	47481260	        30.99 ns/op
Benchmark_TestSuite/Sync_Mixed_NonExisting_Linear_10000000-20                     	 4353163	       457.6 ns/op
Benchmark_TestSuite/Stags_MixedByName_NonExisting_Linear_10000000-20              	12880959	       111.6 ns/op
Benchmark_TestSuite/Stags_MixedByTag_NonExisting_Linear_10000000-20               	96291291	        18.37 ns/op
PASS
ok  	github.com/go-auxiliaries/tagmap/pkg/stags	1542.371s
```

### Benchmark highlights ###

#### Small Dataset (100 keys) ####
- On reading, it is `60%` faster than `sync.Map`. `1.210` against `1.943`
- On reading using tags, it is `466%` faster than `sync.Map`. `0.3430` against `1.943`
- On writing, it is `5049%` faster than `sync.Map`. `8.838` against `455.1`
- On writing using tags, it is `6262%` faster than `sync.Map`. `7.153` against `455.1`
- On deleting, it is `99%` slower than `sync.Map`. `4.107` against `2.058`
- On deleting using tags, it is `53%` slower than `sync.Map`. `3.159` against `2.058`
- On mixed workload, it is `23%` faster than `sync.Map`. `7.310` against `9.042`
- On mixed workload using tags, it is `51%` faster than `sync.Map`. `5.983` against `9.042`

```shell
Benchmark_TestSuite/Sync_Get_Existing_Parallel_100-20         	515149191	         1.943 ns/op
Benchmark_TestSuite/Stags_GetByName_Existing_Parallel_100-20  	1000000000	         1.210 ns/op
Benchmark_TestSuite/Stags_GetByTag_Existing_Parallel_100-20   	1000000000	         0.3865 ns/op
Benchmark_TestSuite/Sync_Set_Existing_Parallel_100-20         	 3612550	       455.1 ns/op
Benchmark_TestSuite/Stags_SetByName_Existing_Parallel_100-20  	148331192	         8.838 ns/op
Benchmark_TestSuite/Stags_SetByTag_Existing_Parallel_100-20   	164656738	         7.153 ns/op
Benchmark_TestSuite/Sync_Delete_Existing_Parallel_100-20      	576646665	         2.058 ns/op
Benchmark_TestSuite/Stags_DeleteByName_Existing_Parallel_100-20         	291168806	         4.107 ns/op
Benchmark_TestSuite/Stags_DeleteByTag_Existing_Parallel_100-20          	376814229	         3.159 ns/op
Benchmark_TestSuite/Sync_Mixed_Existing_Parallel_100-20                 	137217655	         9.042 ns/op
Benchmark_TestSuite/Stags_MixedByName_Existing_Parallel_100-20          	177474196	         7.310 ns/op
Benchmark_TestSuite/Stags_MixedByTag_Existing_Parallel_100-20           	200294023	         5.983 ns/op
```

#### Big Dataset (10000 keys) ####
- On reading, it is `30%` faster than `sync.Map`. `2.971` against `3.863`
- On reading using tags, it is `9635%` faster than `sync.Map`. `0.3968` against `3.863`
- On writing, it is `3499%` faster than `sync.Map`. `10.46` against `376.5`
- On writing using tags, it is `5046%` faster than `sync.Map`. `7.460` against `376.5`
- On deleting, it is `21%` faster than `sync.Map`. `3.324` against `4.023`
- On deleting using tags, it is `371%` faster than `sync.Map`. `0.8529` against `4.023`
- On mixed workload, it is `116%` faster than `sync.Map`. `6.305` against `13.63`
- On mixed workload using tags, it is `287%` faster than `sync.Map`. `3.521` against `13.63`

```shell
Benchmark_TestSuite/Sync_Get_Existing_Parallel_10000-20                 	323890902	         3.863 ns/op
Benchmark_TestSuite/Stags_GetByName_Existing_Parallel_10000-20          	336633704	         2.971 ns/op
Benchmark_TestSuite/Stags_GetByTag_Existing_Parallel_10000-20           	1000000000	         0.3968 ns/op
Benchmark_TestSuite/Sync_Set_Existing_Parallel_10000-20                 	 2795330	       376.5 ns/op
Benchmark_TestSuite/Stags_SetByName_Existing_Parallel_10000-20          	121793810	        10.46 ns/op
Benchmark_TestSuite/Stags_SetByTag_Existing_Parallel_10000-20           	166172490	         7.460 ns/op
Benchmark_TestSuite/Sync_Delete_Existing_Parallel_10000-20              	299370060	         4.023 ns/op
Benchmark_TestSuite/Stags_DeleteByName_Existing_Parallel_10000-20       	365563194	         3.324 ns/op
Benchmark_TestSuite/Stags_DeleteByTag_Existing_Parallel_10000-20        	1000000000	         0.8529 ns/op
Benchmark_TestSuite/Sync_Mixed_Existing_Parallel_10000-20               	85138987	        13.63 ns/op
Benchmark_TestSuite/Stags_MixedByName_Existing_Parallel_10000-20        	191351470	         6.305 ns/op
Benchmark_TestSuite/Stags_MixedByTag_Existing_Parallel_10000-20         	314240161	         3.521 ns/op
```

#### Huge Dataset (10000000 keys) ####
- On reading, it is `5891%` faster than `sync.Map`. `6.487` against `388.7`
- On reading using tags, it is `113422%` faster than `sync.Map`. `0.3781` against `388.7`
- On writing, it is `4273%` faster than `sync.Map`. `12.21` against `534.0`
- On writing using tags, it is `12201%` faster than `sync.Map`. `4.341` against `534.0`
- On deleting, it is `6112%` faster than `sync.Map`. `8.455` against `525.3`
- On deleting using tags, it is `79551%` faster than `sync.Map`. `0.6595` against `525.3`
- On mixed workload, it is `4877%` faster than `sync.Map`. `10.32` against `513.7`
- On mixed workload using tags, it is `24478%` faster than `sync.Map`. `2.090` against `513.7`

```shell
Benchmark_TestSuite/Sync_Get_Existing_Parallel_10000000-20                        	 3378445	       388.7 ns/op
Benchmark_TestSuite/Stags_GetByName_Existing_Parallel_10000000-20                 	230943303	         6.487 ns/op
Benchmark_TestSuite/Stags_GetByTag_Existing_Parallel_10000000-20                  	1000000000	         0.3424 ns/op
Benchmark_TestSuite/Sync_Set_Existing_Parallel_10000000-20                        	 2419239	       534.0 ns/op
Benchmark_TestSuite/Stags_SetByName_Existing_Parallel_10000000-20                 	100000000	        12.21 ns/op
Benchmark_TestSuite/Stags_SetByTag_Existing_Parallel_10000000-20                  	284541192	         4.341 ns/op
Benchmark_TestSuite/Sync_Delete_Existing_Parallel_10000000-20                     	 2621204	       525.3 ns/op
Benchmark_TestSuite/Stags_DeleteByName_Existing_Parallel_10000000-20              	150548632	         8.455 ns/op
Benchmark_TestSuite/Stags_DeleteByTag_Existing_Parallel_10000000-20               	1000000000	         0.6595 ns/op
Benchmark_TestSuite/Sync_Mixed_Existing_Parallel_10000000-20                      	 2950747	       513.7 ns/op
Benchmark_TestSuite/Stags_MixedByName_Existing_Parallel_10000000-20               	121410548	        10.32 ns/op
Benchmark_TestSuite/Stags_MixedByTag_Existing_Parallel_10000000-20                	808862713	         2.090 ns/op
```

#### Benchmark insights ####
On small dataset (100 keys), you can get - `9635%` speedup in reading and `5046%` speedup on writing.
While mixed workload will give you only `51%` improvement due to the slower delete operations.

On huge dataset (10000000 keys), you can get - wapping `113422%` speedup in reading and crazy `12201%` speedup on writing.
And mixed workload will give you same level improvement of `24478%` over `sync.Map`.

In almost all cases it is superior to `sync.Map` and of course to map+sync.RWMutex
