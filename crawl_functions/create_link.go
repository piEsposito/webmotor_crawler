package crawl_functions

import (
	"fmt"
	"strconv"
	"sync"
)

func FeedChannel(page_start int, page_nmr int, c chan int, wg *sync.WaitGroup) {

	for i := page_start; i < page_nmr; i++ {
		c <- i
	}
	wg.Done()

}

func CreateLink(idx int) (string, string) {

	nmr := strconv.Itoa(idx)

	path := fmt.Sprintf("%s.json", nmr)
	link_base_1 := "https://www.webmotors.com.br/api/search/car?url=https://www.webmotors.com.br/carros-usados%%2Festoque%%3Finst%%3Dheader%%3Awebmotors%%3Aheader-deslogado%%3A%%3Acarros-usados-ou-seminovos&actualPage="
	link_base_2 := "&displayPerPage=23&order=1&showMenu=true&showCount=true&showBreadCrumb=true&testAB=false&returnUrl=false"

	link := fmt.Sprintf(link_base_1 + nmr + link_base_2)

	return link, path

}
