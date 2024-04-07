package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Company struct {
	Name string
	Path string
}

func main() {
	http.FileServer(http.Dir("./")) // 替换为你的静态文件路径

	// 处理根路径请求，打开 index.html 文件
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html") // 替换为你的 index.html 文件路径
	})

	//// 将文件服务器处理器注册到路由
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("e:/"))))
	//http.Handle("/image/", http.StripPrefix("/web/", http.FileServer(http.Dir("./images"))))
	//http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))

	// http.HandleFunc("/QueCompany", QueCompanyHandle)
	http.HandleFunc("/QueCompany", QueCompanyHandle2)

	// 启动HTTP服务器并监听指定端口
	port := "8080" // 替换为你希望使用的端口号
	fmt.Printf("Server is running on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}

// 返回所有公司名称，通过json
func QueCompanyHandle(w http.ResponseWriter, r *http.Request) {
	// 设置响应头
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	// 如果是预检请求（OPTIONS），则直接返回成功状态码
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	company := Company{
		Name: "测试",
		Path: "测试",
	}
	jsonData, err := json.Marshal(company)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Println(jsonData)
	// 将 JSON 数据写入响应体
	w.Write(jsonData)
}

// 返回所有公司名称，通过json
func QueCompanyHandle2(w http.ResponseWriter, r *http.Request) {
	// 设置响应头
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	// 如果是预检请求（OPTIONS），则直接返回成功状态码
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	data, err := os.ReadFile("./test.json") // 替换为你的 JSON 文件路径
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/json")

	// 返回 JSON 数据
	w.Write(data)

}
