package tomlconfig

import (
	"godesk/component/message"
	"github.com/BurntSushi/toml"
	"bytes"
	"io/ioutil"
)

type Config struct {
	Server server
}

type server struct {
	Addr string
}

//加载配置文件
func Load() Config {
	var conf Config
	_, err := toml.DecodeFile("config/config.tml", &conf)
	message.CheckFatal(err, "错误","读取配置文件config.tml出错")
	return conf
}

//保存配置文件
func Save(conf Config)  {
	buf := new(bytes.Buffer)
	err := toml.NewEncoder(buf).Encode(conf)
	if(message.CheckError(err, "错误","编码配置文件失败")) {
		err := ioutil.WriteFile("config/config.tml",buf.Bytes(),0644)
		message.CheckError(err, "错误","保存配置文件config.tml失败")
	}
}