package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("========= Go Routines ===========")
	go async1(10)
	go async2(10)
	time.Sleep(15 * time.Second)
	fmt.Println("Main implicit routine exiting.")
}

func async1(num int)  {
	for i:=0; i < num; i++ {
		fmt.Println("Method async 1 running...")
	}
}

func async2(num int)  {
	for i:=0; i < num; i++ {
		fmt.Println("Method async 2 running...")
	}
}
