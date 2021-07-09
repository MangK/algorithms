package http

import (
	"fmt"
	"net/http"
)

var ServerCfg Server

func Start(cfg Server) {
	http.HandleFunc("/", Index)
	http.HandleFunc("/upload", Upload)
	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/setClipboard", SetClipboard)

	addr, err := GetLocalIP()
	if err != nil {
		panic(err)
	}

	serverUrl := "http://" + addr + ":" + cfg.Port

	fmt.Println("Server at:", serverUrl)

	ServerCfg = cfg
	ServerCfg.Addr = serverUrl

	err = http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		panic(err)
	}
}
