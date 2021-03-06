package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Page struct {
	IP string `json:"ip"`
}
func getIP() string {
	pages := getPages()
	val := ""
	for _, p := range pages {
		val = p.IP
	}
	return val
}

func getPages() []Page {
	raw, err := ioutil.ReadFile("./File.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []Page
	json.Unmarshal(raw, &c)
	return c
}
