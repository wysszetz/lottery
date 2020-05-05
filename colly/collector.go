package colly

import (
	"github.com/gocolly/colly"
)

var Colly *colly.Collector

func InitColly() {
	Colly = colly.NewCollector(
		//这次在colly.NewCollector里面加了一项colly.Async(true)，表示抓取时异步的
		colly.Async(true),
		//模拟浏览器
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)
	return
}
