package handler

import (
	"../books"
	"../logic"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func respondBackPatch(writer http.ResponseWriter,failed bool,theBook books.Books) int{
	if failed{
		writer.WriteHeader(http.StatusBadRequest)
		n,err := fmt.Fprintf(writer,"Please enter a valid integer ID")
		HandleFprintf(err,n)
		return 0
	}else{
		data,err := json.MarshalIndent(theBook,"","  ")
		if err!=nil{
			fmt.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			n,err := fmt.Fprintf(writer,"Error occurred")
			HandleFprintf(err,n)
			return 0
		}else{
			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			n,err := writer.Write(data)
			if err != nil{
				fmt.Println(err)
			}
			return n
		}
	}

}
func HandlePatch(writer http.ResponseWriter,r *http.Request)  {
	var newBook books.Book
	if books.Failed{
		writer.WriteHeader(http.StatusInternalServerError)
		n,err := fmt.Fprintf(writer,"Error occurred")
		HandleFprintf(err,n)
	}else{
		vars:= mux.Vars(r)
		id :=  vars["id"]
		reqBody,err := ioutil.ReadAll(r.Body)
		if err!=nil{
			fmt.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			n,err := fmt.Fprintf(writer,"Error occurred")
			HandleFprintf(err,n)
		}else{
			err := json.Unmarshal(reqBody,&newBook)
			if err!=nil{
				writer.WriteHeader(http.StatusBadRequest)
				n,err := fmt.Fprintf(writer,"Please enter information according to format")
				HandleFprintf(err,n)
			}else {
				newBook,failed := logic.UpdateBook(id,newBook)
				respondBackPatch(writer,failed,newBook)
			}
		}
	}

}
