package collect

import (
	"encoding/json"
	"github.com/ibeauser/cryptick/utils"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type inSummary struct {
	Exchange  string
	Book      string
	Timestamp string
	High      string
	Low       string
	Last      string
	Ask       string
	Bid       string
	Volume    string
	Vwap      string
}

func InsertSummary(exchange string, book string, url string) time.Time {
	db := utils.ConnectDB()
	defer db.Close()
	resp, err := http.Get(url)
	utils.CheckErr("http.Get", err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	utils.CheckErr("ioutil.ReadAll", err)

	var dayIn inSummary
	err = json.Unmarshal(body, &dayIn)
	if err != nil {
		log.Printf("error decoding exchange response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("exchange response: %q", body)
	}
	timestamp, err := strconv.ParseInt(dayIn.Timestamp, 10, 64)
	utils.CheckErr("strconv.ParseInt(day.Timestamp, 10, 64)", err)
	high, err := strconv.ParseFloat(dayIn.High, 64)
	utils.CheckErr("strconv.ParseFloat(day.High, 64)", err)
	low, err := strconv.ParseFloat(dayIn.Low, 64)
	utils.CheckErr("strconv.ParseFloat(day.Low, 64)", err)
	last, err := strconv.ParseFloat(dayIn.Last, 64)
	utils.CheckErr("strconv.ParseFloat(day.Last, 64)", err)
	ask, err := strconv.ParseFloat(dayIn.Ask, 64)
	utils.CheckErr("strconv.ParseFloat(day.Ask, 64)", err)
	bid, err := strconv.ParseFloat(dayIn.Bid, 64)
	utils.CheckErr("strconv.ParseFloat(day.Bid, 64)", err)
	volume, err := strconv.ParseFloat(dayIn.Volume, 64)
	utils.CheckErr("strconv.ParseFloat(day.Volume, 64)", err)
	vwap, err := strconv.ParseFloat(dayIn.Vwap, 64)
	utils.CheckErr("strconv.ParseFloat(day.Vwap, 64)", err)

	row, err := db.Query("INSERT INTO tradesummary VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) ON CONFLICT (timestamp) DO NOTHING returning timestamp;",
		exchange, book, timestamp, high, low, last, ask, bid, volume, vwap)
	utils.CheckErr("db.QueryRow", err)
	defer row.Close()
	insertCompleteTime := time.Now()
	return insertCompleteTime
}
