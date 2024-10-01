
# Book Library Management

Этот проект представляет собой простую библиотеку для управления книгами с использованием различных генераторов идентификаторов. Он позволяет добавлять книги в хранилище, изменять идентификаторы с помощью различных алгоритмов хеширования и выполнять поиск книг, а также изменять хранилеще.

## Структура проекта

```bash
PrettyGo/
├── generator/
│   ├── first_generator.go      # Реализация первого генератора хеширования(SHA-1)
│   ├── second_generator.go     # Реализация второго генератора хеширования(MD5)
│   ├── generator_test.go       # Тесты для генераторов
├── library/
│   ├── library.go              # Интерфейсы и реализации для управления библиотекой
│   ├── library_test.go         # Тесты для библиотечного функционала
├── storage/
│   ├── book.go                 # Определение структуры Book
│   ├── storage_map.go          # Реализация хранилища для книг с использованием map
│   ├── storage_slice.go        # Реализация хранилища для книг с использованием slice
│   ├── storage_test.go         # Тесты для хранилища
├── go.mod                      # Модульный файл Go
└── README.md                   # Описание проекта
```
# Общее описание проекта

# Использование
### Добавление книги
Вы можете добавлять книги в хранилище, используя FirstGenerator или SecondGenerator для генерации различных идентификаторов. 
Пример добавления книги с первым типом хеширования
```go 
gen := generator.FirstGenerator{}
book := storage.Book{Title: "Test Book", Author: "Author"}
storage := storage.StorageSlice{}
storage.Add(book, gen)
```
### Поиск книги
Используйте функцию Search для поиска книг по их названию.
Пример как это сделать
```go
gen := generator.FirstGenerator{}
book := storage.Book{Title: "Test Book", Author: "Author"}

// Создаём хранилище с одной книгой
storageSlice := storage.StorageSlice{}
storageSlice.Add(book, gen)

// Ищем книгу по названию с использованием того же генератора
foundBook, ok := Search("Test Book", gen, &storageSlice)
```

### Изменение генератора
Вы можете изменять генератор, это приведет к обновлению id всех книг в хранилище.
Пример как это сделать 
```go
gen1 := generator.FirstGenerator{}
gen2 := generator.SecondGenerator{}

book1 := storage.Book{Title: "Book1", Author: "Author1"}
book2 := storage.Book{Title: "Book2", Author: "Author2"}

storageSlice := storage.StorageSlice{}
storageSlice.Add(book1, gen1)
storageSlice.Add(book2, gen1)

// меняем генератор
library.ChangeGenerator(gen2, &storageSlice)
```


### Преобразование хранилищ
Вы можете изменить тип хранения книг, между StorageSlice и StorageMap.
```go
gen := generator.FirstGenerator{}
book1 := Book{Title: "Book1", Author: "Author1"}
book2 := Book{Title: "Book2", Author: "Author2"}

// Создаём хранилище типа Slice
storageSlice := StorageSlice{}
storageSlice.Add(book1, gen)
storageSlice.Add(book2, gen)

// Преобразуем Slice в Map
storageMap := storageSlice.ChangeStorage()
```

# Тестирование
### Тесты можно запустить с помощью команды:
```bash
go test ./...
```
Эта команда выполнит все тесты, определенные в файлах, заканчивающихся на _test.go.