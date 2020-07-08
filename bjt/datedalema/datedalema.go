package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func dateDalema() {

	// what is expected to happend at the front-end
	fmt.Println("ui==================================================")

	/// we get the date from time, in the form: 2006-01-02 from an HTML
	dateStr := "2020-05-31"
	fmt.Println("Date from HTML:", dateStr)

	/// we convert that to a time date, so that it can be validated, and sent
	/// we parse this into the a request struct so that we can send it
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Printf("failed to convert input date string \"%s\" to time object. \nerr: %v ", dateStr, err)
		return
	}

	fmt.Println()
	fmt.Printf("time object, converted from time string: \ndate string: %s \ntime object: %v", dateStr, date)

	// what happens at the back-end
	fmt.Println()
	fmt.Println("ui==================================================")

	/// we get the date from the request in the RFC3339 format
	/// is will come in a JSON format as a string as part of an object
	/// the date will be isolated, converted in to a time object and sent to the database to be stored
	/// at this point I realised, the date should exist in the database as string. Apart from the HTML, nothing uses it.
	type requestBody struct {
		Date time.Time
	}
	reqBody := requestBody{
		Date: date,
	}
	fmt.Println()
	fmt.Println("Struct to be marshalled and sent")
	fmt.Println(reqBody)

	/// assuming this json was sent
	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Error marshalling json, trying to simulate: request body sent: ", err)
		return
	}
	fmt.Println()
	fmt.Println("JSON request sent and received")
	fmt.Println(string(reqBodyJSON))

	/// assuming the same json was received, and we are unmarshalling it
	type requestBodyWithDateAsString struct {
		Date string
	}
	var reqBodyUmarshal requestBodyWithDateAsString
	err = json.Unmarshal(reqBodyJSON, &reqBodyUmarshal)
	if err != nil {
		fmt.Println("Error unmarshalling json, trying to simulate: request body received: ", err)
		return
	}
	fmt.Println()
	fmt.Println("Struct unmarshalled from received JSON")
	fmt.Println(reqBodyUmarshal)

	rcvd_dateStr := reqBodyUmarshal.Date
	fmt.Println()
	fmt.Println("Date string from the unmarshalled struct")
	fmt.Println(rcvd_dateStr)

	/// convert a1 to time, final struct that goes into db
	dbDate, err := time.Parse(time.RFC3339, rcvd_dateStr)
	if err != nil {
		fmt.Println("Error parsing string to date that would be store into the database")
		fmt.Println("This is the same date in the struct after unmarshal")
		fmt.Println("err :", err)
		return
	}
	fmt.Println()
	fmt.Println("Final date to be send to db: ", dbDate)
}
