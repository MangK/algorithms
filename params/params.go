package params

import (
	myFile "DragFile/file"
	"DragFile/server/http"
	"flag"
	"time"
)

var (
	name        string
	port        string
	password    string
	fielSaveDir string
)

func GetParams() http.Server {
	t := time.Now().Format("0405")

	flag.StringVar(&name, "n", "mangk"+t, "Server Name")
	flag.StringVar(&port, "p", "5488", "Http Server Port")
	flag.StringVar(&password, "P", t, "Http Server Password")
	flag.StringVar(&fielSaveDir, "D", myFile.GetUserWorkSpace(), "File Save Dir")

	flag.Parse()

	// TODO 这里可以先注释掉，等到真正要写入文件时才去创建文件夹
	// if fielSaveDir != myFile.GetUserWorkSpace() {
	myFile.IssetOrCreatDownLoadDir(fielSaveDir)
	// }

	return http.Server{
		Name: name,
		Port: port,
		Addr: "",
		Pwd:  password,
		Dir:  fielSaveDir,
	}
}
