package api

import (
	"bills/dao"
	"bills/model"
	"bufio"
	"encoding/csv"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func LoadWeiChatBills(c *gin.Context) {
	name := c.PostForm("name")
	file, _ := c.FormFile("file")
	_ = c.SaveUploadedFile(file, file.Filename)
	defer os.Remove(file.Filename)
	csvFile, err := os.Open(file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer csvFile.Close()

	db := dao.ConnPool.GetDb()

	scanner := bufio.NewScanner(csvFile)

	// Skip lines until the special string is found
	splitFlag := false
	csvContent := ""
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "微信支付账单明细列表") {
			splitFlag = true
			continue
		}
		if splitFlag {
			csvContent += scanner.Text() + "\n"
		}
	}

	// Now scanner is at the right position, create a csv reader
	reader := csv.NewReader(strings.NewReader(csvContent))
	reader.LazyQuotes = true
	records, err := reader.ReadAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Process the rest of the lines
	bills := make([]model.Bill, 0)
	for _, line := range records[1:] {
		transactionTime, _ := time.Parse("2006-01-02 15:04:05", line[0])
		goods := line[3]
		incomeOrExpenditure := line[4]
		amount, _ := strconv.ParseFloat(strings.Replace(line[5], "¥", "", -1), 64)
		paymentMethod := line[6]
		status := line[7]
		transactionNumber := line[8]
		notes, _ := json.Marshal(line)
		bill := model.Bill{
			ID:              transactionNumber,
			ChargeTime:      transactionTime,
			Value:           amount,
			Channel:         "微信",
			Goods:           goods,
			Tag:             "",
			TransactionType: incomeOrExpenditure,
			SubType:         paymentMethod,
			Status:          status,
			Name:            name,
			Notes:           string(notes),
		}
		bills = append(bills, bill)
	}
	err = db.Create(&bills).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "File processed"})
	return
}
