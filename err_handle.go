package main

import (
	"fmt"
	"errors"
)

func divide(a int, b int) (float64, error) {
	if b == 0 {
		return 0,errors.New("Cannot divide by 0")
	}
	return float64(a/b), nil
}


func main() {
	ans, err := divide(1,0)

	if err != nil {
		fmt.Println("ERROR: ",err)
		return
	} 

	fmt.Println("result is ",ans)

}