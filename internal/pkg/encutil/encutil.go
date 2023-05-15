package encutil

import (
	"encoding/base64"
	"net/url"
)

func UrlEncode(query string) string {
	return url.QueryEscape(query)
}

func PwdEncode(password string) string {
	base64Pwd := base64.StdEncoding.EncodeToString([]byte(password))
	return UrlEncode(base64Pwd)
}
