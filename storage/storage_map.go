package storage

import "PrettyGo/generator"

type StorageMap struct {
	Books map[int]Book
}

func (m *StorageMap) Add(book Book, gen generator.Generator) {
	if m.Books == nil {
		m.Books = make(map[int]Book)
	}
	book.ID = gen.GenerateId(book.Title)
	m.Books[len(m.Books)] = book
}

func (m *StorageMap) BooksList() []Book {
	books := make([]Book, 0, len(m.Books))
	for _, book := range m.Books {
		books = append(books, book)
	}
	return books
}

// ChangeStorage преобразуем storageMap в storageSlice.
func (m *StorageMap) ChangeStorage() StorageSlice {
	books := make([]Book, 0, len(m.Books))
	for _, book := range m.Books {
		books = append(books, book)
	}
	return StorageSlice{Books: books}
}
