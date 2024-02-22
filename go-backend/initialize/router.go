package initialize

import (
	"context"
	"errors"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go-backend/global"
	"go-backend/middleware"
	"go-backend/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func initRouters() *gin.Engine {
	router := gin.Default()
	/*  TODO 前端项目静态资源
	router.StaticFile("/", "./static/dist/index.html")
	router.Static("/assets", "./static/dist/assets")
	router.StaticFile("/favicon.ico", "./static/dist/favicon.ico")
	// 其他静态资源
	router.Static("/public", "./static")
	router.Static("/storage", "./storage/app/public")*/

	// 路由分组
	PublicGroup := router.Group("/")
	router.Use(middleware.GinLogger())
	//routers.SetupUserRouter(PublicGroup)
	//routers.SetupBaseRouter(PublicGroup)
	routers.SetApiGroupRoutes(PublicGroup)
	return router
}

func RunServer() {
	r := initRouters()
	defer CloseDB()
	srv := &http.Server{
		Addr:    ":" + global.GlobConfig.App.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	color.Blue("Server exited")

}
