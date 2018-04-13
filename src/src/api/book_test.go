package api
 import (
	 "testing"
	//  "fmt"
	 "github.com/stretchr/testify/assert"
 )

 func TestBookToJSON(t *testing.T){
	 book := Book{Title: "Micro Service in Go", Author: "Saad Waseem", ISBN: "0123456789"}
	 json := book.ToJSON()
	//  fmt.Println(string(json))
	 assert.Equal(t, "{\"Title\":\"Micro Service in Go\",\"Author\":\"Saad Waseem\",\"ISBN\":\"0123456789\"}",
	string(json), "Book JSON marshalling wrong.")
 }

 func TestBookFromJSON(t *testing.T){
	 json := []byte ("{\"Title\":\"Micro Service in Go\",\"Author\":\"Saad Waseem\",\"ISBN\":\"0123456789\"}")
	 book := FromJSON(json)
	 assert.Equal(t, Book{Title:"Micro Service in Go",Author:"Saad Waseem",ISBN:"0123456789"},
		book, "Book JSON unmarshalling wrong.")
}