package main

import (
	myFile "DragFile/file"
	"DragFile/params"
	"DragFile/server/http"
	"fmt"
)

func main() {
	params := params.GetParams()
	defer myFile.HasFileOrDropDir(params.Dir)

	fmt.Println("Current Config:", params)
	http.Start(params)
}
