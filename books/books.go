package books

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)
type Book struct {
	Id  int64  `json:"id"`
	Author string `json:"author"`
	Language string `json:"language"`
	Link string `json:"link"`
	Pages int64 `json:"pages"`
	Title string `json:"title"`
}
type Books struct {
	Books []Book 
}

type Error struct {
	Error string `json:"error"`
}

var Failed bool = false
var BOOKS Books
var ERROR Error
func readFile(fileName string) *os.File {
	jsonFile,err := os.Open(fileName)
	if err!= nil{
		fmt.Println(err)
		Failed = true
		defer jsonFile.Close()
		return nil
	}else{
		fmt.Println("Successfully opened file",fileName)
		return jsonFile
	}
}

func parseBooks(fp *os.File)  {
	byteValue, _ := ioutil.ReadAll(fp)
	err := json.Unmarshal(byteValue,&BOOKS)
	if err!=nil{
		fmt.Println(err)
		Failed = true
	}else {
		UpdateID()
	}
}
func UpdateID()  {
	for i:=0;i< len(BOOKS.Books);i++{
		BOOKS.Books[i].Id = int64(i) +1
	}
}
func AddBook(book Book)  {
	BOOKS.Books = append(BOOKS.Books,book)
}
func Run() {
	jsonFile := readFile("books.json")
	if jsonFile !=nil{
		parseBooks(jsonFile)
	}
}
