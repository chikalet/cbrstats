package main

import (
	"fmt"
	"log"
	"time"

	"cbrstats/internal/cbr"
	"cbrstats/internal/stats"
)

func main() {
	loc, err := time.LoadLocation("Europe/Madrid")
	if err != nil {
		log.Printf("cannot load Europe/Madrid: %v, using Local", err)
		loc = time.Local
	}
	now := time.Now().In(loc)

	const days = 90
	agg := stats.NewAggregator()

	for i := 0; i < days; i++ {
		d := now.AddDate(0, 0, -i)
		dateReq := fmt.Sprintf("%02d/%02d/%04d", d.Day(), d.Month(), d.Year())

		data, err := cbr.FetchData(dateReq)
		if err != nil {
			log.Printf("date %s: fetch error: %v", dateReq, err)
			continue
		}

		vc, err := cbr.ParseValCurs(data)
		if err != nil {
			log.Printf("date %s: parse error: %v", dateReq, err)
			continue
		}

		for _, v := range vc.Valutes {
			rate, err := cbr.RateForValute(v)
			if err != nil {
				continue
			}
			agg.Add(rate, v.Name+" ("+v.CharCode+")", vc.Date)
		}

		log.Printf("date %s: processed %d valutes", dateReq, len(vc.Valutes))
	}

	res := agg.Result()
	fmt.Println("===========================================")
	fmt.Printf("MAX rate:  %.6f RUB — %s on %s\n", res.MaxRate, res.MaxCurrency, res.MaxDate)
	fmt.Printf("MIN rate:  %.6f RUB — %s on %s\n", res.MinRate, res.MinCurrency, res.MinDate)
	fmt.Println()
	fmt.Printf("Average rate: %.6f RUB\n", res.AverageRate)
	fmt.Printf("Total entries used: %d\n", res.TotalCount)
	fmt.Println("===========================================")
}
