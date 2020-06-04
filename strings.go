package main

import (
		"fmt"
		"strings"
)

func main() {
	s := "sample string"

	// length
	fmt.Println("len ",len(s))

	// iterate over chars with ascii character
	for index,char := range s {
		fmt.Println("index is ",index, " of char ",char)
	}

	// strings are Immutable

	// multi line strings
	multi_line_s := `
	hello
	world
	`
	fmt.Println("multi line ", multi_line_s)

	// substring matching
	does_contain := strings.Contains("substring", "sub")
	fmt.Println("contians or not ", does_contain) // true
	// also checkout HasPrefix and HasSuffix
	
	// join array elements as string
	arr := []string{"a","b","c"}
	joined_str := strings.Join(arr, ",")
	fmt.Println("joined_str ",joined_str)

	// split
	back_arr := strings.Split(joined_str, ",")
	fmt.Println("arr is back ",back_arr)

	// to lower to Upper
	fmt.Println("to upper ", strings.ToUpper("dddd"))
	fmt.Println("to Lower ", strings.ToLower("DDDD"))

	// slices 
	sample_slice := make([]string,2)
	sample_slice[0] = "a"
	new_sample := append(sample_slice, "b")
	
	// issue-1
	// initially slice was "",""
	// on changing 1st element to "a" it has become "a",""
	// on appending be it has become "a","","b" NOT "a","b"

	fmt.Println("sample slice ",sample_slice)
	fmt.Println("new smaple ", new_sample)
	
	// append elements
	newest_sample := append(new_sample,"c","d", "z")
	fmt.Println("latest ",newest_sample)

	// slice
	sliced_sample := newest_sample[2:4]
	fmt.Println("sliced ",sliced_sample)

	// make a copy of slice
	copied := make([]string, 2)
	copy(copied, sliced_sample)
	fmt.Println("copied ",copied)

	// replace
	todo := "eat eat eat"
	part_replaced_todo := strings.Replace(todo,"eat","sleep",2)
	// strings.ReplaceAll only in go version above 1.12
	// total_replaced_todo := strings.ReplaceAll("oink oink oink", "oink", "moo") // moo moo moo
	fmt.Println("part_replaced_todo",part_replaced_todo)

	
}