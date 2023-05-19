// Copyright 2023 Roy Xu <ixw1991@126.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/imxw/h3c-auth.

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
