package handler

import (
	"../data"
	"../logic"
	"../template"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)
var response template.Response

func clearResponse()  {
	response.Books = nil
	response.Error = ""
}
func writeData(writer http.ResponseWriter,byteArray []byte, failed bool,errMessage string) int{
	var toSend []byte
	var err error
	if !failed && !data.Failed{
		toSend = byteArray
	}else{
		response.Error  = errMessage
		toSend,err = json.MarshalIndent(response,"","  ")
		if err!=nil{
			fmt.Println(err)
			writer.Header().Add("Content-Type", "text/plain")
			toSend = []byte("Error occurred while processing")
			writer.WriteHeader(500)
		}
	}
	n,err := writer.Write(toSend)
	if err != nil{
		fmt.Println(err)
	}
	return n
}

func checkErrSetHeader(err error,writer http.ResponseWriter,code int) (http.ResponseWriter,bool) {
	var failed = false
	writer.Header().Add("Content-Type", "application/json")
	if err!=nil || data.Failed{
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
	clearResponse()
	response.Books = data.BOOKS.Books
	booksJson,err := json.MarshalIndent(response,"","  ")
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
	clearResponse()
	vars:= mux.Vars(r)
	keyword:=  vars["title"]
	response.Books = logic.FindByTitleKeyword(keyword).Books
	booksJson,err := json.MarshalIndent(response,"","  ")
	writer,failed := checkErrSetHeader(err,writer,200)
	writeData(writer,booksJson,failed,"Error occurred while processing")
}

func HandleByAuthor(writer http.ResponseWriter,r *http.Request)  {
	fmt.Println("Search By Author Endpoint Hit")
	clearResponse()
	vars:= mux.Vars(r)
	keyword:=  vars["author"]
	response.Books = logic.FindByAuthorKeyword(keyword).Books
	booksJson,err := json.MarshalIndent(response,"","  ")
	writer,failed := checkErrSetHeader(err,writer,200)
	writeData(writer,booksJson,failed,"Error occurred while processing")
}

func checkByPage(flag int,writer http.ResponseWriter,foundBooks template.Books){
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
		response.Books = foundBooks.Books
		booksJson,err := json.MarshalIndent(response,"","  ")
		writer,failed :=checkErrSetHeader(err,writer,200)
		writeData(writer,booksJson,failed,"Error occurred while processing")
	}else{
		response.Error = data
		message,err := json.MarshalIndent(response,"","  ") ///check
		writer,failed:= checkErrSetHeader(err,writer,400)
		writeData(writer,message,failed,"Error occurred while processing")
	}

}
func HandleByPageRange(writer http.ResponseWriter,r *http.Request)  {
	fmt.Println("Search By Page Range Endpoint Hit")
	clearResponse()
	vars:= mux.Vars(r)
	min:=  vars["min"]
	max:= vars["max"]
	foundBook,flag:=logic.FindByPageRange(min,max)
	checkByPage(flag,writer,foundBook)

}

func HandleByLanguage(writer http.ResponseWriter,r *http.Request){
	fmt.Println("Search By Language Endpoint Hit")
	clearResponse()
	vars:= mux.Vars(r)
	keyword :=  vars["lang"]
	response.Books = logic.FindByLanguage(keyword).Books
	booksJson,err := json.MarshalIndent(response,"","  ")
	writer,failed := checkErrSetHeader(err,writer,200)
	writeData(writer,booksJson,failed,"Error occurred while processing")

}

func HandleByID(writer http.ResponseWriter,r *http.Request)  {
	fmt.Println("Search By ID Endpoint Hit")
	clearResponse()
	vars:= mux.Vars(r)
	id :=  vars["id"]
	writer.Header().Add("Content-Type", "application/json")
	data,failed := logic.FindByID(id)
	if failed{
		response.Error = "Please enter valid integer for ID"
		message,err := json.MarshalIndent(response,"","  ")  //check
		writer,failed := checkErrSetHeader(err,writer,400)
		writeData(writer,message,failed,"Error occurred while processing")
	}else{
		response.Books = data.Books
		message,err := json.MarshalIndent(response,"","  ")
		writer,failed := checkErrSetHeader(err,writer,200)
		writeData(writer,message,failed,"Error occurred while processing")
	}
}