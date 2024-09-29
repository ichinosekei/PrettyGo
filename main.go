package PrettyGo

import (
	"PrettyGo/generator"
	"PrettyGo/library"
	"PrettyGo/storage"
)

func main() {
	gen1 := generator.FirstGenerator{}
	gen2 := generator.SecondGenerator{}

	book1 := storage.Book{Title: "Book1", Author: "Author1"}
	book2 := storage.Book{Title: "Book2", Author: "Author2"}

	storageSlice := storage.StorageSlice{}
	storageSlice.Add(book1, gen1)
	storageSlice.Add(book2, gen1)

	// меняем генератор
	library.ChangeGenerator(gen2, &storageSlice)

}
