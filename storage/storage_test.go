package storage

import (
	"PrettyGo/generator"
	"testing"
)

func TestStorageSlice_Add(t *testing.T) {
	gen := generator.FirstGenerator{}
	book := Book{Title: "Test Book", Author: "Author"}

	storage := StorageSlice{}
	storage.Add(book, gen)

	if len(storage.Books) != 1 {
		t.Errorf("Expected 1 book, got %d", len(storage.Books))
	}

	if storage.Books[0].ID == "" {
		t.Error("Expected book ID to be set")
	}
}

func TestStorageMap_Add(t *testing.T) {
	gen := generator.FirstGenerator{}
	book := Book{Title: "Test Book", Author: "Author"}

	storage := StorageMap{}
	storage.Add(book, gen)

	if len(storage.Books) != 1 {
		t.Errorf("Expected 1 book, got %d", len(storage.Books))
	}

	if storage.Books[0].ID == "" {
		t.Error("Expected book ID to be set")
	}
}

func TestChangeStorageSliceToMap(t *testing.T) {
	gen := generator.FirstGenerator{}
	book1 := Book{Title: "Book1", Author: "Author1"}
	book2 := Book{Title: "Book2", Author: "Author2"}

	// Создаём хранилище типа Slice
	storageSlice := StorageSlice{}
	storageSlice.Add(book1, gen)
	storageSlice.Add(book2, gen)

	// Преобразуем Slice в Map
	storageMap := storageSlice.ChangeStorage()

	// Проверяем, что книги корректно перенесены
	if len(storageMap.Books) != 2 {
		t.Errorf("Expected 2 books, got %d", len(storageMap.Books))
	}

	for i, book := range storageSlice.Books {
		if book.ID != storageMap.Books[i].ID {
			t.Errorf("Expected ID %s, got %s", book.ID, storageMap.Books[i].ID)
		}
	}
}

func TestChangeStorageMapToSlice(t *testing.T) {
	gen := generator.FirstGenerator{}
	book1 := Book{Title: "Book1", Author: "Author1"}
	book2 := Book{Title: "Book2", Author: "Author2"}

	// Создаём хранилище типа Map
	storageMap := StorageMap{}
	storageMap.Add(book1, gen)
	storageMap.Add(book2, gen)

	// Преобразуем Map в Slice
	storageSlice := storageMap.ChangeStorage()

	// Проверим книги корректно перенесены ли
	if len(storageSlice.Books) != 2 {
		t.Errorf("Expected 2 books, got %d", len(storageSlice.Books))
	}
	// Проверяем все ли так перенеслось
	for i, book := range storageSlice.Books {
		if book.ID != storageMap.Books[i].ID {
			t.Errorf("Expected ID %s, got %s", book.ID, storageMap.Books[i].ID)
		}
	}
}
