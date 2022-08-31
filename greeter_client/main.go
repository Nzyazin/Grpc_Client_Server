/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package greeter_client

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

type RestResponse struct {
	Kind     string   `json:"kind"`
	Items    []Items  `json:"items"`
	Pageinfo PageInfo `json:"pageInfo"`
}

func (resp *RestResponse) for_print() string {
	//var allInclusive string
	var alterString string
	alterString = fmt.Sprintf("Kind: %s<br>", resp.Kind)
	//allInclusive = "Kind: " + resp.Kind + "\n\n"
	//fmt.Println("Kind:", resp.Kind)
	//fmt.Println("\n")
	for i := 0; i < len(resp.Items); i++ {
		/*fmt.Println("Item number:", i+1)
		fmt.Println("ID:", resp.Items[i].Id)
		fmt.Println("PublishedAt:", resp.Items[i].Snippet.PublishedAt)
		fmt.Println("Title:", resp.Items[i].Snippet.Title)
		fmt.Println("Url:", resp.Items[i].Snippet.Thumbnails.Default.Url)
		fmt.Println("\n")*/
		alterString += fmt.Sprintf("Item number: %v\n<br>", (i + 1))
		//fmt.Println(alterString)
		//allInclusive += "Item number: " + string(i+1) + "\n" + "ID: " + resp.Items[i].Id + "\n" + "PublishedAt: " + resp.Items[i].Snippet.PublishedAt + "\n\n" + "Title:" + resp.Items[i].Snippet.Title + "\n" + "Url:" + resp.Items[i].Snippet.Thumbnails.Default.Url + "\n\n"
	}
	fmt.Println(alterString)
	//fmt.Println("Number of videos from playlist:", len(resp.Items))
	//allInclusive += "Number of videos from playlist:" + string(len(resp.Items))
	return alterString
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

func Do_deal(defaultName string) RestResponse {
	//var

	//Inputting varuable of playlistID
	//fmt.Println("Input your playlist Id: ")
	//fmt.Scanln(&defaultName)
	//fmt.Println("\n")

	var (
	//addr = flag.String("addr", "localhost:50051", "the address to connect to")
	//name = flag.String("name", defaultName, "Name to greet")
	)
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial("0.0.0.0:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Contact the server and print out its response.
	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	var obj RestResponse
	err1 := json.Unmarshal([]byte(r.GetMessage()), &obj)
	if err1 != nil {
		fmt.Println(err1)
	}
	return obj
}
