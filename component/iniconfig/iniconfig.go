package iniconfig

import (
	"github.com/go-ini/ini"
	"godesk/component/message"
	"strings"
)

//加载配置文件
func Load() *ini.File {
	cfg, err := ini.Load("config/config.ini")
	message.CheckFatal(err, "错误","读取配置文件config.ini出错")
	return cfg
}

//保存配置文件
func Save(cfg *ini.File)  {
	err := cfg.SaveTo("config/config.ini")
	message.CheckError(err, "错误","保存配置文件config.ini出错")
}

//获取字符串配置项
func get(key string) *ini.Key{
	cfg := Load()
	path := strings.Split(key, ".")
	if len(path)==2{
		return cfg.Section(path[0]).Key(path[1])
	}else{
		return cfg.Section("").Key(path[0])
	}
}

//获取字符串配置项
func Get(key string) string{
	k := get(key)
	return k.Value()
}

//获取布尔配置项
func GetBool(key string) bool{
	k := get(key)
	b,_ := k.Bool()
	return b
}

//修改配置项
func Set(cfg *ini.File, key string, value string) *ini.File{
	path := strings.Split(key, ".")
	if len(path)==2{
		cfg.Section(path[0]).Key(path[1]).SetValue(value)
	}else{
		cfg.Section("").Key(path[0]).SetValue(value)
	}
	return cfg
}