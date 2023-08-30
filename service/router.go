package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"min/api"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	//api.StartCron()

	//model.Init()

	r := router.Group("/app/v1")

	{
		r.POST("/add/:number", api.AddExpense)
		r.POST("/delete", api.Delete)
		r.POST("/get/month", api.GetMonth)
		r.POST("/get/week", api.GetWeek)
		r.POST("get/day", api.GetDay)
		r.POST("/get/month/:count", api.MulGetMonth)
		r.POST("/get/week/:count", api.MulGetWeek)
		r.POST("get/day/:count", api.MulGetDay)
		r.POST("/get/times", api.GetTimes)
	}

	return router
}

func Run(port string) {

	router := NewRouter()
	// 创建一个 HTTP 服务器
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	// 启动 HTTP 服务器
	go func() {
		fmt.Println("启动 Gin 服务器...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Gin 服务器错误: %s\n", err)
		}
	}()

	// 监听操作系统的中断信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// 接收中断信号
	<-signalChan
	fmt.Println("接收到中断信号，开始优雅关闭...")

	// 创建一个上下文对象，用于控制关闭操作的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭 HTTP 服务器
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("关闭 HTTP 服务器错误: %s\n", err)
	}

	fmt.Println("优雅关闭完成")
}
