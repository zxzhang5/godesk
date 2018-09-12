package main

import (
	"github.com/lxn/walk"
	"log"
	"net/http"
	"github.com/braintree/manners"
	tool "github.com/GeertJohan/go.rice"
	"godesk/component/tomlconfig"
	"golang.org/x/text/message"
	"fmt"
	"github.com/kataras/iris"
)

//下面注释不要删除，使用go generate加入程序图标、信息（main.rc）以及打包静态资源、本地化文件
//go:generate windres -o main-res.syso main.rc
//go:generate rice embed-go
//go:generate gotext -srclang=zh update -out=catalog.go -lang=zh,en

//func main() {
//
//	http.Handle("/", http.FileServer(tool.MustFindBox("resource").HTTPBox()))
//	http.ListenAndServe("127.0.0.1:8000", nil)
//
//	//调用系统函数设置当前进程对高DPi的支持方式,否则系统缩放后会模糊
//	shcore := syscall.NewLazyDLL("Shcore.dll").NewProc("SetProcessDpiAwareness")
//	err := shcore.Find()
//	if err == nil {
//		shcore.Call(uintptr(2))
//	}
//	screenWidth := float32(win.GetSystemMetrics(0))
//	screenHeight := float32(win.GetSystemMetrics(1))
//
//	//创建window窗口
//	//参数一表示创建窗口的样式
//	//SW_TITLEBAR 顶层窗口，有标题栏
//	//SW_RESIZEABLE 可调整大小
//	//SW_CONTROLS 有最小/最大按钮
//	//SW_MAIN 应用程序主窗口，关闭后其他所有窗口也会关闭
//	//SW_ENABLE_DEBUG 可以调试
//	//参数二表示创建窗口的矩形, 参数代表了左上角，右下角的坐标
//	w, err := window.New(
//		//sciter.SW_RESIZEABLE|
//		sciter.SW_ALPHA|
//		sciter.SW_ENABLE_DEBUG,
//		&sciter.Rect{Left: int32(screenWidth * 0.1), Top: int32(screenHeight * 0.1), Right: int32(screenWidth * 0.9), Bottom: int32(screenHeight * 0.9)})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	rice.HandleDataLoad(w.Sciter)
//	tool.MustFindBox("resource")
//
//	//设置托盘图标
//	ni, err := notifyicon.New("icon.ico")
//	if err != nil {
//		log.Fatal(err)
//	}
//	// 消息提示
//	//err = ni.ShowInfo("运行时初始提示信息", "信息详情")
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	log.Print(ni)
//
//	//cmd := exec.Command("tasklist")
//	//out, err := cmd.CombinedOutput()
//	//log.Print(string(out))
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//
//		//加载文件
//		w.LoadFile("rice://resource/view/main.html")
//
//	//获取根元素
//	//root, _ := w.GetRootElement()
//	//元素加载资源
//	//img, _ := root.SelectById("logoImg");
//	//img.AttachEventHandler(&sciter.EventHandler{
//	//	//OnDataArrived 当资源被加载但未使用时调用
//	//	OnDataArrived: func(he *sciter.Element, params *sciter.DataArrivedParams) bool {
//	//		//设置属性，给img标签设置src
//	//		he.SetAttr("src", params.Uri());
//	//		return false;
//	//	},
//	//});
//	//img.Load("http://mat1.gtimg.com/www/images/qq2012/qqLogoFilter.png", sciter.RT_DATA_IMAGE);
//		//显示窗口
//		w.Show()
//		//运行窗口，进入消息循环
//		w.Run()
//}

type MyWindow struct {
	*walk.MainWindow
	ni *walk.NotifyIcon
}

func NewMyWindow() *MyWindow {
	mw := new(MyWindow)
	var err error
	mw.MainWindow, err = walk.NewMainWindow()
	checkError(err)
	return mw
}

func (mw *MyWindow) init() {
	http.Handle("/", http.FileServer(tool.MustFindBox("resource").HTTPBox()))
}

func (mw *MyWindow) RunHttpServer() error {
	return manners.ListenAndServe(":8080", http.DefaultServeMux)
}

func (mw *MyWindow) AddNotifyIcon() {
	var err error
	mw.ni, err = walk.NewNotifyIcon()
	checkError(err)
	mw.ni.SetVisible(true)

	icon, err := walk.Resources.Icon("icon.ico")
	checkError(err)
	mw.SetIcon(icon)
	mw.ni.SetIcon(icon)

	startAction := mw.addAction(nil, "start")
	stopAction := mw.addAction(nil, "stop")
	stopAction.SetEnabled(false)
	startAction.Triggered().Attach(func() {
		go func() {
			err := mw.RunHttpServer()
			if err != nil {
				mw.msgbox("start", "start http server failed.", walk.MsgBoxIconError)
				return
			}
		}()
		startAction.SetChecked(true)
		startAction.SetEnabled(false)
		stopAction.SetEnabled(true)
		mw.msgbox("start", "start http server success.", walk.MsgBoxIconInformation)
	})

	stopAction.Triggered().Attach(func() {
		ok := manners.Close()
		if !ok {
			mw.msgbox("stop", "stop http server failed.", walk.MsgBoxIconError)
		} else {
			stopAction.SetEnabled(false)
			startAction.SetChecked(false)
			startAction.SetEnabled(true)
			mw.msgbox("stop", "stop http server success.", walk.MsgBoxIconInformation)
		}
	})

	helpMenu := mw.addMenu("help")
	mw.addAction(helpMenu, "help").Triggered().Attach(func() {
		walk.MsgBox(mw, "help", "http://127.0.0.1:8080", walk.MsgBoxIconInformation)
	})

	mw.addAction(helpMenu, "about").Triggered().Attach(func() {
		walk.MsgBox(mw, "about", "http server.", walk.MsgBoxIconInformation)
	})

	mw.addAction(nil, "exit").Triggered().Attach(func() {
		mw.ni.Dispose()
		mw.Dispose()
		walk.App().Exit(0)
	})

}

func (mw *MyWindow) addMenu(name string) *walk.Menu {
	helpMenu, err := walk.NewMenu()
	checkError(err)
	help, err := mw.ni.ContextMenu().Actions().AddMenu(helpMenu)
	checkError(err)
	help.SetText(name)

	return helpMenu
}

func (mw *MyWindow) addAction(menu *walk.Menu, name string) *walk.Action {
	action := walk.NewAction()
	action.SetText(name)
	if menu != nil {
		menu.Actions().Add(action)
	} else {
		mw.ni.ContextMenu().Actions().Add(action)
	}

	return action
}

func (mw *MyWindow) msgbox(title, message string, style walk.MsgBoxStyle) {
	mw.ni.ShowInfo(title, message)
	walk.MsgBox(mw, title, message, style)
}

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "Hello world!")
		ctx.View("main.html")
	})
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.TOML("iniconfig/iris.tml")))
	//pkgInfo := x.iprog.Package("golang.org/x/text/message")
	//if pkgInfo == nil {
	//	log.Print("golang.org/x/text/message is not imported")
	//}
	//cfg := iniconfig.Load()
	//iniconfig.Set(cfg,"bbb","111")
	//iniconfig.Set(cfg,"aaa","呵呵")
	//iniconfig.Save(cfg)
	lang  := iniconfig.Get("ui.lang")
	fmt.Print(lang)
	tag := message.MatchLanguage(lang)
	p := message.NewPrinter(tag)
	msg:=p.Sprintln("Hello %s!\n", "howdy")
log.Print(msg)
	mw := NewMyWindow()
	mw.init()
	mw.AddNotifyIcon()
	mw.Run()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
