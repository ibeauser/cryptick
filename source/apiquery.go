package source

import (
	"strings"
)

type Report struct {
	Type string
	Name string
}
type Currency struct {
	Pair string
}
type Query struct {
	Source string
	Type   string
	Pair   string
	Url    string //final endpoint
	Wait   int
}

type Exchange = struct {
	ShortName  string
	LongName   string
	BaseURL    string // base endpoint
	PreQuery   string
	APIwait    int
	Currencies []Currency
	Reports    []Report
	Queries    []Query
}

var QCX = Exchange{
	ShortName: "QCX",
	LongName:  "QuadrigaCX",
	BaseURL:   "https://api.quadrigacx.com/v2/",
	PreQuery:  "?book=",
	APIwait:   10,
	Reports: []Report{
		Report{Name: "ticker", Type: "summary"},
		Report{Name: "transactions", Type: "detail"},
	},
	Currencies: []Currency{
		Currency{Pair: "eth_cad"},
		Currency{Pair: "btc_cad"},
		Currency{Pair: "eth_btc"},
	},
	Queries: []Query{},
}
var BIN = Exchange{
	ShortName: "BIN",
	LongName:  "Binance.com",
	BaseURL:   "https://api.binance.com/api/v1/",
	PreQuery:  "?symbol=",
	APIwait:   5,
	Reports: []Report{
		Report{Name: "trades", Type: "detail"},
		Report{Name: "ticker/24hr", Type: "summary"},
	},
	Currencies: []Currency{
		Currency{Pair: "ETHBTC"},
	},
	Queries: []Query{},
}

func APIQuery() []Query {
	var Queries []Query
	var Exchanges []Exchange
	Exchanges = append(Exchanges, QCX, BIN)
	for x, p := range Exchanges {
		part1 := []string{}
		part1 = append(part1, p.BaseURL)
		wait := p.APIwait
		cx := p.ShortName
		for _, q := range p.Reports {
			part2 := part1
			part2 = append(part2, q.Name)
			reptype := q.Type
			for _, r := range p.Currencies {
				part3 := part2
				part3 = append(part3, p.PreQuery)
				part3 = append(part3, r.Pair)
				pair := r.Pair
				Exchanges[x].Queries = []Query{{cx, reptype, pair, strings.Join(part3, ""), wait}}
				query := Exchanges[x].Queries[0]
				Queries = append(Queries, query)
			}
		}
	}
	return Queries
	panic("function failed")
}
