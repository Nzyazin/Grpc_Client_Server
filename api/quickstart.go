package api

import (
	"io/ioutil"
	"log"
	"net/http"
)

const YOUR_ACCESS_TOKEN = "AIzaSyBnpdbchZIlEdQ-8oyS1vs5NORa6b2fXU0"

func MakeRequest(playlistId string) string {
	//var playlistId = "PLjyyS7xnf5wOkhrhWAo4rWLdnpHRBj6ep"
	//make request to youtube playlist
	resp, err := http.Get("https://youtube.googleapis.com/youtube/v3/playlistItems?part=snippet%2C%20contentDetails&maxResults=25&playlistId=" + playlistId + "&key=" + YOUR_ACCESS_TOKEN)
	if err != nil {
		log.Fatalln(err)
	}

	//reads answer and returns bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//returns formatted bytes
	return string(body)
}
