package handler

import (
	"../books"
	"../logic"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func writeData(writer http.ResponseWriter,byteArray []byte,failed bool) int{
	var data []byte
	var err error
	if !failed{
		data = byteArray
	}else {
		books.ERROR.Error = "Error occurred"
		data,err = json.MarshalIndent(books.ERROR,"","  ")
		if err!=nil{
			fmt.Println(err)
		}
	}
	n,err := writer.Write(data)
	if err != nil{
		panic(err)
	}
	return n
}

func checkErrSetHeader(err error,writer http.ResponseWriter) (http.ResponseWriter,bool) {
	var failed = false
	writer.Header().Add("Content-Type", "application/json")
	if err!=nil || books.Failed{
		if err!=nil{
			fmt.Println(err)
		}
		writer.WriteHeader(http.StatusInternalServerError)
		failed = true
	}else {
		writer.WriteHeader(http.StatusOK)
	}
	return writer,failed

}
func AllBooks(writer http.ResponseWriter,r *http.Request)  {
	fmt.Println("All Books Endpoint Hit")
	booksJson,err := json.MarshalIndent(books.BOOKS,"","  ")
	writer,failed := checkErrSetHeader(err,writer)
	writeData(writer,booksJson,failed)
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
	writer,failed := checkErrSetHeader(err,writer)
	writeData(writer,booksJson,failed)
}

func HandleByAuthor(writer http.ResponseWriter,r *http.Request)  {
	fmt.Println("Search By Author Endpoint Hit")
	vars:= mux.Vars(r)
	keyword:=  vars["author"]
	booksJson,err := json.MarshalIndent(logic.FindByAuthorKeyword(keyword),"","  ")
	writer,failed := checkErrSetHeader(err,writer)
	writeData(writer,booksJson,failed)
}

func checkByPage(err error,flag int,writer http.ResponseWriter,booksJson []byte){
	validInput := false
	var data string
	if flag==1{
		data = "Please enter integer number for page range"
	}else if flag ==2{
		data = "Number of pages should be positive integer"
	}else if flag==3{
		data = "Minimum page should not exceed maximum page"
	}else if flag==0{
		validInput = true
	}
	writer.Header().Add("Content-Type", "application/json")
	if validInput{
		if err==nil && !books.Failed{
			writer.WriteHeader(http.StatusOK)
			writeData(writer,booksJson,false)
		}else{
			writer.WriteHeader(http.StatusInternalServerError)
			writeData(writer,[]byte("Error occurred"),true)
		}
	}else{
		writer.WriteHeader(http.StatusBadRequest)
		books.ERROR.Error = data
		message,err := json.MarshalIndent(books.ERROR,"","  ")
		if err!=nil{
			fmt.Println(err)
		}
		writeData(writer,message,false)
	}

}
func HandleByPageRange(writer http.ResponseWriter,r *http.Request)  {
	fmt.Println("Search By Page Range Endpoint Hit")
	vars:= mux.Vars(r)
	min:=  vars["min"]
	max:= vars["max"]
	data,flag:=logic.FindByPageRange(min,max)
	booksJson,err := json.MarshalIndent(data,"","  ")
	checkByPage(err,flag,writer,booksJson)

}

func HandleByLanguage(writer http.ResponseWriter,r *http.Request){
	fmt.Println("Search By Language Endpoint Hit")
	vars:= mux.Vars(r)
	keyword :=  vars["lang"]
	booksJson,err := json.MarshalIndent(logic.FindByLanguage(keyword),"","  ")
	writer,failed := checkErrSetHeader(err,writer)
	writeData(writer,booksJson,failed)

}