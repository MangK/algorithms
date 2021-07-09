package main

import (
	myFile "DragFile/file"
	"DragFile/params"
	server_http "DragFile/server/http"
	"fmt"
	"github.com/chai2010/winsvc"
	"log"
	"os"
	"path/filepath"
)

var (
	appPath string // 程序的绝对路径
)
func init() {
	var err error
	if appPath, err = winsvc.GetAppPath(); err != nil {
		log.Fatal(err)
	}
	if err := os.Chdir(filepath.Dir(appPath)); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if !winsvc.IsAnInteractiveSession() {
		log.Println("main:", "runService")
		if err := winsvc.RunAsService("AAAAAAAA", start, stop, false); err != nil {
			log.Fatalf("svc.Run: %v\n", err)
		}
		return
	}

	start()
}

func start() {
	param := params.GetParams()
	defer myFile.HasFileOrDropDir(param.Dir)

	fmt.Println("Current Config:", param)
	server_http.Start(param)
}

func stop() {

}
