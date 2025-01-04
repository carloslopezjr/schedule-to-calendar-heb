package main

// read imessage_data file and parse for specific text

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

// create a struct that will be created each time
type Event struct {
	Number       string
	Day          string
	Start_hour   string
	Start_minute string
	S_AM_PM      bool // false means am, true means pm
	End_hour     string
	End_minute   string
	E_AM_PM      bool // false means am, true means pm
}

// Function to extract the most recent response from "54694" and overwrite the file
func extractAndOverwrite(filename string) error {
	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Variables to hold the latest response
	var latestResponse []string
	var tempResponse []string
	inResponseBlock := false

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Check for "54694" to identify a response block
		if strings.Contains(line, "54694") {
			// Start of a new response block
			inResponseBlock = true
			tempResponse = []string{line}
		} else if inResponseBlock {
			// Collect lines within the response block
			if strings.TrimSpace(line) == "" {
				// End of the block
				inResponseBlock = false
				latestResponse = tempResponse
			} else {
				tempResponse = append(tempResponse, line)
			}
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	// If no response was found, return an error
	if len(latestResponse) == 0 {
		return fmt.Errorf("no response from 54694 found in the file")
	}

	// Open the file for writing (truncate it first)
	file, err = os.Create(filename)
	if err != nil {
		return fmt.Errorf("error opening file for writing: %v", err)
	}
	defer file.Close()

	// Write the latest response to the file
	for _, line := range latestResponse {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("error writing to file: %v", err)
		}
	}

	return nil
}

func parse(path string) []Event {
	// Read the entire file into memory
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var dates []Event

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(strings.NewReader(string(data)))

	// Iterate over each line
	for scanner.Scan() {
		line := scanner.Text()

		// Apply your filtering logic here
		if strings.Contains(line, "MO") && !strings.Contains(line, "OFF") {

			var S_AM_PM_bool bool
			var E_AM_PM_bool bool

			// Split by spaces
			parts := strings.Fields(line) // "06", "MO", "5:45A-12:45P"

			// Extract individual parts
			firstTwoDigits := parts[0] // "06"
			day := parts[1]            // "MO"
			timeRange := parts[2]      // "5:45A-12:45P"

			// Split time range further
			times := strings.Split(timeRange, "-") // ["5:45A", "12:45P"]

			// Function to extract hour, minute, and AM/PM from a time string
			extractTime := func(time string) (string, string, string) {
				hour := ""
				minute := ""
				ampm := ""

				// Extract hour and minute
				for i, char := range time {
					if unicode.IsDigit(char) {
						hour += string(char)
					} else if char == ':' {
						minute = time[i+1 : i+3] // Get minutes after ":"
						ampm = time[i+3:]        // Get AM/PM part
						break
					}
				}

				return hour, minute, ampm
			}

			// Extract times
			startHour, startMinute, startAMPM := extractTime(times[0]) // "5:45A"
			endHour, endMinute, endAMPM := extractTime(times[1])       // "12:45P"

			if startAMPM == "P" {
				S_AM_PM_bool = true
			} else {
				S_AM_PM_bool = false
			}

			if endAMPM == "P" {
				E_AM_PM_bool = true
			} else {
				E_AM_PM_bool = false
			}

			mon := Event{
				Number:       firstTwoDigits,
				Day:          day,
				Start_hour:   startHour, 
				Start_minute: startMinute,
				S_AM_PM:      S_AM_PM_bool,
				End_hour:     endHour,
				End_minute:   endMinute,
				E_AM_PM:      E_AM_PM_bool,
			}
			dates = append(dates, mon)
		}
		// Apply your filtering logic here
		if strings.Contains(line, "TU") && !strings.Contains(line, "OFF") {
			var S_AM_PM_bool bool
			var E_AM_PM_bool bool

			// Split by spaces
			parts := strings.Fields(line) // "06", "MO", "5:45A-12:45P"

			// Extract individual parts
			firstTwoDigits := parts[0] // "06"
			day := parts[1]            // "MO"
			timeRange := parts[2]      // "5:45A-12:45P"

			// Split time range further
			times := strings.Split(timeRange, "-") // ["5:45A", "12:45P"]

			// Function to extract hour, minute, and AM/PM from a time string
			extractTime := func(time string) (string, string, string) {
				hour := ""
				minute := ""
				ampm := ""

				// Extract hour and minute
				for i, char := range time {
					if unicode.IsDigit(char) {
						hour += string(char)
					} else if char == ':' {
						minute = time[i+1 : i+3] // Get minutes after ":"
						ampm = time[i+3:]        // Get AM/PM part
						break
					}
				}

				return hour, minute, ampm
			}

			// Extract times
			startHour, startMinute, startAMPM := extractTime(times[0]) // "5:45A"
			endHour, endMinute, endAMPM := extractTime(times[1])       // "12:45P"

			if startAMPM == "P" {
				S_AM_PM_bool = true
			} else {
				S_AM_PM_bool = false
			}

			if endAMPM == "P" {
				E_AM_PM_bool = true
			} else {
				E_AM_PM_bool = false
			}

			tue := Event{
				Number:       firstTwoDigits,
				Day:          day,
				Start_hour:   startHour, // this needs to be changed. If there's two digits in the hour, it will bug out
				Start_minute: startMinute,
				S_AM_PM:      S_AM_PM_bool,
				End_hour:     endHour,
				End_minute:   endMinute,
				E_AM_PM:      E_AM_PM_bool,
			}
			dates = append(dates, tue)
		}
		// Apply your filtering logic here
		if strings.Contains(line, "WE") && !strings.Contains(line, "OFF") {
			var S_AM_PM_bool bool
			var E_AM_PM_bool bool

			// Split by spaces
			parts := strings.Fields(line) // "06", "MO", "5:45A-12:45P"

			// Extract individual parts
			firstTwoDigits := parts[0] // "06"
			day := parts[1]            // "MO"
			timeRange := parts[2]      // "5:45A-12:45P"

			// Split time range further
			times := strings.Split(timeRange, "-") // ["5:45A", "12:45P"]

			// Function to extract hour, minute, and AM/PM from a time string
			extractTime := func(time string) (string, string, string) {
				hour := ""
				minute := ""
				ampm := ""

				// Extract hour and minute
				for i, char := range time {
					if unicode.IsDigit(char) {
						hour += string(char)
					} else if char == ':' {
						minute = time[i+1 : i+3] // Get minutes after ":"
						ampm = time[i+3:]        // Get AM/PM part
						break
					}
				}

				return hour, minute, ampm
			}

			// Extract times
			startHour, startMinute, startAMPM := extractTime(times[0]) // "5:45A"
			endHour, endMinute, endAMPM := extractTime(times[1])       // "12:45P"

			if startAMPM == "P" {
				S_AM_PM_bool = true
			} else {
				S_AM_PM_bool = false
			}

			if endAMPM == "P" {
				E_AM_PM_bool = true
			} else {
				E_AM_PM_bool = false
			}

			wed := Event{
				Number:       firstTwoDigits,
				Day:          day,
				Start_hour:   startHour, // this needs to be changed. If there's two digits in the hour, it will bug out
				Start_minute: startMinute,
				S_AM_PM:      S_AM_PM_bool,
				End_hour:     endHour,
				End_minute:   endMinute,
				E_AM_PM:      E_AM_PM_bool,
			}
			dates = append(dates, wed)
		}
		// Apply your filtering logic here
		if strings.Contains(line, "TH") && !strings.Contains(line, "OFF") {
			var S_AM_PM_bool bool
			var E_AM_PM_bool bool

			// Split by spaces
			parts := strings.Fields(line) // "06", "MO", "5:45A-12:45P"

			// Extract individual parts
			firstTwoDigits := parts[0] // "06"
			day := parts[1]            // "MO"
			timeRange := parts[2]      // "5:45A-12:45P"

			// Split time range further
			times := strings.Split(timeRange, "-") // ["5:45A", "12:45P"]

			// Function to extract hour, minute, and AM/PM from a time string
			extractTime := func(time string) (string, string, string) {
				hour := ""
				minute := ""
				ampm := ""

				// Extract hour and minute
				for i, char := range time {
					if unicode.IsDigit(char) {
						hour += string(char)
					} else if char == ':' {
						minute = time[i+1 : i+3] // Get minutes after ":"
						ampm = time[i+3:]        // Get AM/PM part
						break
					}
				}

				return hour, minute, ampm
			}

			// Extract times
			startHour, startMinute, startAMPM := extractTime(times[0]) // "5:45A"
			endHour, endMinute, endAMPM := extractTime(times[1])       // "12:45P"

			if startAMPM == "P" {
				S_AM_PM_bool = true
			} else {
				S_AM_PM_bool = false
			}

			if endAMPM == "P" {
				E_AM_PM_bool = true
			} else {
				E_AM_PM_bool = false
			}

			thu := Event{
				Number:       firstTwoDigits,
				Day:          day,
				Start_hour:   startHour, // this needs to be changed. If there's two digits in the hour, it will bug out
				Start_minute: startMinute,
				S_AM_PM:      S_AM_PM_bool,
				End_hour:     endHour,
				End_minute:   endMinute,
				E_AM_PM:      E_AM_PM_bool,
			}
			dates = append(dates, thu)
		}
		// Apply your filtering logic here
		if strings.Contains(line, "FR") && !strings.Contains(line, "OFF") {
			var S_AM_PM_bool bool
			var E_AM_PM_bool bool

			// Split by spaces
			parts := strings.Fields(line) // "06", "MO", "5:45A-12:45P"

			// Extract individual parts
			firstTwoDigits := parts[0] // "06"
			day := parts[1]            // "MO"
			timeRange := parts[2]      // "5:45A-12:45P"

			// Split time range further
			times := strings.Split(timeRange, "-") // ["5:45A", "12:45P"]

			// Function to extract hour, minute, and AM/PM from a time string
			extractTime := func(time string) (string, string, string) {
				hour := ""
				minute := ""
				ampm := ""

				// Extract hour and minute
				for i, char := range time {
					if unicode.IsDigit(char) {
						hour += string(char)
					} else if char == ':' {
						minute = time[i+1 : i+3] // Get minutes after ":"
						ampm = time[i+3:]        // Get AM/PM part
						break
					}
				}

				return hour, minute, ampm
			}

			// Extract times
			startHour, startMinute, startAMPM := extractTime(times[0]) // "5:45A"
			endHour, endMinute, endAMPM := extractTime(times[1])       // "12:45P"

			if startAMPM == "P" {
				S_AM_PM_bool = true
			} else {
				S_AM_PM_bool = false
			}

			if endAMPM == "P" {
				E_AM_PM_bool = true
			} else {
				E_AM_PM_bool = false
			}

			fri := Event{
				Number:       firstTwoDigits,
				Day:          day,
				Start_hour:   startHour, // this needs to be changed. If there's two digits in the hour, it will bug out
				Start_minute: startMinute,
				S_AM_PM:      S_AM_PM_bool,
				End_hour:     endHour,
				End_minute:   endMinute,
				E_AM_PM:      E_AM_PM_bool,
			}
			dates = append(dates, fri)
		}
		// Apply your filtering logic here
		if strings.Contains(line, "SA") && !strings.Contains(line, "OFF") {
			var S_AM_PM_bool bool
			var E_AM_PM_bool bool

			// Split by spaces
			parts := strings.Fields(line) // "06", "MO", "5:45A-12:45P"

			// Extract individual parts
			firstTwoDigits := parts[0] // "06"
			day := parts[1]            // "MO"
			timeRange := parts[2]      // "5:45A-12:45P"

			// Split time range further
			times := strings.Split(timeRange, "-") // ["5:45A", "12:45P"]

			// Function to extract hour, minute, and AM/PM from a time string
			extractTime := func(time string) (string, string, string) {
				hour := ""
				minute := ""
				ampm := ""

				// Extract hour and minute
				for i, char := range time {
					if unicode.IsDigit(char) {
						hour += string(char)
					} else if char == ':' {
						minute = time[i+1 : i+3] // Get minutes after ":"
						ampm = time[i+3:]        // Get AM/PM part
						break
					}
				}

				return hour, minute, ampm
			}

			// Extract times
			startHour, startMinute, startAMPM := extractTime(times[0]) // "5:45A"
			endHour, endMinute, endAMPM := extractTime(times[1])       // "12:45P"

			if startAMPM == "P" {
				S_AM_PM_bool = true
			} else {
				S_AM_PM_bool = false
			}

			if endAMPM == "P" {
				E_AM_PM_bool = true
			} else {
				E_AM_PM_bool = false
			}

			sat := Event{
				Number:       firstTwoDigits,
				Day:          day,
				Start_hour:   startHour, // this needs to be changed. If there's two digits in the hour, it will bug out
				Start_minute: startMinute,
				S_AM_PM:      S_AM_PM_bool,
				End_hour:     endHour,
				End_minute:   endMinute,
				E_AM_PM:      E_AM_PM_bool,
			}
			dates = append(dates, sat)
		}
		// Apply your filtering logic here
		if strings.Contains(line, "SU") && !strings.Contains(line, "OFF") {
			var S_AM_PM_bool bool
			var E_AM_PM_bool bool

			// Split by spaces
			parts := strings.Fields(line) // "06", "MO", "5:45A-12:45P"

			// Extract individual parts
			firstTwoDigits := parts[0] // "06"
			day := parts[1]            // "MO"
			timeRange := parts[2]      // "5:45A-12:45P"

			// Split time range further
			times := strings.Split(timeRange, "-") // ["5:45A", "12:45P"]

			// Function to extract hour, minute, and AM/PM from a time string
			extractTime := func(time string) (string, string, string) {
				hour := ""
				minute := ""
				ampm := ""

				// Extract hour and minute
				for i, char := range time {
					if unicode.IsDigit(char) {
						hour += string(char)
					} else if char == ':' {
						minute = time[i+1 : i+3] // Get minutes after ":"
						ampm = time[i+3:]        // Get AM/PM part
						break
					}
				}

				return hour, minute, ampm
			}

			// Extract times
			startHour, startMinute, startAMPM := extractTime(times[0]) // "5:45A"
			endHour, endMinute, endAMPM := extractTime(times[1])       // "12:45P"

			if startAMPM == "P" {
				S_AM_PM_bool = true
			} else {
				S_AM_PM_bool = false
			}

			if endAMPM == "P" {
				E_AM_PM_bool = true
			} else {
				E_AM_PM_bool = false
			}

			sun := Event{
				Number:       firstTwoDigits,
				Day:          day,
				Start_hour:   startHour, // this needs to be changed. If there's two digits in the hour, it will bug out
				Start_minute: startMinute,
				S_AM_PM:      S_AM_PM_bool,
				End_hour:     endHour,
				End_minute:   endMinute,
				E_AM_PM:      E_AM_PM_bool,
			}
			dates = append(dates, sun)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
	return dates
}
