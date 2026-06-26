package repository

import (
	"database/sql"
	"library-api/config"
	"library-api/models"
)

//create
func CreateBook(input models.CreatedBookInput) (models.Book, error) {
	var book models.Book

	query := `
	INSERT INTO book (title, author, isbn, total_copies, available_copies)
	VALUES ($1, $2, $3, $4, $4)
	RETURNING id, title, author, isbn, totla_copies, available_copies, created_at
	`
	// $1, $2, $3, $4 are placeholders — Postgres syntax for parameterized queries.
	// NEVER use string concatenation/fmt.Sprintf for SQL — that's how SQL injection happens.
	// available_copies = total_copies initially (no books borrowed yet), hence $4 used twice.

	row := config.DB.QueryRow(query, input.Title, input.Author, input.ISBN, input.TotalCopies)

	// QueryRow expects exactly ONE row back, so we use .Scan() to map columns -> struct fields
	err := row.Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.ISBN,
		&book.TotalCopies,
		&book.AvailableCopies,
		&book.CreatedAt,
	)
	return book, err
}

//read and get all the books

func GetAllBooks() ([]models.Book, error) {
	query := `SELECT id, title, author, isbn, total_copies, available_aopies, create_at FROM books ORDER BY id`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() //always close rows or you leak DB Connections

	var books []models.Book

	//query return MULTIPLE rows, so we loop with rows.NEXT()
	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.ISBN,
			&book.TotalCopies,
			&book.AvailableCopies,
			&book.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

//read get one book by id
func GetBookByID(id int)(models.Book, error) {
	var book models.Book

	query := `SELECT id, title, author, isbn, total_copies, available_copies, created_at FROM books WHERE id = $1`

	row := config.DB.QueryRow(query, id)
	err := row.Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.ISBN,
		&book.TotalCopies,
		&book.AvailableCopies,
		&book.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return book, err // we'll check for this specific error in the service layer
	}

	return book, err
}

//update
func UpdateBook(id int, input models.UpdateBookInput) (models.Book, error) {
	var book models.Book

	query := `
	UPDATE books
	SET title = $1, author = $2, isbn = $3
	WHERE id = $4
	RETURNING id, title, author, isbn, total_copies, available_copies, created_at
	`

	row := config.DB.QueryRow(query, input.Title, input.Author, input.ISBN, id)
	err := row.Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.ISBN,
		&book.TotalCopies,
		&book.AvailableCopies,
		&book.CreatedAt,
	)

	return book, err
}

//delete
func DeleteBook(id int) error {
	query := `DELETE FROM books WHERE id = $1`

	result, err := config.DB.Exec(query, id)
	if err != nil {
		return err
	}

	//check if anything was actually deleted
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}