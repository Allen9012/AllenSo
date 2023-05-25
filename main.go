package main

import (
	"AllenSo/config/config_init"
	"AllenSo/docs"
	"AllenSo/job/once"
	"AllenSo/router"
	"fmt"
)

func main() {
	// 从配置文件读取配置
	err := config_init.Init()
	if err != nil {
		fmt.Printf("config_init config error : %s \n", err)
		return
	}
	//初始化 post 的数据 只在开始的时候使用一次，注释掉下次就不会再执行了
	if err := once.FetchInitPostList(); err != nil {
		fmt.Printf("fetch post data spider error  %s \n", err)
		return
	}
	// gin-swagger middleware
	docs.SwaggerInfo.BasePath = "/api/v1"
	// 装载路由
	r := router.NewRouter()
	r.Run(":3000")
}
