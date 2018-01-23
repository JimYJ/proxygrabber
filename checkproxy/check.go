package checkproxy

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"log"
	"time"
)

var (
	testUrl = []string{
		"http://www.baidu.com",
		"http://www.google.com",
	}
)
var (
	proxyUrl string
)

func Check(proxylist *[]map[string]string) []map[string]string {
	ch := make(chan map[string]string, len(*proxylist))
	for i := 0; i < 20; i++ {
		proxyUrl = fmt.Sprintf("%s://%s:%s", (*proxylist)[i]["type"], (*proxylist)[i]["ip"], (*proxylist)[i]["port"])
		go checkproxy(proxyUrl, ch, (*proxylist)[i])
	}
	var fineProxyList []map[string]string
	for i := 0; i < 20; i++ {
		fineProxy := <-ch
		if fineProxy != nil {
			fineProxyList = append(fineProxyList, fineProxy)
		}
	}
	return fineProxyList
}

func checkproxy(proxyUrl string, ch chan map[string]string, proxyMap map[string]string) {
	fineProxy := make(map[string]string)
	log.Println(proxyUrl)
	request := gorequest.New().Proxy(proxyUrl).Timeout(10 * time.Second)
	resp, _, err := request.Get("http://m.chn.lottedfs.com/kr").End()
	if err != nil {
		log.Println(err)
	}
	if resp != nil {
		log.Println(resp.StatusCode)
		if resp.StatusCode == 200 {
			fineProxy = proxyMap
		} else {
			fineProxy = nil
		}
	} else {
		fineProxy = nil
	}
	ch <- fineProxy

}
