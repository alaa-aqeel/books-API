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
	r.Body = http.MaxBytesReader(writer, r.Body, 2<<10)
	reqBody,err := ioutil.ReadAll(r.Body)

	if err!=nil{
			fmt.Println(err)
			utils.SendBook(writer, 400, newBook, "Request body was too large to process")

	}else{
			err := json.Unmarshal(reqBody,&newBook)

			if err!=nil{
				utils.SendBook(writer, 400, newBook, "Invalid request body passed")

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
