package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	MakeRequest()
}

func MakeRequest() {
	resp, err := http.Get("https://youtube.googleapis.com/youtube/v3/playlistItems?playlistId=PLjyyS7xnf5wOkhrhWAo4rWLdnpHRBj6ep&key=AIzaSyBnpdbchZIlEdQ-8oyS1vs5NORa6b2fXU0")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
