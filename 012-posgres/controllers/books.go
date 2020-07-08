package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/harrysaini/try-go-web-dev/012-posgres/models"
	"github.com/harrysaini/try-go-web-dev/012-posgres/utils"
	"github.com/julienschmidt/httprouter"
)

// BooksController books routes handler
type BooksController struct {
	dblink models.BooksDataStore
}

// NewBooksController get instance of books routes handler
func NewBooksController(dblink models.BooksDataStore) *BooksController {
	return &BooksController{
		dblink: dblink,
	}
}

// GetAllBooks fetch list of all books
func (bc *BooksController) GetAllBooks(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	books, err := bc.dblink.GetAllBooks()
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, books)
}

// GetBook fetch list of all books
func (bc *BooksController) GetBook(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	isbn := params.ByName("isbn")

	if isbn == "" {
		utils.SendErrorResponse(w, errors.New("isbn missing"), http.StatusInternalServerError)
		return
	}

	book, err := bc.dblink.GetBook(isbn)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, book)
}

// InsertBook add new book
func (bc *BooksController) InsertBook(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var book models.Book

	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	_, err = book.Validate()
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	_, err = bc.dblink.InsertBook(&book)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, nil)
}

// UpdateBook add new book
func (bc *BooksController) UpdateBook(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	isbn := params.ByName("isbn")

	if isbn == "" {
		utils.SendErrorResponse(w, errors.New("isbn missing"), http.StatusInternalServerError)
		return
	}

	_, err := bc.dblink.GetBook(isbn)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	updateDataRaw := make(map[string]interface{})

	err = json.NewDecoder(req.Body).Decode(&updateDataRaw)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	updateData := make(map[string]interface{})

	for _, field := range models.BooksFieldsUdateAllowed {
		if v, ok := updateDataRaw[field]; ok {
			updateData[field] = v
		}
	}

	_, err = bc.dblink.UpdateBook(isbn, updateData)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, nil)
}

// DeleteBook add new book
func (bc *BooksController) DeleteBook(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	isbn := params.ByName("isbn")

	if isbn == "" {
		utils.SendErrorResponse(w, errors.New("isbn missing"), http.StatusInternalServerError)
		return
	}

	_, err := bc.dblink.GetBook(isbn)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	_, err = bc.dblink.DeleteBook(isbn)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, nil)
}
