package main

var bookList []book

type book struct {
	bookName string
	number   int
	author   string
	date     string
	sub []student
}

type student struct {
	name   string
	grade  int
	id     string
	gender string
	book   book
}

func createBook(bookName string, number int, author string, date string) {

}

func searchBook() {

}

func checkStuInfo()  {

}

fun


