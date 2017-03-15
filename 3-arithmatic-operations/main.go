package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Demonstrating Arithmatic operations.")

	var num1 int = 6;
	var num2 int = 4;

	fmt.Printf("%d - %d = %d \n", num1, num2 ,num1 - num2)

	fmt.Printf("%d + %d = %d \n", num1, num2, num1 + num2)

	fmt.Printf("%d * %d = %d \n", num1, num2, num1*num2)

	fmt.Printf("%d / %d = %d \n", num1, num2, (num1/num2))

	fmt.Printf("%d modulous %d = %d \n", num1, num2, num1%num2)

	//Mathematical operations using math library.

	var num3 float64 = 100.00;
	fmt.Printf("Square root of %.4f is %.3f \n", num3, math.Sqrt(num3))

	fmt.Printf("Max value: %.3f \n", math.Max(13.34, num3));

	fmt.Printf("10 Power of 3 is: %f \n", math.Pow(10, 3));

	fmt.Printf("Absolute of 2.453: %f \n", math.Abs(-2.453));

	fmt.Printf("10 Mod of 3: %.0f \n", math.Mod(10, 3));

	fmt.Printf("10/4 is : %.0f \n", math.Remainder(10, 4));

	//max values.
	fmt.Printf("Max int8 is: %d \n", math.MaxInt8)

	fmt.Printf("Max int16 is: %d \n", math.MaxInt16)

	fmt.Printf("Max int32 is: %d \n", math.MaxInt32)

	fmt.Printf("Max int64 is: %d \n", math.MaxInt64)

	fmt.Printf("Max float32 is: %f \n", math.MaxFloat32)




}
