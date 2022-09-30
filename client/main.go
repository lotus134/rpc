package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1. 客户端与微服务端建立连接

	conn, err1 := rpc.Dial("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println(err1)
	}
	//2、当客户端退出的时候关闭连接
	defer conn.Close()

	// 3. 调用远程函数
	var res string
	/*
		入参说明：
		1、第一个参数   helloWord.SayHelloWord      helloWord表示服务名称  SayHelloWord方法名称
		2、第二个参数  给服务端的req 传递数据
		3、第三个参数 需要传入地址  获取微服务端返回的数据
	*/
	err2 := conn.Call("helloWord.SayHelloWord", "小王", &res)
	if err2 != nil {
		fmt.Println(err2)
	}
	//4、获取微服务返回的数据
	fmt.Println(res)
}
