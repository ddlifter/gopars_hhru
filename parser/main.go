package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Item struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("hh.ru"),
	)

	var items []Item

	c.OnHTML("div[class=vacancy-serp-item-body__main-info]", func(h *colly.HTMLElement) {
		item := Item{
			Name:  h.ChildText("h3.bloko-header-section-3"),
			Price: h.ChildText("span[class=bloko-header-section-3]"),
		}
		items = append(items, item)
	})

	c.OnHTML("[class=bloko-button]", func(h *colly.HTMLElement) {
		next_page := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(next_page)
	})

	c.Visit("https://hh.ru/search/vacancy?text=Golang&area=1")
	fmt.Println(items)

}
