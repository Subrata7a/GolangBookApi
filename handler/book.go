package handler

import (
	"GolangBookApi/model"
	"GolangBookApi/utils"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/oklog/ulid/v2"
	"net/http"
	"reflect"
)

var BookDB = make(map[string]model.Book)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook model.Book

	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		(&model.Error{}).GetError(w, http.StatusBadRequest, utils.StatusBadRequest, "Resource not found")
		return
	}
	newBook.UUID = ulid.Make().String()
	curBook := reflect.ValueOf(newBook)

	for i := 0; i < curBook.NumField(); i++ {
		fieldValue := curBook.Field(i)

		if reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface()) {
			(&model.Error{}).GetError(w, http.StatusBadRequest, utils.StatusBadRequest, "Invalid Request")
			return
		}
	}

	BookDB[newBook.UUID] = newBook
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(newBook); err != nil {
		(&model.Error{}).GetError(w, http.StatusBadRequest, utils.StatusBadRequest, "Encoding Error")
		return
	}
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	bookID, err := ulid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		(&model.Error{}).GetError(w, http.StatusBadRequest, utils.StatusBadRequest, "Parsing Error")
		return
	}

	currentBook, exists := BookDB[bookID.String()]
	if !exists {
		(&model.Error{}).GetError(w, http.StatusNotFound, utils.StatusNotFound, "Invalid handler ID")
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(currentBook); err != nil {
		(&model.Error{}).GetError(w, http.StatusBadRequest, utils.StatusBadRequest, "Invalid Encoding")
		return
	}
}

func ListOfBooks(w http.ResponseWriter, r *http.Request) {
	bookList := make([]model.Book, 0, len(BookDB))
	for _, book := range BookDB {
		bookList = append(bookList, book)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(bookList); err != nil {
		(&model.Error{}).GetError(w, http.StatusBadRequest, utils.StatusBadRequest, "Invalid Encoding")
		return
	}
	return
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookID, err := ulid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		(&model.Error{}).GetError(w, http.StatusBadRequest, utils.StatusBadRequest, "Parsing Error")
		return
	}

	var updatedBook model.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		(&model.Error{}).GetError(w, http.StatusBadRequest, utils.StatusBadRequest, "Invalid request body")
		return
	}

	_, exists := BookDB[bookID.String()]
	if !exists {
		(&model.Error{}).GetError(w, http.StatusNotFound, utils.StatusNotFound, "Book not found")
		return
	}

	updatedBook.UUID = bookID.String()

	curBook := reflect.ValueOf(updatedBook)

	for i := 0; i < curBook.NumField(); i++ {
		fieldValue := curBook.Field(i)

		if reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface()) {
			(&model.Error{}).GetError(w, http.StatusBadRequest, utils.StatusBadRequest, "Invalid Request")
			return
		}
	}
	BookDB[bookID.String()] = updatedBook

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(updatedBook); err != nil {
		(&model.Error{}).GetError(w, http.StatusBadRequest, utils.StatusBadRequest, "Invalid Encoding")
		return
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID, err := ulid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		(&model.Error{}).GetError(w, http.StatusBadRequest, utils.StatusBadRequest, "Parsing Error")
		return
	}

	var deletedBook model.Book
	if err := json.NewDecoder(r.Body).Decode(&deletedBook); err != nil {
		(&model.Error{}).GetError(w, http.StatusBadRequest, utils.StatusBadRequest, "Invalid Request Body")
		return
	}

	_, exists := BookDB[bookID.String()]
	if !exists {
		(&model.Error{}).GetError(w, http.StatusNotFound, utils.StatusNotFound, "Book Not Found")
		return
	}

	deletedBook.UUID = bookID.String()
	BookDB[bookID.String()] = deletedBook
	delete(BookDB, bookID.String())

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(deletedBook); err != nil {
		(&model.Error{}).GetError(w, http.StatusBadRequest, utils.StatusBadRequest, "Invalid Encoding")
		return
	}
}
