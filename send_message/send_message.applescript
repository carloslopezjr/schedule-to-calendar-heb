on run {targetPhoneNumber, messageText}
    tell application "Messages"
        set targetService to 1st service whose service type = SMS
        set targetBuddy to buddy targetPhoneNumber of targetService
        send messageText to targetBuddy
    end tell
end run
