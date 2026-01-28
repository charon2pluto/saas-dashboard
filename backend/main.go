package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // 1. 配置 CORS (允许前端跨域访问)
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    })

    // 2. 核心 API：返回仪表盘数据
    r.GET("/api/dashboard", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "WPS Backend Online",
            "stats": gin.H{
                "users":    1204,   // 模拟用户数
                "revenue":  8999,   // 模拟收入
                "growth":   "+15%", // 模拟增长
            },
        })
    })

    // 3. 在 8080 端口启动
    r.Run(":8080") 
}