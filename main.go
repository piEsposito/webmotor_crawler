package main

import (
	"fmt"
	"sync"
	"time"
	"webmotor_crawler/crawl_functions"
)

func main() {

	//to measure time
	start := time.Now()

	//SETTING WG AND CHANNEL
	var wg sync.WaitGroup
	const PAGE_NUMBER int = 5
	const ROUTINE_NUMBER int = 5
	c := make(chan int, PAGE_NUMBER)

	wg.Add(1)
	go crawl_functions.FeedChannel(PAGE_NUMBER, c, &wg)
	wg.Wait()

	fmt.Println("got here")

	for i := 0; i < ROUTINE_NUMBER; i++ {

		wg.Add(1)
		go crawl_functions.CrawlRoutine(&wg, c)

	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Println("Crawling took %s", elapsed)
}
