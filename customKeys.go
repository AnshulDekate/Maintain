package main

import (
	"fmt"
	"sort"
)

// key value store with indexes

// models
type Store struct {
	data map[CustomKey]int
}

type Row struct {
	key       string
	subKey    string
	timestamp int
	val       int
}

type CustomKey struct {
	key       string
	subKey    string
	timestamp int
}

// functions
func (s *Store) getKey(key string) []int {
	var ret []int
	for k, v := range s.data {
		if k.key == key {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *Store) getLastest(key string, subKey string) int {
	type Pair struct {
		timestamp int
		val       int
	}

	var a []Pair
	for k, v := range s.data {
		if k.key == key && k.subKey == subKey {
			a = append(a, Pair{k.timestamp, v})
		}
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].timestamp < a[j].timestamp
	})

	ret := -1
	if len(a) > 0 {
		ret = a[0].val
	}
	return ret
}

func (s *Store) get(key string, subKey string, timestamp int) int {
	val, ok := s.data[CustomKey{key, subKey, timestamp}]
	if ok {
		return val
	} else {
		return -1
	}
}

func (s *Store) deleteKey(key string) {
	var keys []CustomKey
	for k := range s.data {
		if k.key == key {
			keys = append(keys, k)
		}
	}
	for _, k := range keys {
		delete(s.data, k)
	}
}

func (s *Store) deleteSubKey(subKey string) {
	var keys []CustomKey
	for k := range s.data {
		if k.subKey == subKey {
			keys = append(keys, k)
		}
	}
	for _, k := range keys {
		delete(s.data, k)
	}
}

func main() {

	rows := []Row{
		{"a", "asub", 1, 3},
		{"b", "bsub", 1, 3},
		{"a", "asub1", 1, 5},
		{"a", "asub", 2, 5},
		{"a", "asub", 1, 3},
		{"c", "bsub", 1, 3},
	}

	store := Store{make(map[CustomKey]int, 0)}

	for _, row := range rows {
		store.data[CustomKey{row.key, row.subKey, row.timestamp}] = row.val
	}

	fmt.Println(store.data)

	fmt.Println(store.getKey("a"))
	fmt.Println(store.getLastest("a", "asub"))
	fmt.Println(store.get("a", "asub", 2))
	// store.deleteKey("a")
	store.deleteSubKey("bsub")

	fmt.Println(store.data)
}

