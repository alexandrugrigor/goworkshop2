package main

import (
	"fmt"
	"goworkshop/model"
	"goworkshop/importer"
)

func main() {
	model.Authors = importer.ImportAuthors("importer/authors.json")
	fmt.Printf("Imported authors are: %s\n", model.Authors)

	model.Books= importer.ImportBooks("importer/books.json")
	fmt.Printf("Imported books are: %s\n", model.Books)
}
