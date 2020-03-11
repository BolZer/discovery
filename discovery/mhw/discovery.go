package mhw

import (
	"discovery/discovery"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"strings"
)

const WikiSrc = "https://mhworld.kiranico.com/monsters"

func StartDiscovery(callback func(entities []discovery.Entity)) error {
	var buffer []discovery.Entity

	crawler := colly.NewCollector(
		colly.AllowedDomains("mhworld.kiranico.com"),
	)

	crawler.OnHTML("table.table.table-padded:first-child tbody tr", func(element *colly.HTMLElement) {

		if element.Request.URL.Path != "/monsters" {
			return
		}

		entity := discovery.Entity{
			Game: discovery.MHW,
		}

		element.DOM.Find("td").Each(func(index int, selection *goquery.Selection) {

			value := strings.TrimSpace(selection.Text())

			if index == 0 {
				entity.Name = value

				detailPageOfMonster, _ := selection.Find("a").Attr("href")

				entity.Url = detailPageOfMonster
			}

			if index == 1 {
				entity.Weakness.Fire = discovery.FormatStringToInt(value)
			}

			if index == 2 {
				entity.Weakness.Water = discovery.FormatStringToInt(value)
			}

			if index == 3 {
				entity.Weakness.Thunder = discovery.FormatStringToInt(value)
			}

			if index == 4 {
				entity.Weakness.Ice = discovery.FormatStringToInt(value)
			}

			if index == 5 {
				entity.Weakness.Dragon = discovery.FormatStringToInt(value)
			}
		})

		buffer = append(buffer, entity)
	})

	err := crawler.Visit(WikiSrc)

	if err != nil {
		return err
	}

	defer callback(buffer)

	return nil
}
