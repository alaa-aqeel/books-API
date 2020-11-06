package handler

import (
	"../template"
	"../logic"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"../utils"
)

func HandlePatch(writer http.ResponseWriter,r *http.Request)  {
	fmt.Println("Update Book Endpoint Hit")

	var newBook template.Book

	vars:= mux.Vars(r)
	id :=  vars["id"]
	reqBody,err := ioutil.ReadAll(r.Body)
	
	if err!=nil{
		fmt.Println(err)
		utils.SendBook(writer, 500, newBook, "Error occurred while processing")

	}else{
		err := json.Unmarshal(reqBody,&newBook)
		if err!=nil {
			utils.SendBook(writer, 500, newBook, "Error occurred while processing")

		}else {
			book,message := logic.UpdateBook(id,newBook)
			if message != ""{
				utils.SendBook(writer, 400, book, message)

			}else{
				utils.SendBook(writer, 200, book, message)
			}
		}
	}

}
