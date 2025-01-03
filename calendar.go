package main

// import google calendar api, make calls to create eventspackage main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	// Generate the authorization URL
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser: \n%v\n", authURL)

	// Manually instruct the user to copy the authorization code
	fmt.Println("After granting access, paste the authorization code below:")

	var authCode string
	fmt.Scan(&authCode) // Read the authorization code from the terminal

	// Exchange the authorization code for a token
	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// Helper function to convert your Event time to ISO 8601 DateTime
func convertToISO8601(number string, startHour, startMinute string, isPM bool) string {
	// Convert day abbreviation to a full date (mocked for simplicity)
	// You should replace this with proper date handling logic
	currentYear := time.Now().Year()
	currentMonth := time.Now().Month()
	currentDay := time.Now().Day() // only use to check if today is the first
	if number == "1" && currentDay != 1{ // if it's not the first day of the month, that means the week bleeds to a new month
		if currentMonth == 12 { 
			currentMonth = 1 // reset to 1 if on last month of the year (12)
		} else {
			currentMonth += 1 // increment to next month

		}
	}

	date := fmt.Sprintf("%d-%d-%.2s", currentYear, currentMonth, number) // Assuming dates are in November for now

	hour, minute := startHour, startMinute
	if len(hour) == 1 {
		hour = "0" + hour
	}
	if len(minute) == 1 {
		minute = "0" + minute
	}

	hourInt := 0
	fmt.Sscanf(hour, "%d", &hourInt)
	if isPM && hourInt != 12 {
		hourInt += 12
	} else if !isPM && hourInt == 12 {
		hourInt = 0
	}
	hour = fmt.Sprintf("%02d", hourInt)

	return fmt.Sprintf("%sT%s:%s:00-06:00", date, hour, minute) // Mocked timezone as -07:00
}

// Function to insert an event into Google Calendar
func insertEvent(srv *calendar.Service, event Event) {
	startTime := convertToISO8601(event.Number, event.Start_hour, event.Start_minute, event.S_AM_PM)
	endTime := convertToISO8601(event.Number, event.End_hour, event.End_minute, event.E_AM_PM)
	fmt.Println(startTime)
	fmt.Println(endTime)
	calendarEvent := &calendar.Event{
		Summary:     fmt.Sprintf("H-E-B (Work)"),
		Location:    "{insert location}", // Example location
		Description: fmt.Sprintf("Generated event for %s", event.Day),
		Start: &calendar.EventDateTime{
			DateTime: startTime,
			TimeZone: "America/Chicago",
		},
		End: &calendar.EventDateTime{
			DateTime: endTime,
			TimeZone: "America/Chicago",
		},
	}

	calendarId := "bec729f22cb6cf6a349c03a3a4e8aeb738b7c6fe7c760c7c9d315e6684e2e003@group.calendar.google.com"
	createdEvent, err := srv.Events.Insert(calendarId, calendarEvent).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	}
	fmt.Printf("Event created: %s\n", createdEvent.HtmlLink)
}

func run_calendar(data []Event) {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope, calendar.CalendarEventsScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	fmt.Println(data)

	// this is where code will start
	for _, info := range data {
		insertEvent(srv, info)
	}
		

	
}
