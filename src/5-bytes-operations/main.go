package main

import (
	"fmt"
	"bytes"
	//"strings"
)

func main()  {
	fmt.Println("============== Package Bytes usage ===============")
	var bArray1 = []byte("golang");
	var bArray2 = []byte("golang");
	fmt.Println("Size of bArray1 : ", bArray1)
	fmt.Println("Size of bArray2 : ", string(bArray2))

	//bytes.Equal
	fmt.Println("Is bArray1 and bArray2 equal: ", bytes.Equal(bArray1, bArray2))

	//bytes.Compare
	fmt.Println("Compare bArray1 and bArray2 equal: ", bytes.Compare(bArray1, []byte("programing")))

	//create byte array from string.
	var str string = "hello go.";
	var byteFromString  [10]byte;
	//copy method to copy string into byte array.
	copy(byteFromString[:], str)
	fmt.Printf("Byte from string: %s \n", byteFromString)

	//bytes.Contains.
	fmt.Println("Is bArray1 contains go :", bytes.Contains(bArray1, []byte("go")))
	fmt.Println("Is bArray1 contains go :", bytes.ContainsAny(bArray1, "pg"))

	//bytes.HasPrefix
	fmt.Println("Is bArray1 has prefix go: ", bytes.HasPrefix(bArray1, []byte("go")))

	//bytes.HasSuffix
	fmt.Println("Is bArray1 has suffix lang: ", bytes.HasSuffix(bArray1, []byte("lang")))

	//Buffer
	var buffer bytes.Buffer
	buffer.Write([]byte("Written to buffer."))
	fmt.Println("Buffered data as bytes: ", buffer)
	fmt.Println("Buffered data: ", buffer)

	//buffer len.
	fmt.Println("Buffer length: ", buffer.Len())

	fmt.Println("Buffer capacity: ", buffer.Cap())

	var dataRead, err = buffer.ReadString(byte('\n'))

	if err != nil {
		fmt.Println("Data read: ", dataRead)
	}
}
