package api

import (
	"bills/dao"
	"bills/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteOne(c *gin.Context) {
	var form struct {
		ID *string `json:"ID,omitempty"`
	}
	db := dao.ConnPool.GetDb()
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Delete(&model.Bill{}, form.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}
