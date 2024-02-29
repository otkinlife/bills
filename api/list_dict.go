package api

import (
	"bills/dao"
	"bills/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DictData struct {
	Channels         []string `json:"channel"`
	Tags             []string `json:"tag"`
	TransactionTypes []string `json:"transaction_type"`
	Statuses         []string `json:"status"`
	Names            []string `json:"charge_name"`
}

type ListDictResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data DictData `json:"data"`
}

func ListDict(c *gin.Context) {
	db := dao.ConnPool.GetDb()

	var channels []string
	var tags []string
	var transactionTypes []string
	var statuses []string
	var names []string

	db.Model(&model.Bill{}).Distinct().Where("channel <> ''").Pluck("channel", &channels)
	db.Model(&model.Bill{}).Distinct().Where("tag <> ''").Pluck("tag", &tags)
	db.Model(&model.Bill{}).Distinct().Where("transaction_type <> ''").Pluck("transaction_type", &transactionTypes)
	db.Model(&model.Bill{}).Distinct().Where("status <> ''").Pluck("status", &statuses)
	db.Model(&model.Bill{}).Distinct().Where("name <> ''").Pluck("name", &names)

	res := ListDictResponse{
		Code: 200,
		Msg:  "Success",
		Data: DictData{
			Channels:         channels,
			Tags:             tags,
			TransactionTypes: transactionTypes,
			Statuses:         statuses,
			Names:            names,
		},
	}

	c.JSON(http.StatusOK, res)
}
