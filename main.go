package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	Router(r)
	_ = r.Run("0.0.0.0:8228")
}
