package main

import "fmt"

type Input struct {
	Data [][]int `json:"data"`
}

type ZigZagIterator struct {
	data []IteratorI
}

type IteratorI interface {
	hasNext() bool
	next() int
}

type RangeIterator struct {
	start int
	end   int
	step  int
	curr  int
}

type ListIterator struct {
	data []int
	idx  int
}

func NewZigZagIterator() *ZigZagIterator {
	return &ZigZagIterator{make([]IteratorI, 0)}
}
func (it *ZigZagIterator) hasNext() bool {
	return len(it.data) > 0
}
func (it *ZigZagIterator) next() int {
	ret := -1
	for len(it.data) > 0 {
		first := it.data[0]
		it.data = it.data[1:]
		if first.hasNext() {
			ret = first.next()
			if first.hasNext() {
				it.data = append(it.data, first)
			}
			break
		}
	}
	return ret
}


func NewRangeIterator(start, end, step int) *RangeIterator {
	return &RangeIterator{start, end, step, start - step}
}
func (it *RangeIterator) hasNext() bool {
	if it.step > 0 {
		return it.curr+it.step <= it.end
	} else {
		return it.end <= it.curr+it.step
	}
}
func (it *RangeIterator) next() int {
	it.curr = it.curr + it.step
	return it.curr
}

func NewListIterator(a []int) *ListIterator {
	return &ListIterator{a, -1}
}
func (it *ListIterator) hasNext() bool {
	return it.idx+1 < len(it.data)
}
func (it *ListIterator) next() int {
	it.idx += 1
	return it.data[it.idx]
}

func main() {
	a := NewRangeIterator(1, 20, 2)
	b := NewListIterator([]int{11, 12})
	c := NewRangeIterator(100, 10, -20)
	z := NewZigZagIterator()
	z.data = append(z.data, a)
	z.data = append(z.data, b)
	z.data = append(z.data, c)

	for z.hasNext() {
		fmt.Print(z.next(), " ")
	}
}

