package main

import (
	"AllenSo/config/config_init"
	"AllenSo/docs"
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
	// gin-swagger middleware
	docs.SwaggerInfo.BasePath = "/api/v1"
	// 装载路由
	r := router.NewRouter()
	r.Run(":3000")
}
