// Copyright 2023 Roy Xu <ixw1991@126.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/imxw/h3c-auth.

package netutil

import (
	"net"
	"net/http"
	"strings"
	"time"
)

func GetLocalIP() ([]net.IP, error) {
	var ips []net.IP
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, interf := range interfaces {
		// 忽略以 "utun" 开头的接口
		if strings.HasPrefix(interf.Name, "utun") {
			continue
		}

		addrs, err := interf.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			ipAddr, ok := addr.(*net.IPNet)
			if !ok || ipAddr.IP.IsLoopback() || !ipAddr.IP.IsGlobalUnicast() {
				continue
			}
			ips = append(ips, ipAddr.IP)
		}
	}
	return ips, nil
}

func IsIpInNet(ipAddr net.IP, network string) bool {

	_, ipNet, err := net.ParseCIDR(network)
	if err != nil {
		panic(err.Error())
	}

	if ipNet.Contains(ipAddr) {
		return true
	} else {
		return false
	}

}
func IsNetOk() bool {

	// 参考： https://imldy.cn/posts/99d42f85/
	urls := []string{
		"http://connect.rom.miui.com/generate_204",
		"http://wifi.vivo.com.cn/generate_204",
		"http://www.apple.com/library/test/success.html",
		"http://connectivitycheck.platform.hicloud.com/generate_204",
		"http://www.google.com/generate_204",
	}

	for _, url := range urls {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			continue
		}

		client := &http.Client{Timeout: 5 * time.Second}
		res, err := client.Do(req)
		if err != nil {
			continue
		}
		defer res.Body.Close()

		if res.StatusCode == 204 || res.StatusCode == 200 {
			return true
		}
	}

	return false
}
