package router

import (
	"Dogge/config"
	"Dogge/http/api/system"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

/*
 * gin路由启动函数
 */
func Run() {
	if config.RunMode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}

	// 创建gin engine
	r := gin.New()

	// 全局中间件
	// 使用Logger中间件
	if config.RunMode == "release" {
		r.Use(gin.Logger())
	} else {
		r.Use(gin.Logger())
	}

	// 使用 Recovery 中间件
	r.Use(gin.Recovery())

	// 面板 路由
	go runPanel()

	// api 最顶层 路由组
	apiGroup := r.Group("/api")
	// 系统 路由组
	sysGroup := apiGroup.Group("/system")
	{
		sysGroup.GET("/login", system.Login)
	}
	// 应用程序 路由组
	appGroup := apiGroup.Group("/app")
	{
		appGroup.POST("")
	}

	r.Run(config.ServerHost + ":" + config.ServerPort)
}

func runPanel() {
	//staticHandler := http.FileServer(http.Dir(path.Join(config.WorkPath, "panel")))
	//
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	staticHandler.ServeHTTP(writer, request)
	//	//if request.URL.Path != "/" {
	//	//	staticHandler.ServeHTTP(writer, request)
	//	//	return
	//	//}
	//	//_, _ = io.WriteString(writer, "aa")
	//})

	http.Handle("/", http.FileServer(http.Dir(path.Join(config.WorkPath, "panel"))))

	err := http.ListenAndServe(":44445", nil)
	if err != nil {
		panic("静态服务器启动失败！错误信息：" + err.Error())
	}
}