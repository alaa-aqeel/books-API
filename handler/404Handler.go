package handler

import (
	"net/http"
	"../utils"
	"fmt"
)


func Handle404(writer http.ResponseWriter,r *http.Request){
	fmt.Println("Unknown Endpoint Hit")
	utils.SendText(writer, 404, "Page not found")
}
