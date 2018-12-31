package goutil

import (
	"net/url"
)

// MakeAbsoluteURL makes href absolute with repect to baseurl
func MakeAbsoluteURL(href, baseurl string) (string, error) {
	u, err := url.Parse(href)
	if err != nil {
		return "", err
	}
	base, err := url.Parse(baseurl)
	if err != nil {
		return "", err
	}
	u = base.ResolveReference(u)
	return u.String(), nil
}
