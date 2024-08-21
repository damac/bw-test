package main

import (
	"brightwheel-interview/proto/device_rest"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type readingStore struct {
	values   readingValue
	timeList map[string]bool // List of times for dedupe
}

type readingValue struct {
	latest     string // Latest timestamp
	totalCount int32  // Count
}

var readingChannel = make(chan *device_rest.SomeMessage, 10) // Global, I could also wrap the handle functions for http

var dataStore = make(map[string]*readingStore) // readingStore listed by ID's

// Array identified by Device ID's storing a "readingStore"
// Receive the data, unmarshal it, send it over to be processed and stored

func main() {

	go processReading(readingChannel)
	http.HandleFunc("/v1/reading", receiveReadings)
	http.HandleFunc("/v1/latest/", returnLatest)
	http.HandleFunc("/v1/count/", returnCount)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}

func receiveReadings(writer http.ResponseWriter, request *http.Request) {
	parsedData := &device_rest.SomeMessage{}

	if request.Method == "POST" {
		data, err := io.ReadAll(request.Body)
		err = json.Unmarshal(data, parsedData)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
		}
		readingChannel <- parsedData
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func returnCount(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		deviceId := strings.TrimPrefix(request.URL.Path, "/v1/count/")
		if _, exists := dataStore[deviceId]; !exists {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		count := &device_rest.CumulativeCount{
			CumulativeCount: dataStore[deviceId].values.totalCount,
		}
		reply, _ := json.Marshal(count)
		_, _ = fmt.Fprintf(writer, string(reply))
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func returnLatest(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		deviceId := strings.TrimPrefix(request.URL.Path, "/v1/latest/")
		if _, exists := dataStore[deviceId]; !exists {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		latestTimestamp := &device_rest.LatestTimestamp{
			LatestTimestamp: dataStore[deviceId].values.latest,
		}

		reply, _ := json.Marshal(latestTimestamp)
		_, _ = fmt.Fprintf(writer, string(reply))

	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Send messages through channels to be processed and placed into memory
// Run this in the background as a go subroutine
func processReading(input chan *device_rest.SomeMessage) {
	for {
		deviceReading := <-input
		for i := 0; i < len(deviceReading.Readings); i++ {
			recordTime := deviceReading.Readings[i].Timestamp
			if _, known := dataStore[deviceReading.Id]; known { // Known device
				if _, found := dataStore[deviceReading.Id].timeList[recordTime]; !found { // New timestamp
					dataStore[deviceReading.Id].timeList[recordTime] = true // Set this timestamp to true
					if dataStore[deviceReading.Id].values.latest < recordTime {
						dataStore[deviceReading.Id].values.latest = recordTime // If later, set latest
					}
					dataStore[deviceReading.Id].values.totalCount = dataStore[deviceReading.Id].values.totalCount + deviceReading.Readings[i].Count
				} else {
					log.Println("duplicate timestamp detect, record dropped.", deviceReading.Id, recordTime)
				}
			} else {
				dataStore[deviceReading.Id] = &readingStore{
					timeList: make(map[string]bool),
				} // Create a new device and initialize the timeList map for it
				dataStore[deviceReading.Id].timeList[recordTime] = true // Add timestamp
				dataStore[deviceReading.Id].values.latest = recordTime  // We know its latest, its alone
				dataStore[deviceReading.Id].values.totalCount = dataStore[deviceReading.Id].values.totalCount + deviceReading.Readings[i].Count
			}

		}
	}
}
