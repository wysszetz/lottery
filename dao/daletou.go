package dao

import (
	"fmt"
	"strconv"
)

type Dlt struct {
	Id       int    `db:"id"`
	Num      string `db:"num"`
	Num1     int    `db:"num_1"`
	Num2     int    `db:"num_2"`
	Num3     int    `db:"num_3"`
	Num4     int    `db:"num_4"`
	Num5     int    `db:"num_5"`
	Num6     int    `db:"num_6"`
	Num7     int    `db:"num_7"`
	OpenTime string `db:"open_time"`
}

func init() {
	err := InitDB()
	if err != nil {
		panic(err)
	}
}

func NewDlt(number, openTime string, num1, num2, num3, num4, num5, num6, num7 int) *Dlt {
	id, _ := strconv.Atoi(number)
	return &Dlt{
		Id:       id,
		Num:      number,
		Num1:     num1,
		Num2:     num2,
		Num3:     num3,
		Num4:     num4,
		Num5:     num5,
		Num6:     num6,
		Num7:     num7,
		OpenTime: openTime,
	}
}

func (d *Dlt) InsertDlt() {
	result, err := DB.Exec("Insert into dlt(id,num,num_1,num_2,num_3,num_4,num_5,num_6,num_7,open_time) values (?,?,?,?,?,?,?,?,?,?)", d.Id, d.Num, d.Num1, d.Num2, d.Num3, d.Num4, d.Num5, d.Num6, d.Num7, d.OpenTime)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := result.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

func GetDLTRowById(id int) Dlt {
	sqlStr := "select id,num,num_1,num_2,num_3,num_4,num_5,num_6,num_7,open_time from dlt where id=?"
	var dlt Dlt
	err := DB.Get(&dlt, sqlStr, id)
	if err != nil {
		fmt.Printf("get dlt row by id err,err:[%v]\n", err)
	}
	return dlt
}
