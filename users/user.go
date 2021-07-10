package users

import (
	"DragFile/server/http"
	"sync"
)

type User struct {
	http.Server
}

var Users = sync.Map{}
var Change = make(chan bool)

//func GetChangeChan() ch {
//
//}
