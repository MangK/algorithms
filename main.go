package main

import (
	myFile "DragFile/file"
	"DragFile/params"
	server_http "DragFile/server/http"
	"fmt"
)

func main() {
	param := params.GetParams()
	defer myFile.HasFileOrDropDir(param.Dir)

	fmt.Println("Current Config:", param)
	server_http.Start(param)
}
