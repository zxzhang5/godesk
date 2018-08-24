package main

import (
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"log"
	"syscall"
)

func main() {
	//调用系统函数设置当前进程对高DPi的支持方式,否则系统缩放后会模糊
	syscall.NewLazyDLL("Shcore.dll").NewProc("SetProcessDpiAwareness").Call(uintptr(2))
	//创建window窗口
	//参数一表示创建窗口的样式
	//SW_TITLEBAR 顶层窗口，有标题栏
	//SW_RESIZEABLE 可调整大小
	//SW_CONTROLS 有最小/最大按钮
	//SW_MAIN 应用程序主窗口，关闭后其他所有窗口也会关闭
	//SW_ENABLE_DEBUG 可以调试
	//参数二表示创建窗口的矩形, 参数代表了左上角，右下角的坐标
	w, err := window.New(sciter.SW_TITLEBAR|
		sciter.SW_RESIZEABLE|
		sciter.SW_CONTROLS|
		sciter.SW_MAIN|
		sciter.SW_ENABLE_DEBUG,
		&sciter.Rect{Left: 500, Top: 300, Right: 1300, Bottom: 900})
	if err != nil {
		log.Fatal(err)
	}

	//加载文件
	w.LoadFile("resources/views/main.html")
	//设置标题
	w.SetTitle("GoDesk")
	//显示窗口
	w.Show()
	//运行窗口，进入消息循环
	w.Run()
}
