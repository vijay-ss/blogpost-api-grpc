package main

import (
	"context"
	"log"

	pb "github.com/vijay-ss/blogpost-api-grpc/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")

	newBlog := &pb.Blog{
		Id: id,
		AuthorId: "A New Author",
		Title: "A New Title",
		Content: "Content of the first blog, with some additions.",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated.")
}