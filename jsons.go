package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	
	/*
	data types supported are:
		bool for JSON booleans,
		float64 for JSON numbers,
		string for JSON strings, and
		nil for JSON null
	*/

	// making a json
	// 1.struct to JSON
	// to be encoded "The first char of the Key should be a captial one"
	type Person struct {
		Name string
		Age float64
		Is_married bool
	}

	p1 := Person{
		Name: "kau", 
		Age: 11, 
		Is_married: true,
	}

	// marshal it
	jsonData, err := json.Marshal(p1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonData))

	// Receiving and unmarshalling
	my_json := []byte(`
	{
		"name": "ka",
		"age": 12
	}
	
	`)
	var abt_me interface {}
	json.Unmarshal(my_json, &abt_me)
	my_data := abt_me.(map[string]interface {})

	for key,val := range my_data {
		fmt.Println(key, val)
	}
	
	
	// simplest way to handle with JSON
	var test = []byte(`{
		"name": "kau",
		"age" : 19
	}`)

	var x interface{}
	json.Unmarshal(test, &x)
	y, _ := x.(map[string]interface{})
	fmt.Println(y["age"])
}
