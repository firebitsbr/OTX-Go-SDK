// Copyright 2015 The go-otxapi AUTHORS. All rights reserved.
//
// Use of this source code is governed by an Apache
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"otxapi"
	"os"
)

func main() {
	// you can set your api key environment variable however you prefer..Including inline!

	os.Setenv("X_OTX_API_KEY", "mysecretkey")
	fmt.Println("Found API Key in environment var X_OTX_API_KEY! Validating key...")
	fmt.Println(fmt.Sprintf("%s", os.Getenv("X_OTX_API_KEY")))
	client := otxapi.NewClient(nil)
	// get user details, easy way to validate api key
	user_detail, _, err := client.UserDetail.Get()
	if err != nil {
		fmt.Printf("Fatal error: %v\n\n", err)
		fmt.Println("Please ensure your API KEY is valid (i.e. `echo $X_OTX_API_KEY`)")
		os.Exit(1)
	} else {
		fmt.Printf("Successfully authenticated with api key as: %v\n\n", otxapi.Stringify(user_detail.Username))
	}

	// now let's try getting some IOCs
	pulse_detail, _, err := client.PulseDetail.Get("56cdb0a04637f275671672f3")
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	} else {
		fmt.Printf("%v\n\n\n\n", otxapi.Stringify(pulse_detail))
	}
	opt := &otxapi.ListOptions{Page: 1, PerPage: 4}
	pulse_list, _, err := client.ThreatIntel.List(opt)

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	} else {
		fmt.Printf("%v\n\n\n\n", otxapi.Stringify(pulse_list))
	}
}
