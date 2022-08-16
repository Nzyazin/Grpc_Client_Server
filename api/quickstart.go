package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const YOUR_ACCESS_TOKEN = "AIzaSyBnpdbchZIlEdQ-8oyS1vs5NORa6b2fXU0"

func MakeRequest(playlistId string) string {
	resp, err := http.Get(fmt.Sprintf("https://youtube.googleapis.com/youtube/v3/playlistItems?part=snippet%2CcontentDetails&maxResults=25&playlistId=%s&key=%s", playlistId, YOUR_ACCESS_TOKEN))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

	return string(body)
}
