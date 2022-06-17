package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	start := time.Now()
	stocks := [4]string{
		"AAPL",
		"MSFT",
		"PLTR",
		"AMZN",
	}

	parseStocks(stocks)
	fmt.Printf("Completed the code process, took: %f seconds\n", time.Since(start).Seconds())
}

func parseStocks(stocks [4]string) {
	ch := make(chan string)
	var wg sync.WaitGroup

	for _, stock := range stocks {
		wg.Add(1)
		go parseStock(stock, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}

func parseStock(stock string, ch chan string, wg *sync.WaitGroup) {
	defer (*wg).Done()

	c := colly.NewCollector(
		colly.AllowedDomains("finance.yahoo.com"),
		colly.MaxBodySize(0),
		colly.AllowURLRevisit(),
		colly.Async(true),
	)

	// Set max Parallelism and introduce a Random Delay
	_ = c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
	})
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())

	})

	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			var dataSlice []string
			el.ForEach("td", func(_ int, el *colly.HTMLElement) {
				dataSlice = append(dataSlice, el.Text)
			})

			if dataSlice[0] == "Previous Close" {
				ch <- stock + " Price for previous close is: " + dataSlice[1]
			}
		})
	})

	_ = c.Visit("https://finance.yahoo.com/quote/" + stock)

	c.Wait()
}
