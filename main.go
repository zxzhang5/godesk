package main

import (
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"log"
	"syscall"
	"godesk/component/notifyicon"
	tool "github.com/GeertJohan/go.rice"
	rice "github.com/sciter-sdk/go-sciter/rice"
)
//下面注释不要删除，使用go generate加入程序图标、信息（main.rc）以及打包静态资源
//go:generate windres -o main-res.syso main.rc
//go:generate rice embed-go

func main() {
	//调用系统函数设置当前进程对高DPi的支持方式,否则系统缩放后会模糊
	shcore := syscall.NewLazyDLL("Shcore.dll").NewProc("SetProcessDpiAwareness")
	err := shcore.Find()
	if err != nil {
		//加载dll出错
		//log.Print(err)
	}else{
		shcore.Call(uintptr(2))
	}

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
		sciter.SW_POPUP|
		sciter.SW_ENABLE_DEBUG,
		&sciter.Rect{Left: 500, Top: 300, Right: 1300, Bottom: 900})
	if err != nil {
		log.Fatal(err)
	}

	rice.HandleDataLoad(w.Sciter)
	tool.MustFindBox("resource")

	//设置托盘图标
	ni, err := notifyicon.New("icon.ico")
	if err != nil {
		log.Fatal(err)
	}
	// 消息提示
	//err = ni.ShowInfo("运行时初始提示信息", "信息详情")
	//if err != nil {
	//	log.Fatal(err)
	//}
	log.Print(ni)

	//加载文件
	w.LoadFile("rice://resource/view/main.html")
	//设置标题
	w.SetTitle("GoDesk演示")

	//显示窗口
	w.Show()
	//运行窗口，进入消息循环
	w.Run()
}
