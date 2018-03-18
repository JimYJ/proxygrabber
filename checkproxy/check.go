package checkproxy

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"log"
	"time"
)

var (
	testURL = []string{
		"http://www.baidu.com",
		"http://www.google.com",
	}
)
var (
	proxyURL string
)

//Check 检测代理是否失效
func Check(proxylist *[]map[string]string) []map[string]string {
	ch := make(chan map[string]string, len(*proxylist))
	for i := 0; i < 20; i++ {
		proxyURL = fmt.Sprintf("%s://%s:%s", (*proxylist)[i]["type"], (*proxylist)[i]["ip"], (*proxylist)[i]["port"])
		go checkproxy(proxyURL, ch, (*proxylist)[i])
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

func checkproxy(proxyURL string, ch chan map[string]string, proxyMap map[string]string) {
	fineProxy := make(map[string]string)
	log.Println(proxyURL)
	request := gorequest.New().Proxy(proxyURL).Timeout(10 * time.Second)
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
