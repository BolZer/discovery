package mhgu

import (
	"discovery/discovery"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"strings"
)

const WikiSrc = "https://mhgu.kiranico.com/monster"

func StartDiscovery(callback func(entities []discovery.Entity)) error {
	var buffer []discovery.Entity

	crawler := colly.NewCollector(
		colly.AllowedDomains("mhgu.kiranico.com"),
		colly.AllowURLRevisit(),
	)

	entity := discovery.Entity{
		Game: discovery.MHGU,
	}

	crawler.OnHTML(`div[id="state-a"] table tbody tr:first-child`, func(element *colly.HTMLElement) {

		element.DOM.Find("td").Each(func(index int, selection *goquery.Selection) {

			value := strings.TrimSpace(selection.Text())

			// Fire
			if index == 4 {
				entity.Weakness.Fire = discovery.FormatStringToInt(value)
			}

			// Water
			if index == 5 {
				entity.Weakness.Water = discovery.FormatStringToInt(value)
			}

			// Thunder
			if index == 6 {
				entity.Weakness.Thunder = discovery.FormatStringToInt(value)
			}

			// Ice
			if index == 7 {
				entity.Weakness.Ice = discovery.FormatStringToInt(value)
			}

			// Dragon
			if index == 8 {
				entity.Weakness.Dragon = discovery.FormatStringToInt(value)
			}
		})

		buffer = append(buffer, entity)
	})

	crawler.OnHTML("div.card div.card-block div.table-responsive table.table-sm tbody tr td a", func(element *colly.HTMLElement) {

		if element.Request.URL.Path != "/monster" {
			return
		}

		entity.Name = strings.TrimSpace(element.DOM.Text())

		detailPageOfMonster, exists := element.DOM.Attr("href")

		if !exists {
			return
		}

		entity.Url = detailPageOfMonster

		if !strings.Contains(detailPageOfMonster, "monster") {
			return
		}

		err := crawler.Visit(element.Request.AbsoluteURL(detailPageOfMonster))

		if err != nil {
			panic(err)
		}
	})

	err := crawler.Visit(WikiSrc)

	if err != nil {
		return err
	}

	defer callback(buffer)

	return nil
}
