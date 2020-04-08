package handler

import (
	"../books"
	"../logic"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func writeData(writer http.ResponseWriter,byteArray []byte, failed bool,errMessage string) int{
	var data []byte
	var err error
	if !failed && !books.Failed{
		data = byteArray
	}else {
		books.ERROR.Error = errMessage
		data,err = json.MarshalIndent(books.ERROR,"","  ")
		if err!=nil{
			fmt.Println(err)
			writer.Header().Add("Content-Type", "text/plain")
			data = []byte("Error occurred while processing")
			writer.WriteHeader(500)
		}
	}
	n,err := writer.Write(data)
	if err != nil{
		fmt.Println(err)
	}
	return n
}

func checkErrSetHeader(err error,writer http.ResponseWriter,code int) (http.ResponseWriter,bool) {
	var failed = false
	writer.Header().Add("Content-Type", "application/json")
	if err!=nil || books.Failed{
		if err!=nil{
			fmt.Println(err)
		}
		writer.WriteHeader(500)
		failed = true
	}else {
		writer.WriteHeader(code)
	}
	return writer,failed

}
func AllBooks(writer http.ResponseWriter,r *http.Request)  {
	fmt.Println("All Books Endpoint Hit")
	booksJson,err := json.MarshalIndent(books.BOOKS,"","  ")
	writer,failed := checkErrSetHeader(err,writer,200)
	writeData(writer,booksJson,failed,"Error occurred while processing")
}

func HandleHome(writer http.ResponseWriter,r *http.Request) {
	n,err := fmt.Fprintf(writer,"Homepage")
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%d bytes written for homepage.\n", n)
}

func HandleByTitle(writer http.ResponseWriter,r *http.Request)  {
	fmt.Println("Search By Title Endpoint Hit")
	vars:= mux.Vars(r)
	keyword:=  vars["title"]
	booksJson,err := json.MarshalIndent(logic.FindByTitleKeyword(keyword),"","  ")
	writer,failed := checkErrSetHeader(err,writer,200)
	writeData(writer,booksJson,failed,"Error occurred while processing")
}

func HandleByAuthor(writer http.ResponseWriter,r *http.Request)  {
	fmt.Println("Search By Author Endpoint Hit")
	vars:= mux.Vars(r)
	keyword:=  vars["author"]
	booksJson,err := json.MarshalIndent(logic.FindByAuthorKeyword(keyword),"","  ")
	writer,failed := checkErrSetHeader(err,writer,200)
	writeData(writer,booksJson,failed,"Error occurred while processing")
}

func checkByPage(flag int,writer http.ResponseWriter,foundBooks books.Books){
	validInput := false
	var data string
	if flag==1{
		data = "Please enter valid integer number for page range"
	}else if flag ==2{
		data = "Number of pages should be positive integer"
	}else if flag==3{
		data = "Minimum page should not exceed maximum page"
	}else if flag==0{
		validInput = true
	}
	writer.Header().Add("Content-Type", "application/json")
	if validInput{
		booksJson,err := json.MarshalIndent(foundBooks,"","  ")
		writer,failed :=checkErrSetHeader(err,writer,200)
		writeData(writer,booksJson,failed,"Error occurred while processing")
	}else{
		books.ERROR.Error = data
		message,err := json.MarshalIndent(books.ERROR,"","  ")
		writer,failed:= checkErrSetHeader(err,writer,400)
		writeData(writer,message,failed,"Error occurred while processing")
	}

}
func HandleByPageRange(writer http.ResponseWriter,r *http.Request)  {
	fmt.Println("Search By Page Range Endpoint Hit")
	vars:= mux.Vars(r)
	min:=  vars["min"]
	max:= vars["max"]
	foundBook,flag:=logic.FindByPageRange(min,max)
	checkByPage(flag,writer,foundBook)

}

func HandleByLanguage(writer http.ResponseWriter,r *http.Request){
	fmt.Println("Search By Language Endpoint Hit")
	vars:= mux.Vars(r)
	keyword :=  vars["lang"]
	booksJson,err := json.MarshalIndent(logic.FindByLanguage(keyword),"","  ")
	writer,failed := checkErrSetHeader(err,writer,200)
	writeData(writer,booksJson,failed,"Error occurred while processing")

}

func HandleByID(writer http.ResponseWriter,r *http.Request)  {
	fmt.Println("Search By ID Endpoint Hit")
	vars:= mux.Vars(r)
	id :=  vars["id"]
	writer.Header().Add("Content-Type", "application/json")
	data,failed := logic.FindByID(id)
	if failed{
		books.ERROR.Error = "Please enter valid integer for ID"
		message,err := json.MarshalIndent(books.ERROR,"","  ")
		writer,failed := checkErrSetHeader(err,writer,400)
		writeData(writer,message,failed,"Error occurred while processing")
	}else{
		message,err := json.MarshalIndent(data,"","  ")
		writer,failed := checkErrSetHeader(err,writer,200)
		writeData(writer,message,failed,"Error occurred while processing")
	}
}