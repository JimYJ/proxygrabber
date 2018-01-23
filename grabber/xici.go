package grabber

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"strings"
	"time"
)

func GetXicidaili(ch chan *[]map[string]string) {
	request := gorequest.New()
	var url string
	var proxyList []map[string]string
	mapList := make(map[string]string)
	for i := 1; i <= xicidailipage; i++ {
		url = setxXcidailiiUrl(i)
		resp, _, err := request.Get(url).
			Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8").
			Set("Referer", "http://www.xicidaili.com/nn/").
			Set("Host", "www.xicidaili.com").
			Set("User-Agent", GetUserAgent()).
			End()
		if err != nil || resp.StatusCode != 200 {
			if resp != nil {
				printErrors(err, resp.Status)
			} else {
				printErrors(err)
			}
			return
		}
		doc, err2 := goquery.NewDocumentFromResponse(resp)
		if err != nil {
			printErrors(err2)
			return
		}
		var ip, port, types, local string
		doc.Find("#ip_list").Find("tr").Each(func(i int, tr *goquery.Selection) {
			tr.Each(func(j int, td *goquery.Selection) {
				td.Each(func(j int, e *goquery.Selection) {
					strArr := strings.Split(strings.TrimSpace(e.Text()), "\n")
					strArr = removeSpaceElem(strArr)
					ip = strings.TrimSpace(strArr[0])
					port = strings.TrimSpace(strArr[1])
					types = strings.TrimSpace(strArr[4])
					local = strings.TrimSpace(strArr[2])
				})
			})
			value, ok := mapList[ip]
			if ok && value == port {
			} else {
				if ip != "国家" {
					mapList[ip] = port
					proxyList = append(proxyList, map[string]string{"ip": ip, "port": port, "type": types, "local": local})
				}
			}
		})
		time.Sleep(1 * time.Second)
	}
	// log.Println(proxyList)
	ch <- &proxyList
}
