package main

import (
	"fmt"
	"sort"
)

func main() {
	// for dynamic arrays use slices instead
	var names []string

	names = append(names, "kau")
	names = append(names, "kau1")
	names = append(names, "kau2")

	// all names 
	fmt.Println(names)
	
	// only 1st two
	names = names[:2]
	fmt.Println(names)
	
	// itereate thru all elements
	for index,value := range names {
		fmt.Println(index, value)
	}
	

	// check if an element exists
	for _, item := range names {
		if item == "kau" {
			fmt.Println("there")
		}
	}

	// sort an array
	nums := []int{1,3,4,2}
	sort.Ints(nums)
	fmt.Println("sorted ",nums)

	// array of structs
	type Person struct {
		 name string
		 age int
	}
	p := Person{"k",12}
	p1 := Person{"k1", 14}
	p2 := Person{"k1", 9}
	p3 := Person{"k1", 5}
	var persons []Person
	persons = append(persons, p)
	persons = append(persons, p1)
	persons = append(persons, p2)
	persons = append(persons, p3)
	fmt.Println("persons ",persons)

	//sort the array of structs using `sort.SliceStable`
	sort.SliceStable(persons, func (i, j int) bool {
		return persons[i].age < persons[j].age
	 })
	fmt.Println("sorted persons ", persons)
}
 
