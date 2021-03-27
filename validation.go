package main

import "net/url"

func validateUrl(test string) bool {
	_, err := url.ParseRequestURI(test)
	if err != nil {
		return false
	}

	u, err := url.Parse(test)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
