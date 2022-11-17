package controller

import (
	"net/http"

	"github.com/inspirasiprogrammer/go-clean-arc/domain"
	"github.com/inspirasiprogrammer/go-clean-arc/usecase/interactor"
)

type bookController struct {
	bookInteractor interactor.BookInteractor
}

type BookController interface {
	CreateBook(book *domain.Book, c Context) error
	UpdateBook(book *domain.Book, c Context) error
	DeleteBook(book *domain.Book, c Context) error
	GetBooks(c Context) error
	GetBookByISBN(isbn string, c Context) error
	GetBooksByTitle(title string, c Context) error
}

func NewBookController(bi interactor.BookInteractor) BookController {
	return &bookController{bi}
}

func (bc *bookController) CreateBook(book *domain.Book, c Context) error {
	bookDetail, err := bc.bookInteractor.Create(book)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, bookDetail)
}

func (bc *bookController) GetBooks(c Context) error {
	books, err := bc.bookInteractor.Get()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, books)
}

func (bc *bookController) DeleteBook(book *domain.Book, c Context) error {
	if err := bc.bookInteractor.Delete(book); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, book)
}

func (bc *bookController) GetBookByISBN(isbn string, c Context) error {
	book, err := bc.bookInteractor.GetByISBN(isbn)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, book)
}

func (bc *bookController) GetBooksByTitle(title string, c Context) error {
	book, err := bc.bookInteractor.GetByTitle(title)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, book)
}

func (bc *bookController) UpdateBook(book *domain.Book, c Context) error {
	bookResponse, err := bc.bookInteractor.Update(book)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, bookResponse)
}
