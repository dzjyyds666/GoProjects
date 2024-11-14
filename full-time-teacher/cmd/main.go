package main

import (
	"FullTimeTeacher/config"
	"FullTimeTeacher/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//读取配置文件
	config, err := config.LoggingConfig()
	if err != nil {
		fmt.Printf("读取配置文件失败: %v", err)
		return
	}

	// 数据库初始化
	err = models.MySQLConnection(config)
	if err != nil {
		fmt.Printf("数据库连接失败: %v", err)
		return
	}

	r.Run(":9999")
}
