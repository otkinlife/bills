package api

import (
	"bills/dao"
	"bills/model"
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func LoadAliPayBills(c *gin.Context) {
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
		line := scanner.Text()
		if !utf8.Valid([]byte(line)) {
			// Create a transformer to convert GBK to UTF-8
			transformer := simplifiedchinese.GBK.NewDecoder()
			// Use the transformer to convert the text
			utf8Text, err := io.ReadAll(transform.NewReader(strings.NewReader(line), transformer))
			if err != nil {
				fmt.Println("Error converting text:", err)
			} else {
				line = string(utf8Text)
			}
		}
		if strings.Contains(line, "支付宝（中国）网络技术有限公司  电子客户回单") {
			splitFlag = true
			continue
		}
		if splitFlag {
			csvContent += line + "\n"
		}
	}

	// Now scanner is at the right position, create a csv reader
	reader := csv.NewReader(strings.NewReader(csvContent))
	records, err := reader.ReadAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Process the rest of the lines
	bills := make([]model.Bill, 0)
	for _, line := range records[1:] {
		transactionTime, _ := time.Parse("2006-01-02 15:04:05", line[0])
		transactionType := line[1]
		goods := line[4]
		incomeOrExpenditure := line[5]
		amount, _ := strconv.ParseFloat(strings.Replace(line[6], "¥", "", -1), 64)
		paymentMethod := line[7]
		status := line[8]
		transactionNumber := line[9]
		notes, _ := json.Marshal(line)
		bill := model.Bill{
			ID:              transactionNumber,
			ChargeTime:      transactionTime,
			Value:           amount,
			Channel:         "支付宝",
			Goods:           goods,
			Tag:             transactionType,
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
