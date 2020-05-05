package dlt

import (
	"encoding/csv"
	"fmt"
	"lottery/dao"
	"math"
	"os"
	"strconv"
)

const PAGE_SIZE = 500.00

func CalcParity() {
	total := dao.GetDLTCount()
	page := math.Ceil(float64(total) / PAGE_SIZE)

	data := [][]string{{"期号", "开奖日期", "前区1", "前区2", "前区3", "前区4", "前区5", "前区奇数", "前区偶数", "后区1", "后区2", "后区奇数", "后区偶数"}}
	for i := 0; i < int(page); i++ {
		rows := dao.GetDLTRows(i*int(PAGE_SIZE), int(PAGE_SIZE))
		for _, val := range rows {
			var odd, even int
			if val.Num1%2 == 0 {
				even++
			} else {
				odd++
			}
			if val.Num2%2 == 0 {
				even++
			} else {
				odd++
			}
			if val.Num3%2 == 0 {
				even++
			} else {
				odd++
			}
			if val.Num4%2 == 0 {
				even++
			} else {
				odd++
			}
			if val.Num5%2 == 0 {
				even++
			} else {
				odd++
			}

			var after_odd, after_even int
			if val.Num6%2 == 0 {
				after_even++
			} else {
				after_odd++
			}

			if val.Num7%2 == 0 {
				after_even++
			} else {
				after_odd++
			}

			tmp := []string{
				val.Num,
				val.OpenTime,
				strconv.Itoa(val.Num1),
				strconv.Itoa(val.Num2),
				strconv.Itoa(val.Num3),
				strconv.Itoa(val.Num4),
				strconv.Itoa(val.Num5),
				strconv.Itoa(odd),
				strconv.Itoa(even),
				strconv.Itoa(val.Num6),
				strconv.Itoa(val.Num7),
				strconv.Itoa(after_odd),
				strconv.Itoa(after_even),
			}
			data = append(data, tmp)
		}
	}

	fullFile := "./runtime/CalcParity.csv"
	fileObj, err := os.OpenFile(fullFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
	}
	defer fileObj.Close()

	fileObj.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(fileObj)         //创建一个新的写入文件流
	w.WriteAll(data)
	w.Flush()
}
