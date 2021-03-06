package main

import (
	"fmt"
	"sync"
	"time"
	"webmotor_crawler/crawl_functions"
)

func main() {

	/*
		in our main crawling function we do the 2 main tasks

		1. Populate our channel with the page numbers
		2. Create the goroutines to crawl from the channel IDs

	*/

	//to measure time
	start := time.Now()

	//SETTING WG AND CHANNEL
	var wg sync.WaitGroup
	const PAGE_START int = 13782
	const PAGE_NUMBER int = 20000
	const ROUTINE_NUMBER int = 100
	c := make(chan int, PAGE_NUMBER)

	wg.Add(1)
	go crawl_functions.FeedChannel(PAGE_START, PAGE_NUMBER, c, &wg)
	wg.Wait()

	fmt.Println("Channel fed with page numbers")

	//Creating the goroutines
	for i := 0; i < ROUTINE_NUMBER; i++ {

		wg.Add(1)
		go crawl_functions.CrawlRoutine(&wg, c)

	}

	wg.Wait()

	//Some time computation here
	elapsed := time.Since(start)
	fmt.Println("Crawling took %s", elapsed)
}
