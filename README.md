# 安装

- 从https://sciter.com/download/下载sciter-sdk.zip，解压，找到sciter-sdk\bin\64\sciter.dll复制到c:\windows\system32

- 下载安装tdm-gcc：http://tdm-gcc.tdragon.net/download，安装后将tdm-gcc\bin加入到环境变量中

- cmd进入gopath目录src下运行：go get -x github.com/sciter-sdk/go-sciter

# 编译

windres -o main-res.syso main.rc && go build -i -ldflags="-H windowsgui" -o dist/godesk.exe