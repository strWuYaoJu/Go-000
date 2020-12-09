/*
作业：
	基于 errgroup 实现一个 http server 的启动和关闭 ，
	以及 linux signal 信号的注册和处理，要保证能够一个退出，
	全部注销退出
*/


思路：
    新建一个errgroup返回它的一个指针g和errCtx，循环服务列表，通过g.Go()分别启动对应的服务newServer，newServer里面实现启动服务，同时启动一个goroutine读取是否接收到cancel的信号，如果收到取消信号，则Shutdown服务。
    在主goroutine中，通过g.Wait()等待服务出现错误，如果有任何一个服务出错，则调用cancel()发出取消信号，退出所有服务。