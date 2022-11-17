package interactor

import (
	"github.com/inspirasiprogrammer/go-clean-arc/domain"
	"github.com/inspirasiprogrammer/go-clean-arc/usecase/presenter"
	"github.com/inspirasiprogrammer/go-clean-arc/usecase/repository"
)

type bookInteractor struct {
	BookRepository repository.BookRepository
	BookPresenter  presenter.BookPresenter
}

type BookInteractor interface {
	Create(book *domain.Book) (*domain.ResponseDetailBook, error)
	Update(book *domain.Book) (*domain.ResponseDetailBook, error)
	Delete(book *domain.Book) error
	Get() ([]*domain.ResponseBook, error)
	GetByISBN(isbn string) (*domain.ResponseDetailBook, error)
	GetByTitle(title string) ([]*domain.ResponseBook, error)
}

func NewBookInteractor(r repository.BookRepository, p presenter.BookPresenter) BookInteractor {
	return &bookInteractor{r, p}
}

func (bi *bookInteractor) Create(book *domain.Book) (*domain.ResponseDetailBook, error) {
	book, err := bi.BookRepository.Create(book)
	if err != nil {
		return nil, err
	}

	return bi.BookPresenter.ResponseBook(book), nil
}

func (bi *bookInteractor) Update(book *domain.Book) (*domain.ResponseDetailBook, error) {
	book, err := bi.BookRepository.Update(book)
	if err != nil {
		return nil, err
	}

	return bi.BookPresenter.ResponseBook(book), nil
}

func (bi *bookInteractor) Delete(book *domain.Book) error {
	return bi.BookRepository.Delete(book)
}

func (bi *bookInteractor) Get() ([]*domain.ResponseBook, error) {
	books, err := bi.BookRepository.Fetch()
	if err != nil {
		return nil, err
	}

	return bi.BookPresenter.ResponseBooks(books), nil
}

func (bi *bookInteractor) GetByISBN(isbn string) (*domain.ResponseDetailBook, error) {
	book, err := bi.BookRepository.FindByISBN(isbn)
	if err != nil {
		return nil, err
	}

	return bi.BookPresenter.ResponseBook(book), nil
}

func (bi *bookInteractor) GetByTitle(title string) ([]*domain.ResponseBook, error) {
	books, err := bi.BookRepository.FindByTitle(title)
	if err != nil {
		return nil, err
	}

	return bi.BookPresenter.ResponseBooks(books), nil
}
