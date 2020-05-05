package main

import (
	"flag"
	"fmt"
	"lottery/colly"
	"lottery/dlt"
)

func main() {
	var (
		sync    int
		maxPage int64
		i       int64
	)
	flag.IntVar(&sync, "sync", 1, "同步类型，1：同步最新一页，2:全量同步")
	flag.Parse()

	switch sync {
	case 1:
		maxPage = int64(1)
	case 2:
		maxPage = colly.GetDLTMaxPage()
	default:
		fmt.Println("请输入正确的同步类型，1：同步最新一页，2:全量同步")
		return
	}

	for i = 1; i <= maxPage; i++ {
		pageUri := colly.GetDLTPageUrl(i)
		colly.DltColly(pageUri)
	}
	dlt.CalcParity()

}
