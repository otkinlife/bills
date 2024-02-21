package main

import (
	"bills/api"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.GET("/ping", api.Ping)
	r.POST("/load_weichat_bills", api.LoadWeiChatBills)
	r.POST("/load_alipay_bills", api.LoadAliPayBills)
}
