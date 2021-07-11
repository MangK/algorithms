package ui

import (
	server_http "DragFile/server/http"
	"DragFile/users"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"net/url"
	"time"
)

var Width float32 = 800
var Height float32 = 450

var Exit = make(chan bool)

func Run(param server_http.Server) {
	a := app.NewWithID("x.Mangk.DragFile")
	a.SetIcon(theme.FyneLogo())
	w := a.NewWindow("DragFile--文件传输")
	w.Resize(fyne.Size{Width, Height})

	c := buildContainer(param)
	w.SetContent(c)

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
	userCenter := buildUserCenter()
	// 用来测试增加用户
	//go func() {
	//	for true {
	//		t := time.Now().Format("150405")
	//		u := server_http.Server{
	//			Name: t,
	//			Port: t,
	//			Addr: t,
	//			Pwd:  t,
	//			Dir:  t,
	//		}
	//		users.Users.LoadOrStore(u.Name, u)
	//		users.Change <- true
	//		time.Sleep(time.Second * 10)
	//	}
	//}()
	l := container.NewVScroll(container.NewVBox(
		widget.NewLabelWithStyle("Scan QR Code To Send File", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		qr,
		container.NewCenter(
			container.NewHBox(
				widget.NewHyperlink("DragFile", parseURL("DragFile.banmal.cn")),
				widget.NewLabel("-"),
				widget.NewHyperlink(param.Addr, parseURL(param.Addr)),
				widget.NewLabel("-"),
				widget.NewHyperlink("MangK", parseURL("banmal.cn")),
			),
		),
		canvas.NewLine(color.RGBA{
			R: 0,
			G: 0,
			B: 255,
			A: 255,
		}),
		userCenter,
	))
	l.SetMinSize(fyne.Size{Width, Height})
	c := container.NewCenter(l)
	return c
}

func buildUserCenter() *fyne.Container {
	c := container.NewGridWithColumns(6)
	c.Resize(fyne.Size{
		Width:  500,
		Height: Height,
	})

	// 监测用户变换并刷新
	go func(c *fyne.Container) {
		for true {
			<-users.Change
			obj := []fyne.CanvasObject{}
			users.Users.Range(func(key, value interface{}) bool {
				v := value.(server_http.Server)
				obj = append(obj, buildUserBox(v.Name))
				return true
			})
			c.Objects = obj
			c.Refresh()
		}
	}(c)

	return c
}

/**
用户图标生成模块
*/
func buildUserBox(name string) *fyne.Container {
	x := canvas.NewRadialGradient(color.RGBA{
		R: 50,
		G: 100,
		B: 150,
		A: 255,
	}, color.Transparent)
	x.SetMinSize(fyne.Size{60, 60})

	c := container.NewCenter(
		container.NewVBox(
			x,
			widget.NewLabel(name),
			widget.NewButton(name, func() {
				fmt.Println(time.Now())
			}),
		),
	)
	return c
}

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}
