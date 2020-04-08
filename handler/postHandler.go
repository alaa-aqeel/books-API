package handler
import (
	"../books"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handleFprintf(err error,n int)  {
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("%d bytes written\n",n)
}
func respondBack(writer http.ResponseWriter) int {
	data,err := json.MarshalIndent(books.BOOKS.Books[len(books.BOOKS.Books)-1],"","  ")
	if err!=nil{
		fmt.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		n,err := fmt.Fprintf(writer,"Error occurred")
		handleFprintf(err,n)
		return 0
	}else{
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		n,err := writer.Write(data)
		if err != nil{
			panic(err)
		}
		return n
	}
}
func HandlePost(writer http.ResponseWriter,r *http.Request)  {
	var newBook books.Book
	reqBody,err := ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		n,err := fmt.Fprintf(writer,"Error occurred")
		handleFprintf(err,n)
	}else{
		err := json.Unmarshal(reqBody,&newBook)
		if err!=nil{
			writer.WriteHeader(http.StatusBadRequest)
			n,err := fmt.Fprintf(writer,"Please enter information according to format")
			handleFprintf(err,n)
		}else {
			books.AddBook(newBook)
			respondBack(writer)
		}

	}
}
