package main

var bookList []book

var stuList []student

type book struct {
	bookName string
	number   int
	author   string
	date     string
	sub      []student
}

type student struct {
	name   string
	grade  int
	id     string
	gender string
	dep    []book
}

func createBook(bookName string, number int, author string, date string) book {
	i := book{bookName: bookName, number: number, author: author, date: date}
	bookList = append(bookList, i)
	return i
}

func searchBook(searchKind string, searchValue string) (bool, book) {
	switch searchKind {
	case "bookName":
		for _, v := range bookList {
			if v.bookName == searchValue {
				return true, v
			}
		}
	case "author":
		for _, v := range bookList {
			if v.author == searchValue {
				return true, v
			}
		}
	}
	return false, book{}
}

func createStu(name string, grade int, id string, gender string) student{
	i:=student{name:name,grade:grade,id:id,gender:gender}
	stuList = append(stuList,i)
	return i
}

func checkStuInfo() {

}

func borrowBook() {

}

func checkBook() {

}
