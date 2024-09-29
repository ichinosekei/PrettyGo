package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

type FirstGenerator struct {
	active bool
	whichGenerator
}
type SecondGenerator struct {
	active bool
	whichGenerator
}

type whichGenerator interface {
	GenerateId(title string) string
	ChangeGenerator(l storageSlice)
}

type Book struct {
	id     string
	Title  string
	Author string
}

type storageMap struct {
	books map[int]Book
	Library
}

type storageSlice struct {
	books []Book
	Library
}

type Library interface {
	bookS() []Book
	ChangeGenerator(generator whichGenerator, archive Library)
	//ChangeStorage()
	//Add(book Book, generator which_generator)
}

func (l *storageSlice) bookS() []Book {
	return l.books
}

func (mapp *storageMap) bookS() []Book {

	array := make([]Book, 0, len(mapp.books))
	for i, _ := range mapp.books {
		array = append(array, mapp.books[i])
	}
	return array
}

func (FirstGenerator) GenerateId(title string) string {
	hash := sha1.New()
	hash.Write([]byte(title))
	return fmt.Sprintf("%x", hash.Sum(nil))
	//fmt.Println(book.id)
}

func (SecondGenerator) GenerateId(title string) string {
	hash := md5.New()
	hash.Write([]byte(title))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (l *storageSlice) Add(book Book, generator whichGenerator) {
	book.id = generator.GenerateId(book.Title)
	l.books = append(l.books, book)
}

func (mapp *storageMap) Add(book Book, generator whichGenerator) {
	if mapp.books == nil {
		mapp.books = make(map[int]Book)
	}
	book.id = generator.GenerateId(book.Title)
	mapp.books[len(mapp.books)] = book

}

func ChangeGenerator(generator whichGenerator, archive Library) {
	for i, book := range archive.bookS() {
		archive.bookS()[i].id = generator.GenerateId(book.Title)
	}
}

func (v *storageSlice) ChangeStorage() storageMap {
	books := map[int]Book{}
	for i, book := range v.books {
		books[i] = book
	}
	return storageMap{books: books}

}

func (v *storageMap) ChangeStorage() storageSlice {
	books := make([]Book, 0, len(v.books))
	for _, book := range v.books {
		books = append(books, book)
	}
	return storageSlice{books: books}
}

//func ChangeStorage(ArchieveOld *Library) Library {
//    unpoint := *ArchieveOld
//	switch v := unpoint.(type) {
//	case *storageSlice:
//		//storageMap{books: map[int]Book{},nil}
//		ArchieveOld = (v.ChangetoMap())
//		ArchieveOld = v.ChangetoMap()
//		println(v.books)
//		return ArchieveOld
//
//	case *storageMap:
//		v.ChangeToMap()
//		return v.ChangetoSlice()
//	}
//
//	//println("not change")
//	return ArchieveOld
//}

func find(id string, archive Library) (Book, bool) {
	for _, book := range archive.bookS() {
		if book.id == id {
			return book, true
		}
	}
	return Book{}, false
}

func Search(name string, generator whichGenerator, archive Library) (Book, bool) {
	var book Book
	var ok bool
	book, ok = find(generator.GenerateId(name), archive)
	if ok {
		return book, true
	}
	return Book{}, false

}

func main() {
	hash := sha1.New()
	hash.Write([]byte("Book1"))
	s1 := fmt.Sprintf("%x", hash.Sum(nil))

	hash = sha1.New()
	hash.Write([]byte("Book2"))
	s2 := fmt.Sprintf("%x", hash.Sum(nil))

	hash = sha1.New()
	hash.Write([]byte("Book3"))
	s3 := fmt.Sprintf("%x", hash.Sum(nil))

	hash = sha1.New()
	hash.Write([]byte("Book4"))
	s4 := fmt.Sprintf("%x", hash.Sum(nil))

	hash = sha1.New()
	hash.Write([]byte("Book5"))
	s5 := fmt.Sprintf("%x", hash.Sum(nil))

	//fmt.Sprintf("%x", sha1.New().Write([]byte("Book1")).Sum(nil))
	Book1 := Book{s1, "Book1", "Author1"}
	Book2 := Book{s2, "Book2", "Author2"}
	Book3 := Book{s3, "Book3", "Author3"}
	Book4 := Book{s4, "Book4", "Author4"}
	//Book5 := Book{fmt.Sprintf("%x", sha1.New().Write([]byte("Book5")).Sum(nil)), "Book5", "Author5"}
	storage1 := storageSlice{[]Book{Book1, Book2, Book3, Book4}, nil}
	book, ok := (Search("Book5", FirstGenerator{}, &storage1))

	storage1.Add(Book{s5, "Book5", "Author5"}, FirstGenerator{})

	println(storage1.books[4].id)
	ChangeGenerator(SecondGenerator{}, &storage1)
	println(storage1.books[4].id)

	book, ok = (Search("Book3", SecondGenerator{}, &storage1))
	if ok {
		println(book.id)
	} else {
		println("not found")
	}
	///////// map
	println("map")
	mapp := map[int]Book{
		0: Book{s1, "Book1", "Author1"},
		1: Book{s2, "Book2", "Author2"},
		2: Book{s3, "Book3", "Author3"},
		3: Book{s4, "Book4", "Author4"},
		4: Book{s5, "Book5", "Author5"},
	}
	storage2 := storageMap{mapp, nil}
	storage2.Add(Book{s5, "Book5", "Author5"}, FirstGenerator{})
	println(storage2.books[5].id)
	ChangeGenerator(SecondGenerator{}, &storage2)

	////// change storageSlice
	println("change storageSlice")
	storage3 := storageSlice{[]Book{Book1, Book2, Book3, Book4}, nil}
	//r := storageMap{map[int]Book{}, nil}
	def := storage3.ChangeStorage()
	book, ok = (Search("Book3", FirstGenerator{}, &def))
	if ok {
		println(book.id)
	} else {
		println("not found")
	}

}
