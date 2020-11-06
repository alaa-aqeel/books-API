package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"../template"
	"path/filepath"
	"os"
)

var BOOKS template.Books

func readFile(fileName string) *os.File {
	jsonFile,err := os.Open(fileName)
	if err!=nil{
		fmt.Println(err)
		defer jsonFile.Close()
		return nil
	}else{
		fmt.Println("Successfully opened file",fileName)
		return jsonFile
	}
}

func parseBooks(fp *os.File)  bool{
	byteValue, _ := ioutil.ReadAll(fp)
	err := json.Unmarshal(byteValue,&BOOKS)
	if err!=nil{
		fmt.Println(err)
		return false
	}else{
		UpdateID()
		return true
	}
}
func UpdateID()  {
	for i:=0;i< len(BOOKS.Books);i++{
		BOOKS.Books[i].Id = int64(i) +1
	}
}

func SetUp() bool{
	absPath, _ := filepath.Abs("books.json")
	jsonFile := readFile(absPath) //may err
	if jsonFile!=nil{
		return parseBooks(jsonFile)
	}else{
		return false
	}
}

