package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RestResponse struct {
	Kind     string   `json:"kind"`
	Items    []Items  `json:"items"`
	Pageinfo PageInfo `json:"pageInfo"`
}

func (resp *RestResponse) for_print() {
	fmt.Println("Kind:", resp.Kind)
	fmt.Println("\n")
	for i := 0; i < len(resp.Items); i++ {
		fmt.Println("Item number:", i+1)
		fmt.Println("ID:", resp.Items[i].Id)
		fmt.Println("PublishedAt:", resp.Items[i].Snippet.PublishedAt)
		fmt.Println("Title:", resp.Items[i].Snippet.Title)
		fmt.Println("Url:", resp.Items[i].Snippet.Thumbnails.Default.Url)
		fmt.Println("\n")
	}

	fmt.Println("Number of videos from playlist:", len(resp.Items))
}

type PageInfo struct {
	TotalResults int `json:"totalResults"`
}

type Items struct {
	Id      string  `json:"id"`
	Snippet Snippet `json:"snippet"`
}

type Snippet struct {
	PublishedAt string     `json:"publishedAt"`
	Title       string     `json:"title"`
	Thumbnails  Thumbnails `json:"thumbnails"`
}

type Thumbnails struct {
	Default Default `json:"default"`
}

type Default struct {
	Url string `json:"url"`
}

const YOUR_ACCESS_TOKEN = "AIzaSyBnpdbchZIlEdQ-8oyS1vs5NORa6b2fXU0"

func MakeRequest(playlistId string) RestResponse {
	//var playlistId = "PLjyyS7xnf5wOkhrhWAo4rWLdnpHRBj6ep"

	resp, err := http.Get("https://youtube.googleapis.com/youtube/v3/playlistItems?part=snippet%2C%20contentDetails&maxResults=25&playlistId=" + playlistId + "&key=" + YOUR_ACCESS_TOKEN)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var obj RestResponse

	err1 := json.Unmarshal([]byte(body), &obj)
	if err != nil {
		fmt.Println(err1)
	} else {
		obj.for_print()
	}
	return obj
}
