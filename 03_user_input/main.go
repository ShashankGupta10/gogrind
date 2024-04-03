package main

import "fmt"

func main() {
	var num int
	num, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)

}
