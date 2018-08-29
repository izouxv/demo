package util

import (
	"strings"
	"errors"
	"fmt"
	"encoding/json"
	log "github.com/cihub/seelog"
	"bytes"
)

func IPToString(intIp int64) (string, error) {
	i4 := intIp & 255
	i3 := intIp >> 8 & 255
	i2 := intIp >> 16 & 255
	i1 := intIp >> 24 & 255
	if i1 > 255 || i2 > 255 || i3 > 255 || i4 > 255 {
		return "", errors.New("isn't a int ip type")
	}
	ipString := fmt.Sprintf("%d.%d.%d.%d", i1, i2, i3, i4)
	return ipString, nil
}

func IPToInt64(ip string) (int64, error) {
	if ip == "127.0.0.1" {
		return 0, nil
	}
	ips := strings.Split(ip, ".")
	if len(ips) != 4 {
		return 0, errors.New("not a ip")
	}
	var intIP int64
	for k, v := range ips {
		i, err := StrToInt64(v)
		if err != nil || i > 255 {
			return 0, err
		}
		intIP = intIP | i<<uint64(8*(3-k))
	}
	return intIP, nil
}

type TaoBaoIpAddr struct {
	Code	int32	`json:"code"`
	Data	struct {
		Country		string	`json:"country"`
		Region		string	`json:"region"`
		City		string	`json:"city"`
		County 		string	`json:"county"`
		Isp 		string	`json:"isp"`
		CountryId 	string	`json:"country_id"`
	}	`json:"data"`
}

var(
	taoBaoIpToAddr = "http://ip.taobao.com/service/getIpInfo.php?ip="
	localIp = "127.0.0.1"
	localStr = "本地访问"
)

func IPToAddr(chanIp chan string,ip string) {
	var addrStr string
	defer func() {
		chanIp <- addrStr
	}()
	if ip == localIp {
		addrStr = localStr
		return
	}
	buffer := bytes.NewBuffer([]byte(""))
	body := httpReq(taoBaoIpToAddr+ip, "GET",buffer)
	if body != nil {
		log.Info("IPToAddr body is nil")
		addrStr = ""
		return
	}
	var taoBaoAddr TaoBaoIpAddr
	json.Unmarshal(body, &taoBaoAddr)
	if taoBaoAddr.Data.Region == taoBaoAddr.Data.City {
		addrStr = taoBaoAddr.Data.Country+ taoBaoAddr.Data.Region+ taoBaoAddr.Data.County+" "+ taoBaoAddr.Data.Isp+" "+ taoBaoAddr.Data.CountryId
		return
	}
	addrStr = taoBaoAddr.Data.Country+ taoBaoAddr.Data.Region+ taoBaoAddr.Data.City+ taoBaoAddr.Data.County+" "+ taoBaoAddr.Data.Isp+" "+ taoBaoAddr.Data.CountryId
}