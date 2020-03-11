package mh4u

import (
	"discovery/discovery"
	"encoding/json"
	"github.com/gocolly/colly/v2"
	"strings"
)

const WikiSrc = "http://kiranico.com/en/mh4u/monster"

func StartDiscovery(callback func(entities []discovery.Entity)) error {
	var buffer []discovery.Entity

	crawler := colly.NewCollector()

	entity := discovery.Entity{
		Game: discovery.MH4U,
	}

	crawler.OnHTML("html body", func(element *colly.HTMLElement) {

		script := element.DOM.Find("script").Eq(2).Text()

		if script == "" {
			return
		}

		jsonData := script[strings.Index(script, "{") : len(script)-1]
		jsonData = strings.Replace(jsonData, "{};window.js_vars = ", "", -1)

		scrappedEntity := Monster{}
		_ = json.Unmarshal([]byte(jsonData), &scrappedEntity)

		if len(scrappedEntity.Monster.Monsterbodyparts) == 0 {
			return
		}

		entity.Name = scrappedEntity.Monster.LocalName
		entity.Weakness.Fire = scrappedEntity.Monster.Monsterbodyparts[0].Pivot.ResFire
		entity.Weakness.Ice = scrappedEntity.Monster.Monsterbodyparts[0].Pivot.ResIce
		entity.Weakness.Dragon = scrappedEntity.Monster.Monsterbodyparts[0].Pivot.ResDragon
		entity.Weakness.Thunder = scrappedEntity.Monster.Monsterbodyparts[0].Pivot.ResThunder
		entity.Weakness.Water = scrappedEntity.Monster.Monsterbodyparts[0].Pivot.ResWater

		buffer = append(buffer, entity)
	})

	crawler.OnHTML("div.container div.col-sm-2.col-xs-6 p a", func(element *colly.HTMLElement) {

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
