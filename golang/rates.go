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

type AutoGenerated struct {
	Success   bool   `json:"success"`
	Terms     string `json:"terms"`
	Privacy   string `json:"privacy"`
	Timestamp int    `json:"timestamp"`
	Source    string `json:"source"`
	Quotes    struct {
		USDCNY float64 `json:"USDCNY"`
	} `json:"quotes"`
}

func main() {

	for true {
		fetchs()
	}

}

func fetchs() {

	//连接数据库
	db, _ := sql.Open("mysql", "root:123456@/gopar?charset=utf8")

	url := "http://apilayer.net/api/live?access_key=1e341d38f525981c1ab038d2e45810ca&currencies=CNY"

	resp, err := http.Get(url)

	fmt.Println(resp)

	if err != nil {
		fmt.Println(err)
	}

	data, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	var ticker AutoGenerated

	err = json.Unmarshal([]byte(data), &ticker)

	// fmt.Println(ticker.Quotes.USDCNY)

	rates := ticker.Quotes.USDCNY

	//插入
	// stmt, err := db.Prepare("INSERT rate SET rates=?")
	// checkErrs(err)

	// res, err := stmt.Exec(rates)
	// checkErrs(err)

	// id, err := res.LastInsertId()
	// checkErrs(err)

	// fmt.Println(id)

	//更新
	stmt, err := db.Prepare("update rate set USrate=?")
	checkErr(err)

	res, err := stmt.Exec(rates)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	resp, err = http.Get(url)
	fmt.Println(resp)
	time.Sleep(time.Second * 10)
}

func checkErrs(err error) {

	if err != nil {
		panic(err)
	}
}
