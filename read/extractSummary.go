package read

// tradesummary relation, crypto database

import (
	"cryotick/utils"
	"fmt"
	"strconv"
)

type outSummary struct {
	Exchange   string
	Book       string
	Timestamp  int
	High24     float64
	Low24      float64
	Lastprice  float64
	Lowestsell float64
	Highestbuy float64
	Volume     float64
	Vwap       float64
}

func SelectSummary() {
	db := utils.ConnectDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM tradesummary;")
	utils.CheckErr("db.Query", err)
	defer rows.Close()
	var dayOut outSummary
	for rows.Next() {
		err := rows.Scan(&dayOut.Exchange, &dayOut.Book, &dayOut.Timestamp, &dayOut.High24, &dayOut.Low24, &dayOut.Lastprice, &dayOut.Lowestsell, &dayOut.Highestbuy, &dayOut.Volume, &dayOut.Vwap)
		utils.CheckErr("rows.Scan", err)
		fmt.Println(
			"Exchange: ", dayOut.Exchange,
			", Book: ", dayOut.Book,
			", DateTime: ", utils.FromUnixTimestamp(dayOut.Timestamp),
			", Last Price: ", strconv.FormatFloat(dayOut.Lastprice, 'f', 2, 64),
		)
	}
}
