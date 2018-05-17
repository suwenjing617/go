package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Tickers []struct {
	// https://mholt.github.io/json-to-go/

	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	Two4HVolumeUsd   string `json:"24h_volume_usd"`
	MarketCapUsd     string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	MaxSupply        string `json:"max_supply"`
	PercentChange1H  string `json:"percent_change_1h"`
	PercentChange24H string `json:"percent_change_24h"`
	PercentChange7D  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
}

var rank int
var price string
var name string

func main() {
	for true {
		fetch()
	}
}

func fetch() {
	//连接数据库
	db, _ := sql.Open("mysql", "root:123456@/gopar?charset=utf8")

	url := "https://api.coinmarketcap.com/v1/ticker/?start=0&limit=30"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	// fmt.Println(data)
	var data_total Tickers

	err = json.Unmarshal([]byte(data), &data_total)

	for _, v := range data_total {
		// fmt.Printf("第 %d 位 x 的值 = %s\n", i, x.Name)
		name := v.Name
		price := v.PriceUsd
		rank := v.Rank
		//插入

		// stmt, err := db.Prepare("INSERT priceUsd SET name=?,price=?")
		// checkErr(err)

		// res, err := stmt.Exec(name, price)
		// checkErr(err)

		// id, err := res.LastInsertId()
		// checkErr(err)

		// fmt.Println(id)

		//修改
		stmt, err := db.Prepare("update priceUsd set price=?,name=? where id=?")
		checkErr(err)

		res, err := stmt.Exec(price, name, rank)
		checkErr(err)

		affect, err := res.RowsAffected()
		checkErr(err)

		fmt.Println(affect)
		resp, err = http.Get(url)
		fmt.Println(resp)
	}
	time.Sleep(time.Second * 10)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
