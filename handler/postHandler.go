package handler
import (
	"../template"
	"../logic"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"../data"
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
	toSend,err := json.MarshalIndent(data.BOOKS.Books[len(data.BOOKS.Books)-1],"","  ")
	if err!=nil{
		fmt.Println(err)
		PrintErr(writer,http.StatusInternalServerError,"Error occurred while processing")
		return 0
	}else{
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		n,err := writer.Write(toSend)
		if err != nil{
			fmt.Println(err)
		}
		return n
	}
}
func HandlePost(writer http.ResponseWriter,r *http.Request)  {
	var newBook template.Book
	if data.Failed{
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
