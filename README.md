```shell
# 设置环境变量
GOPROXY=https://goproxy.cn,direct

# 打包
GOOS=linux GOARCH=amd64 go build -o myGo main.go

# 运行
chmod +x myGo
nohup ./myGo > app.log 2>&1 &
```