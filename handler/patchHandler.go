package handler

import (
	"../template"
	"../logic"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"../data"
)

func respondBackPatch(writer http.ResponseWriter,failed bool,theBook template.Books) int{
	if failed{
		PrintErr(writer,http.StatusBadRequest,"Please enter a valid integer as ID")
		return 0
	}else{
		toSend,err := json.MarshalIndent(theBook,"","  ")
		if err!=nil{
			fmt.Println(err)
			PrintErr(writer,http.StatusInternalServerError,"Error occurred while processing")
			return 0
		}else{
			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			n,err := writer.Write(toSend)
			if err != nil{
				fmt.Println(err)
			}
			return n
		}
	}

}
func HandlePatch(writer http.ResponseWriter,r *http.Request)  {
	var newBook template.Book
	if data.Failed{
		PrintErr(writer,http.StatusInternalServerError,"Error occurred while processing")
	}else{
		vars:= mux.Vars(r)
		id :=  vars["id"]
		reqBody,err := ioutil.ReadAll(r.Body)
		if err!=nil{
			fmt.Println(err)
			PrintErr(writer,http.StatusInternalServerError,"Error occurred while processing")
		}else{
			err := json.Unmarshal(reqBody,&newBook)
			if err!=nil{
				PrintErr(writer,http.StatusBadRequest,"Please enter information according to format")
			}else {
				newBook,failed := logic.UpdateBook(id,newBook)
				respondBackPatch(writer,failed,newBook)
			}
		}
	}

}
