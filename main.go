package main

import (
	"Dogge/config"
	"Dogge/http/router"
	"fmt"
)

func main()  {
	fmt.Println("Dogge")

	// 初始化配置文件
	config.Init()
	// 启动路由
	router.Run()
}