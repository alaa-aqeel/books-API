package main

import (
	"../sth/handler"
	"./books"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handleRequest()  {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",handler.HandleHome).Methods("GET")
	myRouter.HandleFunc("/all",handler.AllBooks).Methods("GET")
	myRouter.HandleFunc("/title/{title}",handler.HandleByTitle).Methods("GET")
	myRouter.HandleFunc("/author/{author}",handler.HandleByAuthor).Methods("GET")
	myRouter.HandleFunc("/page/min/{min}/max/{max}",handler.HandleByPageRange).Methods("GET")
	myRouter.HandleFunc("/lang/{lang}",handler.HandleByLanguage).Methods("GET")
	myRouter.HandleFunc("/book",handler.HandlePost).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080",myRouter)) // run server

}
func main() {
	books.Run()
	handleRequest()
}
