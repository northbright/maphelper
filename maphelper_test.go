package maphelper_test

import (
	"fmt"
	"log"

	"github.com/northbright/maphelper"
)

func ExampleSortMapByValues() {
	// string->int map:
	m1 := map[string]int{"a": 3, "c": 2, "b": 1}
	// Get sorted keys by value ASC order.
	sortedKeys, err := maphelper.SortMapByValues(m1, false)
	if err != nil {
		log.Printf("maphelper.SortMapByValues() error: %v", err)
	}
	for _, k := range sortedKeys {
		fmt.Printf("%v: %v\n", k, m1[k])
	}

	// string->int32 map:
	m2 := map[string]int32{"a": 300, "c": 200, "b": 100}
	// Get sorted keys by value DESC order.
	sortedKeys, err = maphelper.SortMapByValues(m2, true)
	if err != nil {
		log.Printf("maphelper.SortMapByValues() error: %v", err)
	}
	for _, k := range sortedKeys {
		fmt.Printf("%v: %v\n", k, m2[k])
	}

	// string->int64 map:
	m3 := map[string]int64{"a": 30000, "c": 20000, "b": 10000}
	// Get sorted keys by value ASC order.
	sortedKeys, err = maphelper.SortMapByValues(m3, false)
	if err != nil {
		log.Printf("maphelper.SortMapByValues() error: %v", err)
	}
	for _, k := range sortedKeys {
		fmt.Printf("%v: %v\n", k, m3[k])
	}

	// string->float32 map:
	m4 := map[string]float32{"a": 3.33, "c": 2.22, "b": 1.11}
	// Get sorted keys by value DESC order.
	sortedKeys, err = maphelper.SortMapByValues(m4, true)
	if err != nil {
		log.Printf("maphelper.SortMapByValues() error: %v", err)
	}
	for _, k := range sortedKeys {
		fmt.Printf("%v: %v\n", k, m4[k])
	}

	// string->float64 map:
	m5 := map[string]float64{"a": 33.33, "c": 22.22, "b": 11.11}
	// Get sorted keys by value DESC order.
	sortedKeys, err = maphelper.SortMapByValues(m5, false)
	if err != nil {
		log.Printf("maphelper.SortMapByValues() error: %v", err)
	}
	for _, k := range sortedKeys {
		fmt.Printf("%v: %v\n", k, m5[k])
	}

	// string->string map: unsupported.
	m6 := map[string]string{"a": "3", "c": "2", "b": "1"}
	// Get sorted keys by value DESC order.
	sortedKeys, err = maphelper.SortMapByValues(m6, false)
	if err != nil {
		fmt.Printf("maphelper.SortMapByValues() error: %v", err)
	}
	for _, k := range sortedKeys {
		fmt.Printf("%v: %v\n", k, m5[k])
	}

	// Output:
	// b: 1
	// c: 2
	// a: 3
	// a: 300
	// c: 200
	// b: 100
	// b: 10000
	// c: 20000
	// a: 30000
	// a: 3.33
	// c: 2.22
	// b: 1.11
	// b: 11.11
	// c: 22.22
	// a: 33.33
	// maphelper.SortMapByValues() error: invalid map type
}
