package api

import (
	"bills/dao"
	"bills/model"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"time"
)

func CreateOne(c *gin.Context) {
	var form struct {
		ID              string `json:"ID,omitempty"`
		ChargeTime      string `json:"ChargeTime,omitempty"`
		Value           string `json:"Value,omitempty"`
		Channel         string `json:"Channel,omitempty"`
		Goods           string `json:"Goods,omitempty"`
		Tag             string `json:"Tag,omitempty"`
		TransactionType string `json:"TransactionType,omitempty"`
		SubType         string `json:"SubType,omitempty"`
		Status          string `json:"Status,omitempty"`
		Name            string `json:"Name,omitempty"`
		Notes           string `json:"Notes,omitempty"`
	}
	db := dao.ConnPool.GetDb()
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t, _ := time.Parse("2006-01-02T15:04", form.ChargeTime)
	bill := model.Bill{
		ID:              form.ID,
		ChargeTime:      t,
		Value:           cast.ToFloat64(form.Value),
		Channel:         form.Channel,
		Goods:           form.Goods,
		Tag:             form.Tag,
		TransactionType: form.TransactionType,
		SubType:         form.SubType,
		Status:          form.Status,
		Name:            form.Name,
		Notes:           form.Notes,
	}
	if err := db.Create(&bill).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bill)
}
