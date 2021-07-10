package main

import (
	myFile "DragFile/file"
	"DragFile/params"
	server_http "DragFile/server/http"
	"DragFile/ui"
	"fmt"
)

func main() {
	param := params.GetParams()
	addr, err := server_http.GetLocalIP()
	if err != nil {
		panic(err)
	}
	param.Addr = "http://" + addr + ":" + param.Port

	defer myFile.HasFileOrDropDir(param.Dir)

	fmt.Println("Current Config:", param)
	ui.Run(param)
}

func start(param server_http.Server) {
	server_http.Start(param)
}
