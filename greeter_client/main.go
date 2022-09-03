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

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial("server:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	//deserialization from bytes to struct
	err1 := json.Unmarshal([]byte(r.GetMessage()), &obj)
	if err1 != nil {
		fmt.Println(err1)
	}
	return obj
}
