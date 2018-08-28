package notifyicon

import (
	"log"
	"github.com/lxn/walk"
)

//图盘图标设置程序，结构体NotifyIcon引用需要带包名
func Config(filename string) (*walk.NotifyIcon, error){
	icon, err := walk.Resources.Icon(filename)
	if err != nil {
		log.Fatal(err)
	}
	// 新建托盘图标
	ni, err := walk.NewNotifyIcon()
	if err != nil {
		log.Fatal(err)
	}
	// 设置图标icon
	if err := ni.SetIcon(icon); err != nil {
		log.Fatal(err)
	}
	if err := ni.SetToolTip("鼠标焦点提示"); err != nil {
		log.Fatal(err)
	}

	// 鼠标左键点击
	ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button != walk.LeftButton {
			return
		}
		//err := ni.ShowCustom("鼠标左键点击","鼠标左键点击提示")
		//if err != nil {
		//	log.Fatal(err)
		//}
	})

	// 新建右击菜单选项
	exitAction := walk.NewAction()
	if err := exitAction.SetText("退出"); err != nil {
		log.Fatal(err)
	}
	// 定义点击选项触发事件
	exitAction.Triggered().Attach(func() { walk.App().Exit(0) })

	// 右击菜单加入选项
	if err := ni.ContextMenu().Actions().Add(exitAction); err != nil {
		log.Fatal(err)
	}

	// 显示托盘图标
	if err := ni.SetVisible(true); err != nil {
		log.Fatal(err)
	}

	return ni, nil
}
