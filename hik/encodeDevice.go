package hik

import (
	"crypto/tls"
	"fmt"

	"io"
	"net/http"
	"net/url"
	"strings"
)

func GetEncodeDevice() {

	ip := IPaddr
	reqUrl := fmt.Sprintf("https://%s/sdmc/ui/vms/encoder/fetchPageQuery?", ip)

	// 创建 URL 对象并设置请求参数
	params := url.Values{}
	params.Set("pageNo", "1")    // 页数
	params.Set("pageSize", "20") // 每页数量
	params.Set("unitIdRange", "true")
	params.Set("syncAlready", "false")
	params.Set("regionId", "root000000")
	params.Set("hasEzviz", "true")

	keys := []string{"pageNo", "pageSize", "unitIdRange", "syncAlready", "regionId", "hasEzviz"}
	encodeParams := make([]string, len(keys))
	for i, key := range keys {
		encodeParams[i] = fmt.Sprintf("%s=%s", key, params.Get(key))
	}

	reqUrl = reqUrl + strings.Join(encodeParams, "&")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", reqUrl, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Cookie", Cookie)
	req.Header.Add("Host", IPaddr)
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(reqUrl)
	fmt.Println(string(body))
}

//	url := "https://113.55.126.45/sdmc/ui/vms/encoder/fetchPageQuery?pageNo=1&pageSize=100&unitIdRange=true&syncAlready=false&regionId=root000000&hasEzviz=true"
