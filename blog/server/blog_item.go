package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	pb "github.com/vijay-ss/blogpost-api-grpc/blog/proto"
)

type BlogItem struct {
	ID 			primitive.ObjectID	`bson:"_id,omitempty"`
	AuthorID 	string 				`bson:"author_id"`
	Title 		string 				`bson:"title"`
	Content 	string 				`bson:"content"`
}

func documentToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id: 		data.ID.Hex(),
		AuthorId:	data.AuthorID,
		Title: 		data.Title,
		Content: 	data.Content,
	}
}