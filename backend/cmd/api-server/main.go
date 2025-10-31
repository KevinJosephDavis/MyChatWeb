package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinjosephdavis/chatroom/internal/controller"
	"github.com/kevinjosephdavis/chatroom/internal/repository"
	"github.com/kevinjosephdavis/chatroom/internal/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func startHTTPServer() {
	fmt.Println("HTTP服务器在8080端口监听...")

	// 1.初始化数据库
	dsn := "root:123456@tcp(localhost:3306)/chatroom?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败，错误信息：", err)
		return
	}

	// 2.自动创建表
	//db.AutoMigrate() 暂时注释，设计好所有表结构再打开

	// 3.初始化各层组件
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// 4.创建Gin引擎
	r := gin.Default()

	// 5.注册路由
	api := r.Group("/api")
	{
		api.POST("/register", userController.Register)
		api.POST("/login", userController.Login)
	}

	//6.添加测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong from gin!",
		})
	})

	//启动HTTP服务器
	if err := r.Run(":8080"); err != nil {
		fmt.Println("HTTP服务器启动失败，错误信息：", err)
	}

}

func main() {
	startHTTPServer()
}
