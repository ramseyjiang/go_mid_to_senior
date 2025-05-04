package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const logFileName = "logs.txt"

// LogEntry is a struct for a log entry
type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

// receiveLog receives a log message from the API service and appends it to the log file
func receiveLog(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	logMessage := string(body)
	if len(logMessage) == 0 {
		fmt.Println("Received log content length is:", len(logMessage))
	} else {
		fmt.Println("Received log:", logMessage)
	}

	logEntry := createLogEntry(logMessage)
	err := appendLogEntryToFile(logEntry, logFileName)
	if err != nil {
		fmt.Printf("Error appending log entry to file: %v\n", err)
	}
}

// createLogEntry creates a LogEntry struct from a log message
func createLogEntry(logMessage string) LogEntry {
	return LogEntry{
		Timestamp: time.Now(),
		Message:   logMessage,
	}
}

// appendLogEntryToFile appends a log entry to the log file
func appendLogEntryToFile(logEntry LogEntry, logFileName string) error {
	// Open the log file in append mode, creating it if it doesn't exist
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening log file: %v", err)
	}

	// Marshal the log entry to JSON
	logEntryJSON, err := json.Marshal(logEntry)
	if err != nil {
		return fmt.Errorf("error marshaling log entry: %v", err)
	}

	// Write the JSON log entry to the file, followed by a newline
	_, err = logFile.WriteString(string(logEntryJSON) + "\n")
	if err != nil {
		return fmt.Errorf("error writing log entry to file: %v", err)
	}

	return nil
}

func main() {
	http.Handle("/log/add", http.HandlerFunc(receiveLog))

	log.Fatal(http.ListenAndServe(":8081", nil))
}
