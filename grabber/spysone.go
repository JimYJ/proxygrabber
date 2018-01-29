package grabber

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/nladuo/go-phantomjs-fetcher"
)

func GetSpyone() {
	// request := gorequest.New()
	// var proxyList []map[string]string
	// mapList := make(map[string]string)

	// b, err := ioutil.ReadAll(resp.Body)
	// log.Println(string(b), err)

	// resp, _, err := request.Post(spysone).
	// 	Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8").
	// 	Set("Referer", "http://spys.one/en/http-proxy-list/").
	// 	Set("Host", "spys.one").
	// 	Set("User-Agent", GetUserAgent()).
	// 	Send("xpp=5&xf1=4&xf2=2").
	// 	End()
	// if err != nil || resp.StatusCode != 200 {
	// 	if resp != nil {
	// 		printErrors(err, resp.Status)
	// 	} else {
	// 		printErrors(err)
	// 	}
	// 	return
	// }
	// doc, err2 := goquery.NewDocumentFromResponse(resp)
	// if err2 != nil {
	// 	printErrors(err2)
	// 	return
	// }
	// text := doc.Find("table").Text()

	// log.Println(text)

	//create a fetcher which seems to a httpClient
	fetcher, err := phantomjs.NewFetcher(2016, nil)
	defer fetcher.ShutDownPhantomJSServer()
	if err != nil {
		panic(err)
	}
	//inject the javascript you want to run in the webpage just like in chrome console.
	js_script := "function(){document.getElementById('kw').setAttribute('value', 'github');document.getElementById('su').click();}"
	//run the injected js_script at the end of loading html
	js_run_at := phantomjs.RUN_AT_DOC_END
	//send httpGet request with injected js
	resp, err := fetcher.GetWithJS(spysone, js_script, js_run_at)
	if err != nil {
		panic(err)
	}

	//select search results by goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.Content))
	if err != nil {
		panic(err)
	}
	fmt.Println("Results:")
	doc.Find(".c-container h3 a").Each(func(i int, contentSelection *goquery.Selection) {
		fmt.Println(i+1, "-->", contentSelection.Text())
	})
}
