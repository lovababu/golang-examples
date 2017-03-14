package main

import (
	"fmt"
)
//Variables demo.
func main() {

	// %T prints the type of variable.
	var intValue  int = 1;
	fmt.Printf("%T is type of intValue and value is %d \n", intValue, intValue);

	var intValue8 int8 = 1;
	fmt.Printf("%T is type of intValue8 and value is %d \n", intValue8, intValue8);

	var floatValue  float32 = 1.2345;
	fmt.Printf("%T is type of floatValue and value is %.2f \n", floatValue, floatValue);

	var floatValue64 float64 = 1.2345;
	fmt.Printf("%T is type of floatValue64 and value is %.3f \n", floatValue64, floatValue64);

	//If there is no explicit type, go decide the type based on the value assigned.
	var noTypeInt = 5;
	fmt.Printf("%T is type of noTypeInt and value is %d \n", noTypeInt, noTypeInt);

	var noTypeFloat = 5.45678;
	fmt.Printf("%T is type of noTypeFloat and value is %.3f \n", noTypeFloat, noTypeFloat);

	var strValue  string = "Hello, World..!";
	fmt.Printf("%T is type of strValue and value is %s \n", strValue, strValue);

	var noTypeStrValue = "Hello, World..!";
	fmt.Printf("%T is type of noTypeStrValue and value is %s \n", noTypeStrValue, noTypeStrValue);

}
