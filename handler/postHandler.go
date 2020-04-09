package handler
import (
	"../books"
	"../logic"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HandleFprintf(err error,n int)  {
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("%d bytes written\n",n)
}
func PrintErr(writer http.ResponseWriter,statusCode int, errMessage string){
	writer.WriteHeader(statusCode)
	n,err := fmt.Fprintf(writer,errMessage)
	HandleFprintf(err,n)
}
func respondBackPost(writer http.ResponseWriter) int {
	data,err := json.MarshalIndent(books.BOOKS.Books[len(books.BOOKS.Books)-1],"","  ")
	if err!=nil{
		fmt.Println(err)
		PrintErr(writer,http.StatusInternalServerError,"Error occurred while processing")
		return 0
	}else{
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		n,err := writer.Write(data)
		if err != nil{
			fmt.Println(err)
		}
		return n
	}
}
func HandlePost(writer http.ResponseWriter,r *http.Request)  {
	var newBook books.Book
	if books.Failed{
		PrintErr(writer,http.StatusInternalServerError,"Error occurred while processing")
	}else{
		reqBody,err := ioutil.ReadAll(r.Body)
		if err!=nil{
			fmt.Println(err)
			PrintErr(writer,http.StatusInternalServerError,"Error occurred while processing")
		}else{
			err := json.Unmarshal(reqBody,&newBook)
			if err!=nil{
				PrintErr(writer,http.StatusBadRequest,"Please enter information according to format")
			}else {
				failed := logic.AddABook(newBook)
				if failed{
					PrintErr(writer,http.StatusBadRequest,"Please enter valid integer for number of pages")
				}else{
					respondBackPost(writer)
				}
			}

		}
	}

}
