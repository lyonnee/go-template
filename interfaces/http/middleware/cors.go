package middleware

import (
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/cors"
)

func CORS() app.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "DELETE", "PUT", "GET", "OPTIONS", "PATCH", "UPDATE", "HEAD"},
		AllowHeaders: []string{
			"Authorization",     // 标准认证头部 (Bearer Token)
			"Content-Length",    // 标准请求体长度
			"X-CSRF-Token",      // 保留 (CSRF 防护，虽非标准但广泛使用)
			"Accept",            // 标准内容类型偏好
			"Origin",            // 标准跨域请求源
			"Host",              // 标准请求主机
			"Connection",        // 标准连接控制
			"Accept-Encoding",   // 标准编码偏好
			"Accept-Language",   // 标准语言偏好
			"User-Agent",        // 标准用户代理
			"If-Modified-Since", // 标准缓存验证
			"Cache-Control",     // 标准缓存控制
			"Content-Type",      // 标准内容类型
			"Pragma",            // 标准缓存控制 (HTTP/1.0 兼容)
			"Cookie",            // 标准会话管理 (替代 session)
			"Set-Cookie",        // 标准会话管理 (替代 session)
			"Sec-Fetch-Dest",    // 标准请求目的地 (替代 X-Requested-With)
			"Sec-Fetch-Mode",    // 标准请求模式 (替代 X-Requested-With)
			"Sec-Fetch-Site",    // 标准请求来源站点
			"DNT",               // 标准 Do Not Track (隐私)
		},
		ExposeHeaders: []string{
			"Content-Length",               // 响应体字节长度
			"Access-Control-Allow-Origin",  // CORS：允许访问的域名（如 "https://example.com"）
			"Access-Control-Allow-Headers", // CORS：允许的请求头部（如 "Content-Type, Authorization"）
			"Access-Control-Allow-Methods", // CORS：允许的 HTTP 方法（如 "GET, POST"）
			"Cache-Control",                // 缓存策略（如 "no-cache", "max-age=3600"）
			"Content-Type",                 // 内容类型（如 "application/json"）
			"Last-Modified",                // 资源最后修改时间（用于缓存验证）
		},
		MaxAge:           12 * time.Hour,
		AllowCredentials: false,
	})
}
