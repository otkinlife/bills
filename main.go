package main

import (
	"bills/middlewares"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.Cors())
	Router(r)
	_ = r.Run("0.0.0.0:8228")
}
