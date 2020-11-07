package handler

import (
	"../data"
	"../logic"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"../utils"
)

func AllBooks(writer http.ResponseWriter,r *http.Request)  {

	fmt.Println("All Books Endpoint Hit")
	utils.SendBooks(writer, 200, data.BOOKS)

}

func HandleHome(writer http.ResponseWriter,r *http.Request) {

	_, err := fmt.Fprintf(writer,"Homepage")

	if err!=nil{
		fmt.Println(err)

	}
}

func HandleByTitle(writer http.ResponseWriter,r *http.Request)  {

	fmt.Println("Search By Title Endpoint Hit")

	vars := mux.Vars(r)
	keyword :=  vars["title"]

	utils.SendBooks(writer, 200, logic.FindByTitleKeyword(keyword))

}

func HandleByAuthor(writer http.ResponseWriter,r *http.Request)  {

	fmt.Println("Search By Author Endpoint Hit")

	vars := mux.Vars(r)
	keyword := vars["author"]

	utils.SendBooks(writer, 200, logic.FindByAuthorKeyword(keyword))

}

func HandleByPageRange(writer http.ResponseWriter,r *http.Request)  {

	fmt.Println("Search By Page Range Endpoint Hit")

	vars := mux.Vars(r)
	min :=  vars["min"]
	max := vars["max"]

	foundBooks,err := logic.FindByPageRange(min,max)
	if err != "" {
		utils.SendError(writer, 400, err)

	}else {
		utils.SendBooks(writer, 200, foundBooks)

	}

}

func HandleByLanguage(writer http.ResponseWriter,r *http.Request){

	fmt.Println("Search By Language Endpoint Hit")

	vars:= mux.Vars(r)
	keyword :=  vars["lang"]

	utils.SendBooks(writer, 200, logic.FindByLanguage(keyword))

}

func HandleByID(writer http.ResponseWriter,r *http.Request)  {

	fmt.Println("Search By ID Endpoint Hit")

	vars:= mux.Vars(r)
	id :=  vars["id"]

	
	book,err := logic.FindByID(id)

	if err != ""{
		utils.SendBook(writer, 400, book, err)

	}else{
		utils.SendBook(writer, 200, book, err)
		
	}
	
}