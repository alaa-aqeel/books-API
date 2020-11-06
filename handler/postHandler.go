package handler
import (
	"../template"
	"../logic"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"../utils"
)

func HandlePost(writer http.ResponseWriter,r *http.Request) {
	fmt.Println("Add Book Endpoint Hit")
	var newBook template.Book
	reqBody,err := ioutil.ReadAll(r.Body)
	if err!=nil{
			fmt.Println(err)
			utils.SendBook(writer, 500, newBook, "Error occurred while processing")

	}else{
			err := json.Unmarshal(reqBody,&newBook)
			if err!=nil{
				utils.SendBook(writer, 500, newBook, "Error occurred while processing")

			}else {
				book, message := logic.AddABook(newBook)
				if message != ""{
					utils.SendBook(writer, 400, book, message)
	
				}else{
					utils.SendBook(writer, 201, book, message)
				}
			}

		}

}
