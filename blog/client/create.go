package main

import (
	"context"
	"log"

	pb "github.com/vijay-ss/blogpost-api-grpc/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("---createBlog was invoked---")

	blog := &pb.Blog{
		AuthorId: "Iron Mike",
		Title: "Spar Wars",
		Content: "Dogs gotta eat!",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected errpr: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res)

	return res.Id
}