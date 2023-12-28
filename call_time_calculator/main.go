package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Call struct {
	From      string
	To        string
	StartTime time.Time
	EndTime   time.Time
}

func main() {
	// Open the file
	file, err := os.Open("call_time_calculator/call_details.csv")
	if err != nil {
		log.Println("cannot open the file: ", err)
		return
	}
	// Defer file.Close
	defer file.Close()

	// Read the CSV and parse into the []Call
	reader := csv.NewReader(file)
	var calls []Call

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println("cannot read the file: ", err)
			return
		}

		startTime, _ := time.Parse(time.RFC3339, record[2])
		endTime, _ := time.Parse(time.RFC3339, record[3])

		calls = append(calls, Call{
			From:      record[0],
			To:        record[1],
			StartTime: startTime,
			EndTime:   endTime,
		})
	}

	// Check the name and calculate the actual time

	name := "john"

	var totalIncoming time.Duration
	var totalOutgoing time.Duration

	for _, call := range calls {
		if strings.EqualFold(call.From, name) {
			totalOutgoing += call.EndTime.Sub(call.StartTime)
		} else if strings.EqualFold(call.To, name) {
			totalIncoming += call.EndTime.Sub(call.StartTime)
		}
	}

	fmt.Printf("totalIncoming: %v\n", totalIncoming)
	fmt.Printf("totalOutgoing: %v\n", totalOutgoing)
}
