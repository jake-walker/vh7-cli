package main

import (
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/pterm/pterm"
	"github.com/quark-links/quark-cli/quark"
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

func prettyDate(date quark.UtcTime) string {
	if (date == quark.UtcTime{}) {
		return "never"
	}
	return time.Time(date).Format("2 Jan 2006")
}

func prettyTime(date time.Time) string {
	return date.Format("2 Jan 2006 15:04")
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
