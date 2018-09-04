package main

import (
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"log"
	"syscall"
	"godesk/component/notifyicon"
	tool "github.com/GeertJohan/go.rice"
	"github.com/sciter-sdk/go-sciter/rice"
	"github.com/lxn/win"
)
//下面注释不要删除，使用go generate加入程序图标、信息（main.rc）以及打包静态资源
//go:generate windres -o main-res.syso main.rc
//go:generate rice embed-go

func main() {
	//调用系统函数设置当前进程对高DPi的支持方式,否则系统缩放后会模糊
	shcore := syscall.NewLazyDLL("Shcore.dll").NewProc("SetProcessDpiAwareness")
	err := shcore.Find()
	if err == nil {
		shcore.Call(uintptr(2))
	}

	screenWidth := float32(win.GetSystemMetrics(0))
	screenHeight := float32(win.GetSystemMetrics(1))

	//创建window窗口
	//参数一表示创建窗口的样式
	//SW_TITLEBAR 顶层窗口，有标题栏
	//SW_RESIZEABLE 可调整大小
	//SW_CONTROLS 有最小/最大按钮
	//SW_MAIN 应用程序主窗口，关闭后其他所有窗口也会关闭
	//SW_ENABLE_DEBUG 可以调试
	//参数二表示创建窗口的矩形, 参数代表了左上角，右下角的坐标
	w, err := window.New(
		//sciter.SW_RESIZEABLE|
		sciter.SW_ALPHA|
		sciter.SW_ENABLE_DEBUG,
		&sciter.Rect{Left: int32(screenWidth * 0.1), Top: int32(screenHeight * 0.1), Right: int32(screenWidth * 0.9), Bottom: int32(screenHeight * 0.9)})
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

	//cmd := exec.Command("tasklist")
	//out, err := cmd.CombinedOutput()
	//log.Print(string(out))
	//if err != nil {
	//	log.Fatal(err)
	//}

		//加载文件
		w.LoadFile("rice://resource/view/main.html")

	//获取根元素
	//root, _ := w.GetRootElement()
	//元素加载资源
	//img, _ := root.SelectById("logoImg");
	//img.AttachEventHandler(&sciter.EventHandler{
	//	//OnDataArrived 当资源被加载但未使用时调用
	//	OnDataArrived: func(he *sciter.Element, params *sciter.DataArrivedParams) bool {
	//		//设置属性，给img标签设置src
	//		he.SetAttr("src", params.Uri());
	//		return false;
	//	},
	//});
	//img.Load("http://mat1.gtimg.com/www/images/qq2012/qqLogoFilter.png", sciter.RT_DATA_IMAGE);
		//显示窗口
		w.Show()
		//运行窗口，进入消息循环
		w.Run()
}
