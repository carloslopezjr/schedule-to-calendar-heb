package main

import (
	"fmt"
)


func main() {

	// run parsing process
	filepath := "/Users/carlos/Desktop/schedule-to-calendar-heb/imessage_data/54694.txt" // change to desired txt path
	data := parse(filepath)

	fmt.Println(data) // debugging purposes

	// run google calendar process




}