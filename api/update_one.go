package api

import (
	"bills/dao"
	"bills/model"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

func UpdateOne(c *gin.Context) {
	var form struct {
		ID              *string `json:"ID,omitempty"`
		ChargeTime      *string `json:"ChargeTime,omitempty"`
		Value           *string `json:"Value,omitempty"`
		Channel         *string `json:"Channel,omitempty"`
		Goods           *string `json:"Goods,omitempty"`
		Tag             *string `json:"Tag,omitempty"`
		TransactionType *string `json:"TransactionType,omitempty"`
		SubType         *string `json:"SubType,omitempty"`
		Status          *string `json:"Status,omitempty"`
		Name            *string `json:"Name,omitempty"`
		Notes           *string `json:"Notes,omitempty"`
	}
	db := dao.ConnPool.GetDb()
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var bill *model.Bill
	db.First(&bill, form.ID)
	if bill == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "bill not found"})
		return
	}
	updates := make(map[string]interface{})
	if form.ID != nil {
		updates["ID"] = *form.ID
	}
	//if form.ChargeTime != nil && *form.ChargeTime != "" && *form.ChargeTime != bill.ChargeTime.Format("2006-01-02T15:04") {
	//	t, _ := time.Parse("2006-01-02T15:04", *form.ChargeTime)
	//	updates["ChargeTime"] = t
	//}
	if form.Value != nil && *form.Value != "" && cast.ToFloat64(*form.Value) != bill.Value {
		updates["Value"] = cast.ToFloat64(*form.Value)
	}
	if form.Channel != nil && *form.Channel != "" && *form.Channel != bill.Channel {
		updates["Channel"] = *form.Channel
	}
	if form.Goods != nil && *form.Goods != "" && *form.Goods != bill.Goods {
		updates["Goods"] = *form.Goods
	}
	if form.Tag != nil && *form.Tag != "" && *form.Tag != bill.Tag {
		updates["Tag"] = *form.Tag
	}
	if form.TransactionType != nil && *form.TransactionType != "" && *form.TransactionType != bill.TransactionType {
		updates["TransactionType"] = *form.TransactionType
	}
	if form.SubType != nil && *form.SubType != "" && *form.SubType != bill.SubType {
		updates["SubType"] = *form.SubType
	}
	if form.Status != nil && *form.Status != "" && *form.Status != bill.Status {
		updates["Status"] = *form.Status
	}
	if form.Name != nil && *form.Name != "" && *form.Name != bill.Name {
		updates["Name"] = *form.Name
	}
	if form.Notes != nil && *form.Notes != "" && *form.Notes != bill.Notes {
		updates["Notes"] = *form.Notes
	}

	if err := db.Model(&model.Bill{}).Where("id = ?", *form.ID).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bill)
}
