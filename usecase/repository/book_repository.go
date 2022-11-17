package repository

import "github.com/inspirasiprogrammer/go-clean-arc/domain"

type BookRepository interface {
	Create(book *domain.Book) (*domain.Book, error)
	Update(book *domain.Book) (*domain.Book, error)
	Delete(book *domain.Book) error
	Fetch() ([]*domain.Book, error)
	FindByISBN(isbn string) (*domain.Book, error)
	FindByTitle(title string) ([]*domain.Book, error)
}
