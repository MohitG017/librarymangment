package main

import (
	"context"
	"log"
	"time"

	pb "libraraymanagment/generated"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewLibraryServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Add a new book
	book := &pb.Book{Title: "Go Programming", Author: "John Doe"}
	res, err := c.AddBook(ctx, &pb.AddBookRequest{Book: book})
	if err != nil {
		log.Fatalf("could not add book: %v", err)
	}
	log.Printf("Added Book: %v", res.Book)

	// Get books
	booksRes, err := c.GetBooks(ctx, &pb.GetBooksRequest{})
	if err != nil {
		log.Fatalf("could not get books: %v", err)
	}
	log.Printf("Books: %v", booksRes.Books)

	// Update a book
	updateBook := &pb.Book{
		Id:         res.Book.Id,
		Title:      "Advanced Go Programming",
		Author:     "Jane Doe",
		IsBorrowed: false,
	}
	updRes, err := c.UpdateBook(ctx, &pb.UpdateBookRequest{Book: updateBook})
	if err != nil {
		log.Fatalf("could not update book: %v", err)
	}
	log.Printf("Updated Book: %v", updRes.Book)

	// Delete a book
	delRes, err := c.DeleteBook(ctx, &pb.DeleteBookRequest{BookId: res.Book.Id})
	if err != nil {
		log.Fatalf("could not delete book: %v", err)
	}
	log.Printf("Delete response: %v", delRes.Message)
}
