package maphelper

import (
	"fmt"
	"sort"
)

// strIntPair represents string->int pair.
type strIntPair struct {
	key   string
	value int
}

// strInt32Pair represents string->int32 pair.
type strInt32Pair struct {
	key   string
	value int32
}

// strInt64Pair represents string->int64 pair.
type strInt64Pair struct {
	key   string
	value int64
}

// strFloat32Pair represents string->float32 pair.
type strFloat32Pair struct {
	key   string
	value float32
}

// strFloat64Pair represents string->float64 pair.
type strFloat64Pair struct {
	key   string
	value float64
}

// SortMapByValues sorts the string -> number maps by values.
//
// Params:
//     m: supported map type:
//        map[string]int, map[string]int32, map[string]int64, map[string]float32, map[string]float64
//     descOder: true for DESC order, false for ASC order.
// Return:
//     Slice contains sorted keys by values.
//     It'll return 'invalid map type' for unsupported map types.
func SortMapByValues(m interface{}, descOrder bool) ([]string, error) {
	var sortedKeys []string

	switch m := m.(type) {
	case map[string]int:
		arr := []strIntPair{}
		for k, v := range m {
			arr = append(arr, strIntPair{k, v})
		}

		sort.SliceStable(arr, func(i, j int) bool {
			if !descOrder {
				return arr[i].value < arr[j].value
			}
			return arr[i].value > arr[j].value
		})

		for _, v := range arr {
			sortedKeys = append(sortedKeys, v.key)
		}
		return sortedKeys, nil

	case map[string]int32:
		arr := []strInt32Pair{}
		for k, v := range m {
			arr = append(arr, strInt32Pair{k, v})
		}

		sort.SliceStable(arr, func(i, j int) bool {
			if !descOrder {
				return arr[i].value < arr[j].value
			}
			return arr[i].value > arr[j].value
		})

		for _, v := range arr {
			sortedKeys = append(sortedKeys, v.key)
		}
		return sortedKeys, nil

	case map[string]int64:
		arr := []strInt64Pair{}
		for k, v := range m {
			arr = append(arr, strInt64Pair{k, v})
		}

		sort.SliceStable(arr, func(i, j int) bool {
			if !descOrder {
				return arr[i].value < arr[j].value
			}
			return arr[i].value > arr[j].value

		})

		for _, v := range arr {
			sortedKeys = append(sortedKeys, v.key)
		}
		return sortedKeys, nil

	case map[string]float32:
		arr := []strFloat32Pair{}
		for k, v := range m {
			arr = append(arr, strFloat32Pair{k, v})
		}

		sort.SliceStable(arr, func(i, j int) bool {
			if !descOrder {
				return arr[i].value < arr[j].value
			}
			return arr[i].value > arr[j].value

		})

		for _, v := range arr {
			sortedKeys = append(sortedKeys, v.key)
		}
		return sortedKeys, nil

	case map[string]float64:
		arr := []strFloat64Pair{}
		for k, v := range m {
			arr = append(arr, strFloat64Pair{k, v})
		}

		sort.SliceStable(arr, func(i, j int) bool {
			if !descOrder {
				return arr[i].value < arr[j].value
			}
			return arr[i].value > arr[j].value

		})

		for _, v := range arr {
			sortedKeys = append(sortedKeys, v.key)
		}
		return sortedKeys, nil

	default:
		return []string{}, fmt.Errorf("invalid map type")

	}
}
