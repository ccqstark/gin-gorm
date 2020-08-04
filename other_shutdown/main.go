package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r:=gin.Default()
	r.GET("/test",func(c *gin.Context){
		time.Sleep(10 * time.Second)
		c.String(200,"hello test")
	})
	//创建服务器
	srv:=&http.Server{
		Addr:":8085",
		Handler:r,
	}

	//启动服务器进行监听
	go func() {
		if err:=srv.ListenAndServe();err!=nil && err!=http.ErrServerClosed{
			log.Fatalf("listen: %s\n",err)
		}
	}()

	quit:=make(chan os.Signal)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit
	log.Println("shutdown Server...")

	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	if err:=srv.Shutdown(ctx);err!=nil {
		log.Fatal("server shutdown:",err)
	}
	log.Println("server exiting")
}
