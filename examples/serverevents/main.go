package main

import (
	"encoding/json"
	"fmt"
	mgo "go-marathon"
	"os"
)

func pretty(value interface{}) string {
	prettyBytes, _ := json.MarshalIndent(value, "", "\t")
	return string(prettyBytes)
}

func main() {
	config := mgo.NewDefaultConfig()
	config.URL = "http://10.131.163.183:8080"
	config.LogOutput = os.Stdout
	if client, err := mgo.NewClient(config); err != nil {
		fmt.Printf("Something went wrong %v", err.Error())
	} else {
		if client.ListenEvents() != nil {
			fmt.Printf("Something went wrong %v", err.Error())
		} else {
			for {
				fmt.Println(pretty(<-client.GetEvents()))
			}
		}
	}
}
