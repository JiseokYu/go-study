package main

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Writer    string    `json:"writer"`
	Publisher string    `json:"publisher"`
	CreateAt  time.Time `json:"create_at"`
}

func NewBook(name string, writer string, publisher string) *Book {
	book := new(Book)
	u, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	book.Id = u.String()
	book.Name = name
	book.Writer = writer
	book.Publisher = publisher
	book.CreateAt = time.Now()
	return book
}
