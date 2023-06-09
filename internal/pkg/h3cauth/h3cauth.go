// Copyright 2023 Roy Xu <ixw1991@126.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/imxw/h3c-auth.

package h3cauth

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/imxw/h3c-auth/internal/pkg/encutil"
)

type Config struct {
	Username string `yaml:"username,omitempty" validate:"required"`
	Password string `yaml:"password,omitempty" validate:"required"`
	IpAddr   string `yaml:"ipAddr,omitempty" validate:"required,ip"`
	Port     string `yaml:"port,omitempty" validate:"required,hostname_port"`
}

func Auth(cfg Config) error {

	username := cfg.Username
	password := cfg.Password
	serverIP := cfg.IpAddr
	serverPort := cfg.Port
	baseUrl := fmt.Sprintf("http://%s%s", serverIP, serverPort)
	appRootUrl := baseUrl + "/portal"
	authUrl := appRootUrl + "/pws?t=li&ifEmailAuth=false"

	encodeRootUrl := encutil.UrlEncode(appRootUrl)
	encodePwd := encutil.PwdEncode(password)
	method := "POST"

	queryParams := fmt.Sprintf("userName=%s&userPwd=%s&serviceType=&language=chinese&usermac=mac&entrance=0&customPageId=3&send_dynamic_pwd_type=0&pwdMode=0&portalProxyIP=%s&portalProxyPort=50200&dcPwdNeedEncrypt=1&assignIpType=0&appRootUrl=%s&userurl=&userip=&basip=&wlannasid=&wlanssid=&loginVerifyCode=&userDynamicPwddd=&manualUrl=&manualUrlEncryptKey=", username, encodePwd, serverIP, encodeRootUrl)

	payload := strings.NewReader(queryParams)

	client := &http.Client{}
	req, err := http.NewRequest(method, authUrl, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Origin", baseUrl)
	req.Header.Add("Referer", appRootUrl+"/index_pad.jsp")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "zh-CN,zh-Hans;q=0.9")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Accept", "text/plain, */*; q=0.01")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK || body == nil {
		return errors.New("认证失败")
	}
	return nil
}
