package main

import (
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/jake-walker/vh7-cli/vh7"
	"github.com/pterm/pterm"
)

func fullUrl(link string) string {
	url, err := url.Parse(apiUrl)

	if err != nil {
		pterm.Fatal.Println("There was a problem building the URL", err)
	}

	url.Path = path.Join(url.Path, link)
	return url.String()
}

func cleanLink(link string) string {
	link = strings.Replace(link, apiUrl, "", 1)
	link = strings.ReplaceAll(link, "/", "")
	return link
}

func prettyDate(date vh7.UtcTime) string {
	if (date == vh7.UtcTime{}) {
		return "never"
	}
	return time.Time(date).Format("2 Jan 2006")
}
