package grabber

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"strings"
	"time"
)

func GetPcdaili(ch chan *[]map[string]string) {
	request := gorequest.New()
	var url string
	var proxyList []map[string]string
	mapList := make(map[string]string)
	for types := 1; types <= 4; types++ {
		if types == 2 || types == 3 {
			continue
		}
		for i := 1; i <= pcdailipage; i++ {
			url = setPcdailiURL(types, i)
			resp, _, err := request.Get(url).
				Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8").
				Set("Referer", "http://www.pcdaili.com/index.php").
				Set("Host", "www.pcdaili.com").
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
			doc.Find(".table").Find("tbody").Find("tr").Each(func(i int, tr *goquery.Selection) {
				tr.Each(func(j int, td *goquery.Selection) {
					strArr := strings.Split(td.Text(), "\n")
					if len(strArr) >= 5 {
						ip = strings.TrimSpace(strArr[1])
						port = strings.TrimSpace(strArr[2])
						types = strings.TrimSpace(strArr[4])
						local = strings.TrimSpace(strArr[5])
					}
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
	}
	ch <- &proxyList
}
