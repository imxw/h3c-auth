// Copyright 2023 Roy Xu <imxw1991@126.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/imxw/h3c-auth.

package netutil

import (
	"net"
	"net/http"
)

func GetLocalIP() (ip net.IP, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP, nil
	}
	return
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

	url := "http://connect.rom.miui.com/generate_204"
	method := "GET"

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	return res.StatusCode == 204
}
