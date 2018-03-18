package grabber

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

//GetKuaidaili 从快代理抓取
func GetKuaidaili(ch chan *[]map[string]string) {
	request := gorequest.New()
	var url string
	var proxyList []map[string]string
	mapList := make(map[string]string)
	for i := 1; i <= kuaidailipage; i++ {
		url = setKuaidailiURL(i)
		resp, _, err := request.Get(url).
			Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8").
			Set("Referer", "https://www.kuaidaili.com/free").
			Set("Host", "www.kuaidaili.com").
			Set("User-Agent", GetUserAgent()).
			End()
		if err != nil {
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
		doc.Find(".table").Find("tbody").Find("tr").Each(func(i int, tr *goquery.Selection) {
			tr.Each(func(j int, td *goquery.Selection) {
				strArr := strings.Split(td.Text(), "\n")
				ip = strings.TrimSpace(strArr[1])
				port = strings.TrimSpace(strArr[2])
				types = strings.TrimSpace(strArr[4])
				local = strings.TrimSpace(strArr[5])
			})
			value, ok := mapList[ip]
			if ok && value == port {
			} else {
				mapList[ip] = port
				proxyList = append(proxyList, map[string]string{"ip": ip, "port": port, "type": types, "local": local})
			}
		})
		time.Sleep(1 * time.Second)
	}
	ch <- &proxyList
}
