package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// 设置文件服务器的根目录
	rootDir := "E:/Test/" // 替换为你的本地文件路径

	// 创建一个文件服务器实例
	fs := http.FileServer(http.Dir(rootDir))

	// 定义一个处理函数，用于处理请求并将请求重定向到文件服务器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 获取请求的文件路径
		requestedFile := filepath.Join(rootDir, r.URL.Path)
		fmt.Println(requestedFile)
		// 检查请求的文件是否存在
		if _, err := os.Stat(requestedFile); err == nil {
			// 请求的文件存在，将请求重定向到文件服务器
			fs.ServeHTTP(w, r)
		} else {
			// 请求的文件不存在，返回404错误
			http.NotFound(w, r)
		}
	})

	// 启动HTTP服务器并监听指定端口
	port := "8080" // 替换为你希望使用的端口号
	fmt.Printf("Server is running on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
