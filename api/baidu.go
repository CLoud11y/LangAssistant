package api

import (
	"LangAssistant/config"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// baidu api response struct
type BaiduResp struct {
	From        string `json:"from"`
	To          string `json:"to"`
	TransResult []struct {
		Src string `json:"src"`
		Dst string `json:"dst"`
	} `json:"trans_result"`
}

var baseUrl string = "https://fanyi-api.baidu.com/api/trans/vip/translate"
var from string = "en"
var to string = "zh"
var appid string = config.Conf.BAIDU_API.Appid
var salt string = config.Conf.BAIDU_API.Salt
var key string = config.Conf.BAIDU_API.Key

func BaiduTranslate(text string) (*BaiduResp, error) {
	// check if text is empty
	if text == "" {
		return nil, fmt.Errorf("text is empty")
	}
	fmt.Println("baidu translate text: ", text)

	// create request url that fits baidu api
	s1 := fmt.Sprintf("%s%s%s%s", appid, text, salt, key)
	sign := fmt.Sprintf("%x", md5.Sum([]byte(s1)))
	url := fmt.Sprintf("%s?q=%s&from=%s&to=%s&appid=%s&salt=%s&sign=%s", baseUrl, url.QueryEscape(text), from, to, appid, salt, sign)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("baidu response: ", string(body))

	data := &BaiduResp{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
