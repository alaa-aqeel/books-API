package utils

import (
	"../template"
	"encoding/json"
	"fmt"
	"net/http"
)

var response template.Response
var single template.SingleResponse

func clearResponse()  {
	response.Books = make([]template.Book, 0)
	response.Error = ""
}

func SendText(writer http.ResponseWriter,statusCode int, message string){
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	toSend := []byte("{\"message\": \"" + message + "\" }")
	_, err := writer.Write(toSend)
	
	if err != nil{
		fmt.Println(err)
	}
}

func SendError(writer http.ResponseWriter,statusCode int, errMessage string){
	clearResponse()
	response.Error = errMessage

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	toSend,err := json.MarshalIndent(response,"","  ")

	if err!=nil{
		fmt.Println(err)
		SendText(writer, 500, "Error occurred while processing")

	}else{
		_,err := writer.Write(toSend)
		if err != nil{
			fmt.Println(err)
		}
	}
}

func SendBooks(writer http.ResponseWriter,statusCode int, books template.Books){
	clearResponse()
	response.Books = books.Books

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	toSend,err := json.MarshalIndent(response,"","  ")

	if err!=nil{
		fmt.Println(err)
		SendError(writer, 500, "Error occurred while processing")
	}else{
		_,err := writer.Write(toSend)
		if err != nil{
			fmt.Println(err)
		}
	}
}

func SendBook(writer http.ResponseWriter,statusCode int, book template.Book, errMessage string){
	single.Book = template.Book{}
	single.Error = ""

	if errMessage!=""{
		single.Error = errMessage

	}
	single.Book = book

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	
	toSend,err := json.MarshalIndent(single,"","  ")
	
	if err!=nil{
		fmt.Println(err)
		SendText(writer, 500, "Error occurred while processing")

	}else{
		_,err := writer.Write(toSend)

		if err != nil{
			fmt.Println(err)
		}
	}
}
