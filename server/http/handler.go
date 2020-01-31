package http

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(indexTmp()))
}

func Ping(w http.ResponseWriter, r *http.Request) {
	cfg := ServerCfg
	cfg.Dir = ""
	cfg.Pwd = ""
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	w.Write(jsonData)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		f, err := os.OpenFile(ServerCfg.Dir+"/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		io.Copy(f, file)
	}
	w.Write([]byte("<script>window.location.replace('" + ServerCfg.Addr + "');</script>"))
}
