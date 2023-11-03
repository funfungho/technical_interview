package main

import "fmt"

type BagI interface {
	Add(int)       // add an item
	IsEmpty() bool // is the bag empty?
	Size() int     // number of items in the bag
}

type IntBag map[int]bool

func (b *IntBag) Add(item int) {
	(*b)[item] = true
}

func (b *IntBag) IsEmpty() bool {
	return b.Size() == 0
}

func (b *IntBag) Size() int {
	return len(*b)
}

func Bag() BagI {
	b := IntBag(map[int]bool{})
	return &b
}

func main() {
	// create a bag
	bag := Bag()

	// test bag is empty
	if !bag.IsEmpty() {
		fmt.Printf("Bag should be empty\n")
	}

	// add items into the bag
	items := []int{1, 2, 3}
	for _, item := range items {
		bag.Add(item)
	}

	// test bag size
	if bag.Size() != len(items) {
		fmt.Printf("Bag size should be %d\n", len(items))
	}

	for _, item := range items {
		bagMap := *(bag.(*IntBag))
		if _, ok := bagMap[item]; !ok {
			fmt.Printf("Bag should contain %d\n", item)
		}
	}

	fmt.Println("Bag tests passed.")
}
