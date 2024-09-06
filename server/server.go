package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "libraraymanagment/generated"
	"libraraymanagment/models"

	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type server struct {
	pb.UnimplementedLibraryServiceServer
	db *gorm.DB
}

func (s *server) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	book := models.Book{
		Title:      req.Book.Title,
		Author:     req.Book.Author,
		IsBorrowed: false,
	}
	result := s.db.Create(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pb.AddBookResponse{
		Book: &pb.Book{
			Id:         int32(book.ID),
			Title:      book.Title,
			Author:     book.Author,
			IsBorrowed: book.IsBorrowed,
		},
	}, nil
}

func (s *server) GetBooks(ctx context.Context, req *pb.GetBooksRequest) (*pb.GetBooksResponse, error) {
	var books []models.Book
	result := s.db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	var pbBooks []*pb.Book
	for _, book := range books {
		pbBooks = append(pbBooks, &pb.Book{
			Id:         int32(book.ID),
			Title:      book.Title,
			Author:     book.Author,
			IsBorrowed: book.IsBorrowed,
		})
	}

	return &pb.GetBooksResponse{Books: pbBooks}, nil
}

func (s *server) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	var book models.Book
	result := s.db.First(&book, req.BookId)
	if result.Error != nil {
		return nil, result.Error
	}

	if book.IsBorrowed {
		return nil, fmt.Errorf("book already borrowed")
	}

	book.IsBorrowed = true
	s.db.Save(&book)

	return &pb.BorrowBookResponse{
		Book: &pb.Book{
			Id:         int32(book.ID),
			Title:      book.Title,
			Author:     book.Author,
			IsBorrowed: book.IsBorrowed,
		},
	}, nil
}

func (s *server) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	var book models.Book
	result := s.db.First(&book, req.BookId)
	if result.Error != nil {
		return nil, fmt.Errorf("book not found")
	}

	s.db.Delete(&book)
	return &pb.DeleteBookResponse{
		Message: "Book deleted successfully",
	}, nil
}

func (s *server) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	var book models.Book
	result := s.db.First(&book, req.Book.Id)
	if result.Error != nil {
		return nil, fmt.Errorf("book not found")
	}

	book.Title = req.Book.Title
	book.Author = req.Book.Author
	book.IsBorrowed = req.Book.IsBorrowed
	s.db.Save(&book)

	return &pb.UpdateBookResponse{
		Book: &pb.Book{
			Id:         int32(book.ID),
			Title:      book.Title,
			Author:     book.Author,
			IsBorrowed: book.IsBorrowed,
		},
	}, nil
}

func main() {
	// Set up database connection
	db, err := gorm.Open(sqlite.Open("library.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&models.Book{})

	// Set up gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterLibraryServiceServer(grpcServer, &server{db: db})

	log.Printf("gRPC server started on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
