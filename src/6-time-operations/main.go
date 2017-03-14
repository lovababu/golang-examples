package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=========== Date packge usage. ===============")
	var cTime = time.Now;  //time.Now returns Time object, where as time.Now() returns Time struct.
	fmt.Println("Current time is: " , cTime())

	//sleep the current execution.
	time.Sleep(time.Second)
	fmt.Println("After sleeping..")

	//Timer .
	c := time.Tick(1 * time.Second)
	var i int = 0;
	for now := range c {
		fmt.Printf("%v %s\n", now, statusUpdate())
		i += 1
		if i == 3 {
			break
		}
	}

	//seconds conversion.
	second := time.Second
	fmt.Println("seconds .. : ", int64(second/time.Millisecond))
	//add days to current time.
	var fTime = cTime().AddDate(0, 0, 2);

	fmt.Println("future time is: ", fTime)
	//time comparison.
	fmt.Println("is fTime is After cTime: ", fTime.After(cTime()))
	fmt.Println("is cTime is before fTime: ", cTime().Before(fTime))

	//fetch year, month and day from Time.
	var year, month, day = cTime().Date();
	fmt.Printf("Year: %d, Month: %d, Day: %d \n", year, month, day)

	//fetch hours, minutes and seconds from Time.
	var hours, min, sec = cTime().Clock();
	fmt.Printf("Hours: %d, Minutes: %d, Seconds: %d \n", hours, min, sec)

	//fetch day of the month.
	var dayInMonth = cTime().Day()
	fmt.Printf("day is : %d \n", dayInMonth)

	//week day of time.
	fmt.Printf("Week day: %s \n", cTime().Weekday());
	//time zone location.
	fmt.Printf("Time zone Location is: %s \n", cTime().Local().Location().String())

	var zone, offSet = cTime().Zone();
	fmt.Printf("Zone is: %s, %d \n", zone, offSet)
}

func statusUpdate() (string) {
	return "Hello, tick tick.."
}