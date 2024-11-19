package library

import (
	"PrettyGo/generator"
	"PrettyGo/storage"
)

type Library interface {
	BooksList() []storage.Book
	Add(book storage.Book, gen generator.Generator)
}

func ChangeGenerator(gen generator.Generator, lib Library) {
	for i, book := range lib.BooksList() {
		lib.BooksList()[i].ID = gen.GenerateId(book.Title)
	}
}

func Search(name string, gen generator.Generator, archive Library) (storage.Book, bool) {
	bookID := gen.GenerateId(name)
	for _, book := range archive.BooksList() {
		if book.ID == bookID {
			return book, true
		}
	}
	return storage.Book{}, false
}
