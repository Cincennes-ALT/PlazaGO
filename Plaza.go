package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Wellcome to PlazaGO!")

	CurrentSong()
}

// gets the current song
func CurrentSong() {
	getAPIEredmeny := json.Marshal(GetAPI("status", "", "json"))

}

// api sample
func GetAPI(WhatType string, WhatValue string, ReturnType string) string {
	url := "https://api.plaza.one/" + WhatType + "/" + WhatValue
	method := "GET"
	client := &http.Client{}

	request, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		var hiba error = err
		return (hiba.Error())
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		var hiba error = err
		return (hiba.Error())
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		var hiba error = err
		return (hiba.Error())
	}

	// fmt.Println(string(body))
	switch ReturnType {
	case "string":
		vissza := string(body)
		return vissza
	case "array":
		var vissza [1]string
		vissza[0] = string(body)
		return vissza[0]
	case "json":
		vissza, err := json.Marshal(body)
		if err != nil {
			fmt.Println(err)
		}
		return string(vissza)
	default:
		return string(body)
	}
}
