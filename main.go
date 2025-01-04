package main

func main() {

	// run parsing process
	filepath := "/Users/carlos/Desktop/schedule-to-calendar-heb/imessage_data/54694.txt" // change to desired txt path
	extractAndOverwrite(filepath)
	data := parse(filepath)

	// run google calendar process
	run_calendar(data)

}