package services

import (
	"database/sql"
	"errors"
	"library-api/models"
	"library-api/repository"

	"github.com/go-playground/locales/id"
)

func CreateBookService(input models.CreatedBookInput) (models.Book, error) {
//business rule examle: could check for duplicate ISBN here before inserting
return repository.CreateBook(input)
}

func GetAllBookService() ([]models.Book, error) {
	return repository.GetAllBooks()
}

func GetBookByIDService(id int) (models.Book, error) {
	book, err := repository.GetBookByID(id) 
	if err == sql.ErrNoRows {
		return book, errors.New("book not found")
	}
	return book, err
}

func DeleteBookService(id int) error {
	err := repository.DeleteBook(id)
	if err == sql.ErrNoRows {
		return errors.New("book not found")
	}
	return err
}