# GithubIps

**由于天朝问题，有时候某个IP被墙了，就会出现图片显示不全或者显示不了的问题**

本项目是用Golang写的简单程序，用于获取GitHub的最新IP，解决无法显示GitHub的图片  

主要使用第三方包[soup](https://github.com/anaskhan96/soup).

我已经编译好了Mac、Linux和Windows三个平台的可执行文件  
大家可以按自己的平台自己下载或自行编译。  
* `Mac` [🔗](https://github.com/Sanbolee/GithubIps/releases/download/v1.0/GithubIps_mac)
* `Linux` [🔗](https://github.com/Sanbolee/GithubIps/releases/download/v1.0/GithubIps_linux)
* `Windows` [🔗](https://github.com/Sanbolee/GithubIps/releases/download/v1.0/GithubIps_windows.exe)

## 运行  
Mac
```bash
chmod +x GithubIps_mac
./GithubIps_mac
```
Linux
```bash
chmod +x GithubIps_linux
./GithubIps_linux
```
Windows
```bash
双击666即可
```

## 结果
```bash
正在输出结果，请稍等：
185.199.108.153   assets-cdn.github.com
199.232.68.133   avatars2.githubusercontent.com
199.232.68.133   avatars1.githubusercontent.com
140.82.112.4   gist.github.com
140.82.112.4   github.com
199.232.68.133   avatars4.githubusercontent.com
199.232.68.133   avatars3.githubusercontent.com
199.232.68.133   avatars7.githubusercontent.com
199.232.68.133   avatars0.githubusercontent.com
199.232.68.133   avatars6.githubusercontent.com
199.232.68.133   camo.githubusercontent.com
199.232.68.133   avatars5.githubusercontent.com
199.232.68.133   avatars8.githubusercontent.com
199.232.68.133   raw.githubusercontent.com
199.232.68.133   gist.githubusercontent.com
199.232.68.133   cloud.githubusercontent.com
输出结果完毕，请复制粘贴到hosts文件。
Mac系统和Linux系统hosts文件位置：/etc/hosts
Win系统hosts文件位置：c:/windows/system32/drivers/etc/hosts
```

## 编译 
Mac 下压缩编译 Linux 和 Windows 64位可执行程序
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" main.go
```
Linux 下压缩编译 Mac 和 Windows 64位可执行程序
```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s" main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" main.go
```
Windows 下压缩编译 Mac 和 Linux 64位可执行程序
```bash
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -ldflags "-w -s" main.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-w -s" main.go
```