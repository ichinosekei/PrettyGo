package library

import (
	"PrettyGo/generator"
	"PrettyGo/storage"

	"testing"
)

func TestChangeGenerator(t *testing.T) {
	gen1 := generator.FirstGenerator{}
	gen2 := generator.SecondGenerator{}

	book1 := storage.Book{Title: "Book1", Author: "Author1"}
	book2 := storage.Book{Title: "Book2", Author: "Author2"}

	storageSlice := storage.StorageSlice{}
	storageSlice.Add(book1, gen1)
	storageSlice.Add(book2, gen1)

	// проверяем сгенерированы ли id
	firstID := storageSlice.Books[0].ID
	if firstID == "" {
		t.Errorf("Expected first generator ID, got empty string")
	}

	// меняем генератор
	ChangeGenerator(gen2, &storageSlice)

	// проверяем поменялось ли
	secondID := storageSlice.Books[0].ID
	if firstID == secondID {
		t.Errorf("Expected different ID after changing generator, but got the same")
	}
}

func TestSearch(t *testing.T) {
	gen := generator.FirstGenerator{}
	book := storage.Book{Title: "Test Book", Author: "Author"}

	// Создаём хранилище с одной книгой
	storageSlice := storage.StorageSlice{}
	storageSlice.Add(book, gen)

	// Ищем книгу по названию с использованием того же генератора
	foundBook, ok := Search("Test Book", gen, &storageSlice)
	if !ok {
		t.Error("Expected book to be found")
	}

	if foundBook.Title != "Test Book" {
		t.Errorf("Expected 'Test Book', got %s", foundBook.Title)
	}
}

func TestSearch_NotFound(t *testing.T) {
	gen := generator.FirstGenerator{}
	book := storage.Book{Title: "Test Book", Author: "Author"}

	// Создаём хранилище с одной книгой
	storageSlice := storage.StorageSlice{}
	storageSlice.Add(book, gen)

	// Ищем несуществующую книгу
	_, ok := Search("Nonexistent Book", gen, &storageSlice)
	if ok {
		t.Error("Expected book not to be found")
	}
}
