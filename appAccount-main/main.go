package main

import (
	"log"
	"net/http"

	"app/api"
)

// 全局 CORS 中间件（允许所有来源）
func withCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 允许所有来源
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// 允许的方法
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 允许的请求头
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// 处理预检请求
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent) // 204
			return
		}

		// 正常请求继续
		h.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	// API 路由
	mux.Handle(
		"/api/shared-accounts",
		withCORS(http.HandlerFunc(api.SharedAccountHandler)),
	)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
