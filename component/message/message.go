package message

import (
	"syscall"
	"github.com/lxn/win"
	"github.com/lxn/walk"
)

//显示系统对话框,返回值6 yes 7 no
func Show(title string, text string, style uint32) int {
	textUtf,_ := syscall.UTF16PtrFromString(text)
	titleUtf,_ := syscall.UTF16PtrFromString(title)
	ret := win.MessageBox(0,
		textUtf,
		titleUtf,
		style)
	return int(ret)
}

//成功提示
func Success(title string, text string){
	Show(title, text, win.MB_OK|win.MB_ICONINFORMATION)
}

//错误提示不退出
func Error(title string, text string){
	Show(title, text, win.MB_OK|win.MB_ICONERROR)
}

//错误提示并退出
func Fatal(title string, text string){
	Show(title, text, win.MB_OK|win.MB_ICONERROR)
	walk.App().Exit(0)
}

//询问对话框，返回值true是，false否
func Confirm(title string, text string) bool{
	ret := Show(title, text, win.MB_YESNO|win.MB_ICONQUESTION)
	if ret == 6{
		return true
	}else{
		return false
	}
}