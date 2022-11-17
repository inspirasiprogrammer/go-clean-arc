package presenter

import "github.com/inspirasiprogrammer/go-clean-arc/domain"

type BookPresenter interface {
	ResponseBooks(books []*domain.Book) []*domain.ResponseBook
	ResponseBook(book *domain.Book) *domain.ResponseDetailBook
}