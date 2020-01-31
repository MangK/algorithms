package file

import (
	"fmt"
	"os"
	"os/user"
)

func IssetOrCreatDownLoadDir(path string) {
	os.MkdirAll(path, os.ModePerm)
}

func HasFileOrDropDir(path string) {
	fmt.Println("hahahah")
}

func GetUserWorkSpace() string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	return user.HomeDir + "/dfDownload"
}
