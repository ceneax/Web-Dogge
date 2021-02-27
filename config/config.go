package config

import (
	"Dogge/util"
	"os"
	"path/filepath"
	"time"

	"github.com/go-ini/ini"
)

/*
 * 全局常量
 */
const (
	WorkPath = "./Dogge-Data" // 工作目录
)

/*
 * 全局变量
 * 读取配置文件获取值
 */
var (
	cfg *ini.File // config配置文件

	RunMode string // 运行模式：debug | release

	ServerHost string // 服务端运行监听地址
	ServerPort string    // 服务端运行监听端口

	LogReadTimeout  time.Duration // 读取超时时间
	LogWriteTimeout time.Duration // 写入超时时间
	LogFilePath     string        // 日志文件路径
	LogFileName     string        // 日志文件名
)

/*
 * 初始化
 */
func Init() {
	// 获取工作路径
	//workPath, err := os.Getwd()
	//if err != nil {
	//	// 工作路径获取失败
	//	panic(err)
	//}

	const fileName = "config.ini"
	// 生成config.ini配置文件的完整路径
	var filePath = filepath.Join(WorkPath, fileName)

	// 判断文件是否存在
	if !util.FileExists(filePath) {
		appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic(err)
		}

		filePath = filepath.Join(appPath, fileName)
		if !util.FileExists(filePath) {
			panic("初始化错误：配置文件config.ini不存在！")
		}
	}

	var err error
	cfg, err = ini.Load(filePath)
	if err != nil {
		panic("初始化错误：配置文件config.ini读取失败！原因：" + err.Error())
	}

	loadBase()
	loadServer()
}

/*
 * 加载基本参数
 */
func loadBase() {
	RunMode = cfg.Section("").Key("RUN_MODE").MustString("debug")
}

/*
 * 加载server节点参数
 */
func loadServer() {
	sec, err := cfg.GetSection("server")
	if err != nil {
		panic("初始化错误：配置文件的server节点数据读取失败！原因：" + err.Error())
	}

	ServerHost = sec.Key("SERVER_HOST").MustString("localhost")
	ServerPort = sec.Key("SERVER_PORT").MustString("44444")
	LogReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	LogWriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	LogFilePath = sec.Key("LOG_FILEPATH").MustString("log.txt")
}