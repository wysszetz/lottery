package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gopkg.in/ini.v1"
	"lottery/conf"
	"lottery/logger"
)

var (
	DB *sqlx.DB
)

var config = new(conf.AppConfig)

func InitDB() (err error) {
	err = ini.MapTo(config, "./conf/config.ini")
	if err != nil {
		logger.NewFileLogger("error").Error("解析配置异常，err：【%v】\n", err)
		return
	}
	//数据库信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", config.MysqlConfig.UserName, config.MysqlConfig.Password, config.MysqlConfig.IpAddress, config.MysqlConfig.Port, config.MysqlConfig.DbName, config.MysqlConfig.Charset)
	//链接
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		logger.NewFileLogger("error").Error("mysql connect failed, detail is [%v]\n", err.Error())
		return
	}
	err = DB.Ping()
	if err != nil {
		logger.NewFileLogger("error").Error("mysql ping failed, detail is [%v]\n", err)
	}
	return
}
