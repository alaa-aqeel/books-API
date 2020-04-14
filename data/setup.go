package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"../template"
	"path/filepath"
)

var Failed bool = false
var BOOKS template.Books
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

func SetUp() {
	absPath, _ := filepath.Abs("books.json")
	jsonFile := readFile(absPath) //may err
	if jsonFile !=nil{
		parseBooks(jsonFile)
	}else{
		os.Exit(1)
	}
}

