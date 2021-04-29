package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)


func main() {

	g, ctx := errgroup.WithContext(context.Background())
	fmt.Println("服务启动start！")
	addr := ":9091"
	s :=&http.Server{
		Addr: addr,
		Handler:http.DefaultServeMux,
	}
	g.Go(func() error {
		http.HandleFunc("/test1", func(writer http.ResponseWriter, request *http.Request) {
			fmt.Println("tes1")
			writer.Write([]byte("tes1"))
		})
		return s.ListenAndServe()
	})
	g.Go(func() error {
		exit := make(chan os.Signal)
		//监听 Ctrl+C 信号
		signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-exit:
			fmt.Println("进程已被取消~")
			return s.Shutdown(ctx)
		}
	})
	err := g.Wait()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("服务启动成功！")
	if ctx.Err() !=nil {
		fmt.Println(ctx.Err())
		fmt.Println("服务关闭成功！")
		os.Exit(0)
	}

}
