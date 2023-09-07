package router

import (
	"frontend-websocket-test-framework/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 自己往service包里面加接口

func Router(r *gin.Engine) {

	// 跨域
	r.Use(func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	})

	r.GET("/device", service.GetRouterList())
	r.GET("/usableDevice", service.GetUsableList())
}
