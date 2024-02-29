package main

import (
	"bills/api"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	// 静态文件
	r.Static("/static", "./static")

	r.GET("/ping", api.Ping)
	r.POST("/load_weichat_bills", api.LoadWeiChatBills)
	r.POST("/load_alipay_bills", api.LoadAliPayBills)
	r.POST("/list", api.List)
	r.GET("/list_dict", api.ListDict)
	r.POST("/create_one", api.CreateOne)
	r.POST("/update_one", api.UpdateOne)
	r.POST("/delete_one", api.DeleteOne)
}
