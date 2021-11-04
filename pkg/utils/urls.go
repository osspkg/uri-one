package utils

import "net/url"

func IsValidUrl(data string) bool {
	if len(data) == 0 {
		return false
	}
	uri, err := url.Parse(data)
	if err != nil {
		return false
	}
	if len(uri.Scheme) == 0 || len(uri.Host) == 0 {
		return false
	}
	return true
}
