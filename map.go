package main

import (
	"fmt"
)

func main() {
	// to add elements later u need to initalize the map using make 
	// if used map literal like ...
	// var sample_map[string]string
	// and if u try to assign the code will panic: "Assignment to entry in nil map"

	// Hence use make
	sample_map := make(map[string]string)
	sample_map["name"] = "kau"
	sample_map["age"] = "1"
	sample_map["phone"] = "10000"

	fmt.Println(sample_map)
	fmt.Println("len of map ",len(sample_map))

	// check is some key is there or not
	_,is_there := sample_map["name"]
	fmt.Println("is there ", is_there)

	// delete 
	delete(sample_map, "age")

	// iterate through all key value pairs
	for key,value := range sample_map {
		fmt.Println(key, value)
	}


}