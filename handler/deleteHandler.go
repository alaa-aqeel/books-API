package handler

import (
	"../logic"
	"fmt"
	"net/http"
	"../utils"
	"github.com/gorilla/mux"
)


func HandleDelete(writer http.ResponseWriter,r *http.Request){
	fmt.Println("Delete Book Endpoint Hit")

	vars:= mux.Vars(r)
	id :=  vars["id"]

	success, err := logic.DeleteABook(id)

	if !success && err != ""{
		utils.SendText(writer, 400, err)

	}else{
		utils.SendText(writer, 200, "Book was successfully deleted")
	}
}
