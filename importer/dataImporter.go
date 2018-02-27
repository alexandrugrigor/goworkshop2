package importer

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"goworkshop/model"
)

func ImportAuthors(filename string) []model.AuthorDto {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Failed to read authors.json file: %s\n", err)
		panic(err)
	}
	var authors []model.AuthorDto
	err = json.Unmarshal(fileContent, &authors)
	if err != nil {
		fmt.Printf("Failed to unmarshal authors: %s\n", err)
		panic(err)
	}
	return authors
}

func ImportBooks(filename string) []model.BookDto {
	fileContent, err := ioutil.ReadFile(filename )
	if err != nil {
		fmt.Printf("Failed to read book.json file: %s\n", err)
		panic(err)
	}
	var books []model.BookDto
	err = json.Unmarshal(fileContent, &books)
	if err != nil {
		fmt.Printf("Failed to unmarshal books: %s\n", err)
		panic(err)
	}
	return books
}