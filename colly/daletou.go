package colly

import (
	"fmt"
	"github.com/gocolly/colly"
	"lottery/dao"
	"lottery/logger"
	strconv "strconv"
	"strings"
)

type LotteryResult struct {
	Number       string
	FirstNumber  int
	SecondNumber int
	ThirdNumber  int
	FourthNumber int
	FifthNumber  int
	SixNumber    int
	SevenNumber  int
	OpenTime     string
}

var url = "https://www.lottery.gov.cn/historykj/history_%d.jspx?_ltype=dlt"

func init() {
	InitColly()
}

func GetDLTPageUrl(page int64) string {
	return fmt.Sprintf(url, page)
}

func GetDLTMaxPage() int64 {
	var page int64
	//限制采集规则
	Colly.Limit(&colly.LimitRule{DomainGlob: "*.lottery.gov.cn.*", Parallelism: 5})
	Colly.OnRequest(func(r *colly.Request) {
		logger.NewFileLogger("info").Info("正在请求访问：%s\n", r.URL)
	})
	Colly.OnError(func(_ *colly.Response, err error) {
		logger.NewFileLogger("error").Error("访问请求异常：%v\n", err)
	})
	Colly.OnResponse(func(r *colly.Response) {
		logger.NewFileLogger("info").Info("已经完成访问：%s\n", r.Request.URL)
	})

	Colly.OnHTML("#selectobj", func(e *colly.HTMLElement) {
		pageStr := strings.Replace(e.Text, " ", "", -1)
		pages := strings.Split(pageStr, "\n")
		pages = append(pages[:0], pages[1:]...)
		pages = append(pages[:len(pages)-1])
		getLastPage := pages[len(pages)-1]
		page, _ = strconv.ParseInt(getLastPage, 10, 64)
	})

	Colly.OnScraped(func(r *colly.Response) {

	})
	Colly.Visit(GetDLTPageUrl(1))
	Colly.Wait()
	return page
}

func DltColly(url string) {
	//限制采集规则
	Colly.Limit(&colly.LimitRule{DomainGlob: "*.lottery.gov.cn.*", Parallelism: 5})
	Colly.OnRequest(func(r *colly.Request) {
		logger.NewFileLogger("info").Info("正在请求访问：%s\n", r.URL)
	})
	Colly.OnError(func(_ *colly.Response, err error) {
		logger.NewFileLogger("error").Error("访问请求异常：%v\n", err)
	})
	Colly.OnResponse(func(r *colly.Response) {
		logger.NewFileLogger("info").Info("已经完成访问：%s\n", r.Request.URL)
	})

	lotteryRes := make([]*LotteryResult, 0)
	Colly.OnHTML("body > div.yyl > div.yylMain > div.result > table > tbody > tr", func(e *colly.HTMLElement) {
		lottery := LotteryResult{
			Number:   e.ChildText("td:first-child"),
			OpenTime: e.ChildText("td:nth-child(20)"),
		}
		firstNumber, err := strconv.Atoi(e.ChildText("td:nth-child(2)"))
		if err == nil {
			lottery.FirstNumber = firstNumber
		}

		secondNumber, err := strconv.Atoi(e.ChildText("td:nth-child(3)"))
		if err == nil {
			lottery.SecondNumber = secondNumber
		}

		thirdNumber, err := strconv.Atoi(e.ChildText("td:nth-child(4)"))
		if err == nil {
			lottery.ThirdNumber = thirdNumber
		}

		fourthNumber, err := strconv.Atoi(e.ChildText("td:nth-child(5)"))
		if err == nil {
			lottery.FourthNumber = fourthNumber
		}
		fifthNumber, err := strconv.Atoi(e.ChildText("td:nth-child(6)"))
		if err == nil {
			lottery.FifthNumber = fifthNumber
		}

		sixNumber, err := strconv.Atoi(e.ChildText("td:nth-child(7)"))
		if err == nil {
			lottery.SixNumber = sixNumber
		}

		sevenNumber, err := strconv.Atoi(e.ChildText("td:nth-child(8)"))
		if err == nil {
			lottery.SevenNumber = sevenNumber
		}

		lotteryRes = append(lotteryRes, &lottery)
	})

	Colly.OnScraped(func(r *colly.Response) {
		for _, val := range lotteryRes {
			id, _ := strconv.Atoi(val.Number)
			check := dao.GetDLTRowById(id)
			if check.Id == 0 {
				dlt := dao.NewDlt(val.Number, val.OpenTime, val.FirstNumber, val.SecondNumber, val.ThirdNumber, val.FourthNumber, val.FifthNumber, val.SixNumber, val.SevenNumber)
				dlt.InsertDlt()
			}
			break
		}
	})
	Colly.Visit(url)
	Colly.Wait()
}
