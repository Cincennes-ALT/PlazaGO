package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Wellcome to PlazaGO!")
	GetAPI("status", "")
}

// api sample
func GetAPI(WhatType string, WhatValue string) { // the link to get api
	url := "https://api.plaza.one/" + WhatType + "/" + WhatValue // link puzzle
	method := "GET"
	client := &http.Client{}

	// send the request
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	//get the request
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	// body of the request
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// switch case based on the link
	switch WhatType {
	case "status":
		var prepare map[string]map[string]string // its a map in a map sor the "song" key
		json.Unmarshal(body, &prepare)
		var fwdSong map[string]string = prepare["song"]
		currentSong(fwdSong)
	}

}

func currentSong(gotSatus map[string]string) {
	// i kinda made that too complex
	var songMeta SongNow = SongNow{fmt.Sprint(gotSatus["title"]), fmt.Sprint(gotSatus["artist"]), fmt.Sprint(gotSatus["album"]), fmt.Sprint(gotSatus["artwork_src"])}

	title := songMeta.title
	artist := songMeta.artist
	album := songMeta.album
	// albumCover := songMeta.albumCover

	fmt.Printf("Currently playing: %s by %s from %s\n", title, artist, album)
}

/*
	structs
*/
// Now playing sungs structs
type SongNow struct {
	artist, title, album, albumCover string
}
