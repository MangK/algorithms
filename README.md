# 编译windows

- brew install mingw-w64
- CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64 go build

# 可执行文件隐藏命令行窗口

- Windows
    - -ldflags -H=windowsgui
- Mac
    - 暂无

# 打包

- fyne package -os linux -icon icon.png
- fyne package -os windows -icon icon.png