package main

import "fmt"

func main()  {
	fmt.Println("Sum of 6 and 4 is: ", sum(6, 4))
}

func sum(num1, num2 int32) int32 {
	return num1 + num2
}