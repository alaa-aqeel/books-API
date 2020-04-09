package handler

import (
	"../books"
	"../logic"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)


func HandleDelete(writer http.ResponseWriter,r *http.Request){
	if books.Failed{
		PrintErr(writer,http.StatusInternalServerError,"Error occurred while processing")
	}else{
		vars:= mux.Vars(r)
		id :=  vars["id"]
		failed,found := logic.DeleteABook(id)
		if failed{
			PrintErr(writer,http.StatusBadRequest,"Please enter a valid integer as ID")
		}else if !found{
			PrintErr(writer,http.StatusBadRequest,"No Book found with id: "+id)
		}else{
			writer.WriteHeader(http.StatusOK)
			n,err := fmt.Fprintf(writer,"Book with id: "+id+" is deleted")
			HandleFprintf(err,n)
		}
	}
}
