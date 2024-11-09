package main

import (
	"fmt"
	"sort"
)

type Person struct {
	first string
	age   int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.first, p.age)
}

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)
}

type ByName []Person

func (bn ByName) Len() int {
	return len(bn)
}

func (bn ByName) Swap(i, j int) {
	bn[i], bn[j] = bn[j], bn[i]
}

func (bn ByName) Less(i, j int) bool {
	return bn[i].first < bn[j].first
}
