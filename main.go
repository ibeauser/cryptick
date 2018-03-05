package main

import (
	"cryptick/collect"
	// "./read"
	"cryptick/source"
	"fmt"
	"time"
)

// const (
// 	timerPeriod time.Duration = 15 * time.Second // Interval to wake up on.
// )

func main() {
	// datastore.SelectTradeSummary()
	// wait := timerPeriod

	queries := source.APIQuery()
	for {
		for _, query := range queries {
			if query.Source == "QCX" && query.Type == "summary" {
				starttime := time.Now()
				insertTime := collect.InsertSummary(query.Source, query.Pair, query.Url)
				difference := insertTime.Sub(starttime)
				fmt.Printf("%s, %s, %s, %s, %s\n", query.Source, query.Type, query.Pair, query.Url, difference)
				time.Sleep(time.Second * time.Duration(query.Wait))
			}
			// read.SelectSummary()
		}
	}
}
