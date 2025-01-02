package main

// read imessage_data file and parse for specific text

import (
	"fmt"
	"os"
	"log"
	"strings"
	"bufio"
)

// create a struct that will be created each time 
type Event struct {
	Number string
	Day string
	Start_hour string
	Start_minute string
	S_AM_PM bool // false means am, true means pm
	End_hour string
	End_minute string
	E_AM_PM bool // false means am, true means pm
}


func parse(path string) []interface{} {
	// Read the entire file into memory
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var dates []interface{}

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(strings.NewReader(string(data)))

	// Iterate over each line
	for scanner.Scan() {
		line := scanner.Text()

		// Apply your filtering logic here
		if strings.Contains(line, "MO") && !strings.Contains(line, "OFF") {

			var S_AM_PM_bool bool
			var E_AM_PM_bool bool

			if line[10:11] == "P" {
				S_AM_PM_bool = true
			} else {
				S_AM_PM_bool = false
			}

			if line[16:17] == "P" {
				E_AM_PM_bool = true
			} else {
				E_AM_PM_bool = false
			}

			mon := Event {
				Number: line[0:2],
				Day: line[3:5],
				Start_hour: line[6:7],
				Start_minute: line[8:10],
				S_AM_PM: S_AM_PM_bool,
				End_hour: line[12:13],
				End_minute: line[14:16],
				E_AM_PM: E_AM_PM_bool,
			}
			dates = append(dates, mon)
		}
		// Apply your filtering logic here
		if strings.Contains(line, "TU") && !strings.Contains(line, "OFF"){
			var S_AM_PM_bool bool
			var E_AM_PM_bool bool

			if line[10:11] == "P" {
				S_AM_PM_bool = true
			} else {
				S_AM_PM_bool = false
			}

			if line[16:17] == "P" {
				E_AM_PM_bool = true
			} else {
				E_AM_PM_bool = false
			}

			tue := Event {
				Number: line[0:2],
				Day: line[3:5],
				Start_hour: line[6:7],
				Start_minute: line[8:10],
				S_AM_PM: S_AM_PM_bool,
				End_hour: line[12:13],
				End_minute: line[14:16],
				E_AM_PM: E_AM_PM_bool,
			}
			dates = append(dates, tue)
		}
		// Apply your filtering logic here
		if strings.Contains(line, "WE") && !strings.Contains(line, "OFF"){
			var S_AM_PM_bool bool
			var E_AM_PM_bool bool

			if line[10:11] == "P" {
				S_AM_PM_bool = true
			} else {
				S_AM_PM_bool = false
			}

			if line[16:17] == "P" {
				E_AM_PM_bool = true
			} else {
				E_AM_PM_bool = false
			}

			wed := Event {
				Number: line[0:2],
				Day: line[3:5],
				Start_hour: line[6:7],
				Start_minute: line[8:10],
				S_AM_PM: S_AM_PM_bool,
				End_hour: line[12:13],
				End_minute: line[14:16],
				E_AM_PM: E_AM_PM_bool,
			}
			dates = append(dates, wed)
		}
		// Apply your filtering logic here
		if strings.Contains(line, "TH") && !strings.Contains(line, "OFF"){
			var S_AM_PM_bool bool
			var E_AM_PM_bool bool

			if line[10:11] == "P" {
				S_AM_PM_bool = true
			} else {
				S_AM_PM_bool = false
			}

			if line[16:17] == "P" {
				E_AM_PM_bool = true
			} else {
				E_AM_PM_bool = false
			}

			thu := Event {
				Number: line[0:2],
				Day: line[3:5],
				Start_hour: line[6:7],
				Start_minute: line[8:10],
				S_AM_PM: S_AM_PM_bool,
				End_hour: line[12:13],
				End_minute: line[14:16],
				E_AM_PM: E_AM_PM_bool,
			}
			dates = append(dates, thu)
		}
		// Apply your filtering logic here
		if strings.Contains(line, "FR") && !strings.Contains(line, "OFF"){
			var S_AM_PM_bool bool
			var E_AM_PM_bool bool

			if line[10:11] == "P" {
				S_AM_PM_bool = true
			} else {
				S_AM_PM_bool = false
			}

			if line[16:17] == "P" {
				E_AM_PM_bool = true
			} else {
				E_AM_PM_bool = false
			}

			fri := Event {
				Number: line[0:2],
				Day: line[3:5],
				Start_hour: line[6:7],
				Start_minute: line[8:10],
				S_AM_PM: S_AM_PM_bool,
				End_hour: line[12:13],
				End_minute: line[14:16],
				E_AM_PM: E_AM_PM_bool,
			}
			dates = append(dates, fri)
		}
		// Apply your filtering logic here
		if strings.Contains(line, "SA") && !strings.Contains(line, "OFF"){
			var S_AM_PM_bool bool
			var E_AM_PM_bool bool

			if line[10:11] == "P" {
				S_AM_PM_bool = true
			} else {
				S_AM_PM_bool = false
			}

			if line[16:17] == "P" {
				E_AM_PM_bool = true
			} else {
				E_AM_PM_bool = false
			}

			sat := Event {
				Number: line[0:2],
				Day: line[3:5],
				Start_hour: line[6:7],
				Start_minute: line[8:10],
				S_AM_PM: S_AM_PM_bool,
				End_hour: line[12:13],
				End_minute: line[14:16],
				E_AM_PM: E_AM_PM_bool,
			}
			dates = append(dates, sat)
		}
		// Apply your filtering logic here
		if strings.Contains(line, "SU") && !strings.Contains(line, "OFF"){
			var S_AM_PM_bool bool
			var E_AM_PM_bool bool

			if line[10:11] == "P" {
				S_AM_PM_bool = true
			} else {
				S_AM_PM_bool = false
			}

			if line[16:17] == "P" {
				E_AM_PM_bool = true
			} else {
				E_AM_PM_bool = false
			}

			sun := Event {
				Number: line[0:2],
				Day: line[3:5],
				Start_hour: line[6:7],
				Start_minute: line[8:10],
				S_AM_PM: S_AM_PM_bool,
				End_hour: line[12:13],
				End_minute: line[14:16],
				E_AM_PM: E_AM_PM_bool,
			}
			dates = append(dates, sun)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
	return dates
}