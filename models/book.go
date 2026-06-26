package models

import "time"

type Book struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	ISBN            string `json:"isbn"`
	TotalCopies     int    `json:"total_copies"`
	AvailableCopies int    `json:"available_copies"`
	CreatedAt       time.Time `json:"created_at"`
}

//what the client sends us when creating  a book
type CreatedBook struct {
	Title string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	ISBN string `json:"isbn" binding:"required"`
	TotalCopies int `json:"total_copies" binding:"required, min=1"`
}

//what the clients sends us whn updating a book
type UpdateBookInput struct {
	Title string `json:"title"`
	Author string `json:"author"`
	ISBN string `json:"isbn"`
}