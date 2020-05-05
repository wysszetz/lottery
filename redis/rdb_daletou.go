package redis

import (
	"lottery/colly"
)

const ZsetDaletouKey = "daletou"

func AddResultZset(dlt *colly.LotteryResult) {
	// ZADD
	//num, err := Rdb.ZAdd(ZsetDaletouKey, dlt...).Result()
	//if err != nil {
	//	fmt.Printf("zadd failed, err:%v\n", err)
	//	return
	//}
	//fmt.Printf("zadd %d succ.\n", num)
}
