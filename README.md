# maphelper

Package maphelper provides functions to sort maps.

[![Build Status](https://travis-ci.org/northbright/maphelper.svg?branch=master)](https://travis-ci.org/northbright/maphelper)
[![Go Report Card](https://goreportcard.com/badge/github.com/northbright/maphelper)](https://goreportcard.com/report/github.com/northbright/maphelper)
[![GoDoc](https://godoc.org/github.com/northbright/maphelper?status.svg)](https://godoc.org/github.com/northbright/maphelper)

#### Requirements
* Go1.8 and later

  It uses [sort.SliceStable](https://godoc.org/sort#SliceStable) which need Go1.8 and later.

#### Examples
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

        // Output:
        // b: 1
        // c: 2
        // a: 3

#### Documentation
* [API References](https://godoc.org/github.com/northbright/maphelper) 

#### License
* [MIT License](LICENSE)
