package api

import (
	"bills/dao"
	"bills/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Request struct {
	Page            int       `json:"page"`
	PageSize        int       `json:"page_size"`
	DateRange       []string  `json:"date_range"`  // Assume it's in the format of ["2024-01-01", "2024-12-31"]
	ValueRange      []float64 `json:"value_range"` // Assume it's in the format of [0.0, 100.0]
	Channel         string    `json:"channel"`
	Tag             string    `json:"tag"`
	TransactionType string    `json:"transaction_type"`
	Status          string    `json:"status"`
	Name            string    `json:"name"`
}

type Response struct {
	Code     int          `json:"code"`
	Msg      string       `json:"msg"`
	List     []model.Bill `json:"list"`
	Total    int64        `json:"total"`
	Page     int          `json:"page"`
	PageSize int          `json:"page_size"`
}

func List(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := dao.ConnPool.GetDb()

	var bills []model.Bill
	var total int64

	tx := db.Model(&model.Bill{})
	if len(req.DateRange) == 2 {
		start, _ := time.Parse("2006-01-02", req.DateRange[0])
		end, _ := time.Parse("2006-01-02", req.DateRange[1])
		tx = tx.Where("charge_time BETWEEN ? AND ?", start, end)
	}

	if len(req.ValueRange) == 2 {
		tx = tx.Where("value BETWEEN ? AND ?", req.ValueRange[0], req.ValueRange[1])
	}

	if req.Channel != "" {
		tx = tx.Where("channel = ?", req.Channel)
	}

	if req.Tag != "" {
		tx = tx.Where("tag = ?", req.Tag)
	}

	if req.TransactionType != "" {
		tx = tx.Where("transaction_type = ?", req.TransactionType)
	}

	if req.Status != "" {
		tx = tx.Where("status = ?", req.Status)
	}

	if req.Name != "" {
		tx = tx.Where("name = ?", req.Name)
	}
	tx.Count(&total)
	err := tx.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).Order("charge_time desc").Find(&bills).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := Response{
		Code:     200,
		Msg:      "Success",
		List:     bills,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	c.JSON(http.StatusOK, res)
}
