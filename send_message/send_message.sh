#!/bin/bash

echo "Starting..."

sleep 5

echo "Sending Message to 54694"

sleep 5

# PHONE_NUMBER="54694"
# MESSAGE="schedule"

# osascript send_message.applescript "$PHONE_NUMBER" "$MESSAGE"

# Get current date in YYYY-MM-DD format
current_date=$(date +%Y-%m-%d)

# Get the next date in YYYY-MM-DD format
next_date=$(date -v+1d +%Y-%m-%d)

# Wait for 1 minute
# echo "Waiting for 2 minutes before executing the command..."
# sleep 120

# Run the imessage-exporter command with the calculated dates
# imessage-exporter -f txt -t "54694" -o "/Users/carlos/Desktop/schedule-to-calendar-heb/imessage_data" -s "$current_date" -e "$next_date"

../start.sh


echo "Schedule has been uploaded to Google Calendar :D"