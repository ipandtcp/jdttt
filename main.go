package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://ggog.6.owaap.com")
	if err != nil {
		fmt.Println("请求出错:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("状态码:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	fmt.Println("响应内容:")
	fmt.Println(string(body))
}
