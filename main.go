package main

import (
	"./handler"
	"./data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func handleRequest()  {

	myRouter := mux.NewRouter().StrictSlash(true)
	
	myRouter.HandleFunc("/",handler.HandleHome).Methods("GET")
	myRouter.HandleFunc("/all",handler.AllBooks).Methods("GET")
	myRouter.HandleFunc("/title={title}",handler.HandleByTitle).Methods("GET")
	myRouter.HandleFunc("/author={author}",handler.HandleByAuthor).Methods("GET")
	myRouter.HandleFunc("/page/min={min}&max={max}",handler.HandleByPageRange).Methods("GET")
	myRouter.HandleFunc("/lang={lang}",handler.HandleByLanguage).Methods("GET")
	myRouter.HandleFunc("/id={id}",handler.HandleByID).Methods("GET")

	myRouter.HandleFunc("/book",handler.HandlePost).Methods("POST")

	myRouter.HandleFunc("/id={id}",handler.HandlePatch).Methods("PATCH")

	myRouter.HandleFunc("/id={id}",handler.HandleDelete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080",myRouter)) // run server

}

func main() {

	if data.SetUp(){
		handleRequest()

	}else{
		os.Exit(1)
		
	}
}
