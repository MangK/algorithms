package main

import (
	myFile "DragFile/file"
	"DragFile/params"
	server_http "DragFile/server/http"
	"DragFile/ui"
	"fmt"
)

var Exit = make(chan bool)

func main() {
	param := params.GetParams()
	addr, err := server_http.GetLocalIP()
	if err != nil {
		panic(err)
	}
	param.Addr = "http://" + addr + ":" + param.Port

	defer myFile.HasFileOrDropDir(param.Dir)

	fmt.Println("Current Config:", param)

	go server_http.Start(param)

	ui.Run(param)

}
