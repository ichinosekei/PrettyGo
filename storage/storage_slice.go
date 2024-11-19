package storage

import "PrettyGo/generator"

type StorageSlice struct {
	Books []Book
}

func (s *StorageSlice) Add(book Book, gen generator.Generator) {
	book.ID = gen.GenerateId(book.Title)
	s.Books = append(s.Books, book)
}

func (s *StorageSlice) BooksList() []Book {
	return s.Books
}

// ChangeStorage преобразуем storageSlice в storageMap.
func (v *StorageSlice) ChangeStorage() StorageMap {
	books := make(map[int]Book)
	for i, book := range v.Books {
		books[i] = book
	}
	return StorageMap{Books: books}
}
