package h3cauth

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/imxw/h3c-auth/internal/pkg/encutil"
)

const (
	serverIP   = "10.0.100.20"
	serverPort = "8080"
	baseUrl    = "http://" + serverIP + ":" + serverPort
	appRootUrl = baseUrl + "/portal"
	authUrl    = appRootUrl + "/pws?t=li&ifEmailAuth=false"
)

var (
	userName = os.Getenv("USER")
	password = os.Getenv("PASSWORD")
)

func Auth() error {
	fmt.Println(userName, password)
	encodeRootUrl := encutil.UrlEncode(appRootUrl)
	encodePwd := encutil.PwdEncode(password)
	method := "POST"

	queryParams := fmt.Sprintf("userName=%s&userPwd=%s&serviceType=&language=chinese&usermac=mac&entrance=0&customPageId=3&send_dynamic_pwd_type=0&pwdMode=0&portalProxyIP=%s&portalProxyPort=50200&dcPwdNeedEncrypt=1&assignIpType=0&appRootUrl=%s&userurl=&userip=&basip=&wlannasid=&wlanssid=&loginVerifyCode=&userDynamicPwddd=&manualUrl=&manualUrlEncryptKey=", userName, encodePwd, serverIP, encodeRootUrl)

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
	if body == nil {
		return errors.New("认证失败")
	}
	return nil
}
