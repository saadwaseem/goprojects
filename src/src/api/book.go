package api

import (
	"encoding/json"
	"net/http"
)

type Book struct{
	Title string
	Author string
	ISBN string
}

// marshalling of book type
func (b Book) ToJSON() []byte{
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

func FromJSON(data []byte) Book{
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err!= nil {
		panic(err)
	}
	return book
}

// Books slice
var Books = []Book{Book{Title:"Book1",Author:"author1",ISBN:"isbn1"},
Book{Title:"Book2",Author:"author2",ISBN:"isbn2"}}

// Handles request for book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request){
	b, err := json.Marshal(Books)
	if err != nil{
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf8")
	w.Write(b)
}