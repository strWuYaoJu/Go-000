package main

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

/*
作业：
	基于 errgroup 实现一个 http server 的启动和关闭 ，
	以及 linux signal 信号的注册和处理，要保证能够一个退出，
	全部注销退出
*/

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	g, errCtx := errgroup.WithContext(ctx)
	var serverAddrs = []string{
		"127.0.0.1:10001",
		"127.0.0.1:10002",
		"127.0.0.1:10003",
	}
	for _, v := range serverAddrs {
		addr := v
		g.Go(func() error {
			return newServer(errCtx, addr)
		})
	}

	if err := g.Wait(); err != nil {
		cancel() //退出所有服务
		fmt.Printf("err:%v \n", err)
		return
	}
	fmt.Println("ok !")
}

type serverObject struct {
}

//ServeHTTP 实现Handler接口中的ServeHTTP方法
func (s *serverObject) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	switch r.URL.Path {
	case "/server1":
		fmt.Fprintln(w, "hello server1")
	case "/server2":
		fmt.Fprintln(w, "hello server2")
	case "/server3":
		fmt.Fprintln(w, "hello server3")
	default:
		fmt.Fprintln(w, "hello server1")
	}
}

// server 新建服务
func newServer(ctx context.Context, addr string) error {

	s := http.Server{
		Addr:    addr,
		Handler: &serverObject{},
	}
	ctx1 := ctx
	go func() {
		<-ctx1.Done() //收到cancel信号，Shutdown服务
		s.Shutdown(context.Background())
		fmt.Printf("%s Shutdown!\n", addr)
	}()

	fmt.Printf("%s start!\n", addr)
	return s.ListenAndServe()
}
