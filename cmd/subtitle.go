package cmd

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func GetSubtitleLink(imdbid string) string {
	c := colly.NewCollector()

	subtitleDownloadLink := ""

	toBeInstalled := true

	c.OnHTML("td[align='center'] a", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		if strings.Contains(link, "subtitleserve") && toBeInstalled {
			toBeInstalled = false
			subtitleDownloadLink = fmt.Sprintf("https://www.opensubtitles.org/%s", link)
			fmt.Printf("%s", subtitleDownloadLink)
		}
	})

	link := fmt.Sprintf("https://www.opensubtitles.org/pb/search/sublanguageid-pob/imdbid-%s/sort-7/asc-0", imdbid)

	c.Visit(link)
	c.Wait()

	return subtitleDownloadLink
}
