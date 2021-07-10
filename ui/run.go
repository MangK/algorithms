package ui

import (
	server_http "DragFile/server/http"
	"DragFile/users"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/skip2/go-qrcode"
	"image"
	"net/url"
	"time"
)

func Run(param server_http.Server) {
	a := app.NewWithID("x.Mangk.DragFile")
	a.SetIcon(theme.FyneLogo())
	w := a.NewWindow("DragFile--文件传输")
	w.Resize(fyne.Size{800, 450})

	w.SetContent(buildContainer(param))

	w.ShowAndRun()
}

func buildQrcode(server string) image.Image {
	qrObj, err := qrcode.New(server, qrcode.Highest)
	if err != nil {
		panic(err)
	}
	qrObj.DisableBorder = true

	return qrObj.Image(100)
}

func buildContainer(param server_http.Server) *fyne.Container {
	qr := canvas.NewImageFromImage(buildQrcode(param.Addr))
	qr.FillMode = canvas.ImageFillContain
	qr.SetMinSize(fyne.Size{100, 100})
	bottom := container.NewHBox()
	go func() {
		for true {
			t := time.Now().Format("150405")
			u := server_http.Server{
				Name: t,
				Port: t,
				Addr: t,
				Pwd:  t,
				Dir:  t,
			}
			users.Users.LoadOrStore(u.Name, u)
			users.Change <- true
			time.Sleep(time.Second * 10)
		}
	}()
	go func(c *fyne.Container) {
		for true {
			<-users.Change
			obj := []fyne.CanvasObject{}
			users.Users.Range(func(key, value interface{}) bool {
				v := value.(server_http.Server)
				obj = append(obj, widget.NewLabel(v.Name))
				return true
			})
			c.Objects = obj
			c.Refresh()
		}
	}(bottom)
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("Scan QR Code To Send File", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		qr,
		container.NewHBox(
			widget.NewHyperlink("DragFile", parseURL("DragFile.banmal.cn")),
			widget.NewLabel("-"),
			widget.NewHyperlink(param.Addr, parseURL(param.Addr)),
			widget.NewLabel("-"),
			widget.NewHyperlink("MangK", parseURL("banmal.cn")),
		),
		bottom,
	))
}

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}
